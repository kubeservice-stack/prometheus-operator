// Copyright 2017 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/textparse"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"

	"github.com/prometheus-operator/prometheus-operator/pkg/k8sutil"
)

func SourceToIOReader(source string) (io.Reader, error) {
	if strings.HasPrefix(source, "http") {
		return URLToIOReader(source)
	}
	return PathToOSFile(source)
}

func PathToOSFile(relativePath string) (*os.File, error) {
	path, err := filepath.Abs(relativePath)
	if err != nil {
		return nil, fmt.Errorf("failed generate absolute file path of %s: %w", relativePath, err)
	}

	manifest, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}

	return manifest, nil
}

func URLToIOReader(url string) (io.Reader, error) {
	var resp *http.Response
	timeout := 30 * time.Second

	err := wait.PollUntilContextTimeout(context.Background(), time.Second, timeout, false, func(_ context.Context) (bool, error) {
		var err error
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return true, nil
		}
		return false, nil
	})

	if err != nil {
		return nil, fmt.Errorf(
			"waiting for %v to return a successful status code timed out. Last response from server was: %v: %w",
			url,
			resp,
			err,
		)
	}

	return resp.Body, nil
}

// WaitForPodsReady waits for a selection of Pods to be running and each
// container to pass its readiness check.
func (f *Framework) WaitForPodsReady(ctx context.Context, namespace string, timeout time.Duration, expectedReplicas int, opts metav1.ListOptions) error {
	return wait.PollUntilContextTimeout(ctx, time.Second, timeout, false, func(ctx context.Context) (bool, error) {
		pl, err := f.KubeClient.CoreV1().Pods(namespace).List(ctx, opts)
		if err != nil {
			return false, err
		}

		runningAndReady := 0
		for _, p := range pl.Items {
			isRunningAndReady, err := k8sutil.PodRunningAndReady(p)
			if err != nil {
				return false, nil
			}

			if isRunningAndReady {
				runningAndReady++
			}
		}

		if runningAndReady == expectedReplicas {
			return true, nil
		}
		return false, nil
	})
}

func (f *Framework) WaitForPodsRunImage(ctx context.Context, namespace string, expectedReplicas int, image string, opts metav1.ListOptions) error {
	return wait.PollUntilContextTimeout(ctx, time.Second, time.Minute*5, false, func(ctx context.Context) (bool, error) {
		pl, err := f.KubeClient.CoreV1().Pods(namespace).List(ctx, opts)
		if err != nil {
			return false, err
		}

		runningImage := 0
		for _, p := range pl.Items {
			if podRunsImage(p, image) {
				runningImage++
			}
		}

		if runningImage == expectedReplicas {
			return true, nil
		}
		return false, nil
	})
}

func WaitForHTTPSuccessStatusCode(timeout time.Duration, url string) error {
	var resp *http.Response
	err := wait.PollUntilContextTimeout(context.Background(), time.Second, timeout, false, func(_ context.Context) (bool, error) {
		var err error
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return true, nil
		}
		return false, nil
	})

	if err != nil {
		return fmt.Errorf(
			"waiting for %v to return a successful status code timed out. Last response from server was: %v: %w",
			url,
			resp,
			err,
		)
	}
	return nil
}

func podRunsImage(p v1.Pod, image string) bool {
	for _, c := range p.Spec.Containers {
		if image == c.Image {
			return true
		}
	}

	return false
}

// ProxyGetPod executes an HTTP(S) request against the default port of the pod
// using the Proxy API.
func (f *Framework) ProxyGetPod(ctx context.Context, scheme, namespace, pod, path string) ([]byte, error) {
	b, err := f.KubeClient.
		CoreV1().
		Pods(namespace).
		ProxyGet(scheme, pod, "", path, nil).
		DoRaw(ctx)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ProxyPostPod expects resourceName as "[protocol:]podName[:portNameOrNumber]".
// protocol is optional and the valid values are "http" and "https".
// Without specifying protocol, "http" will be used.
// podName is mandatory.
// portNameOrNumber is optional.
// Without specifying portNameOrNumber, default port will be used.
func (f *Framework) ProxyPostPod(namespace, resourceName, path, body string) *rest.Request {
	return f.KubeClient.
		CoreV1().
		RESTClient().
		Post().
		Namespace(namespace).
		Resource("pods").
		SubResource("proxy").
		Name(resourceName).
		Suffix(path).
		Body([]byte(body)).
		SetHeader("Content-Type", "application/json")
}

// GetMetricValueFromPod sends an HTTP(S) request to the /metrics endpoint of the pod
// using the Proxy API, parses the response and returns the flot64 value of the
// first series matching the metric name.
// If protocol is empty, HTTP is used.
// If portNumberOfName is empty, the default pod's port is used.
func (f *Framework) GetMetricValueFromPod(ctx context.Context, protocol, ns, podName, portNumberOrName, metricName string) (float64, error) {
	b, err := f.KubeClient.
		CoreV1().
		Pods(ns).
		ProxyGet(protocol, podName, portNumberOrName, "/metrics", nil).
		DoRaw(ctx)
	if err != nil {
		return 0, fmt.Errorf("error reading /metrics: %w", err)
	}

	return getMetricValue(b, metricName)
}

// GetMetricValueFromService sends an HTTP(S) request to the /metrics endpoint
// of the service using the Proxy API, parses the response and returns the
// flot64 value of the first series matching the metric name.
// If protocol is empty, HTTP is used.
// If portNumberOfName is empty, the default pod's port is used.
func (f *Framework) EnsureMetricsFromService(ctx context.Context, protocol, ns, service, portNumberOrName string, metrics ...string) error {
	if len(metrics) == 0 {
		return fmt.Errorf("need to provide at least 1 metric to check")
	}

	b, err := f.KubeClient.
		CoreV1().
		Services(ns).
		ProxyGet(protocol, service, portNumberOrName, "/metrics", nil).
		DoRaw(ctx)
	if err != nil {
		return fmt.Errorf("error reading /metrics: %w", err)
	}

	for _, m := range metrics {
		_, err = getMetricValue(b, m)
		if err != nil {
			return fmt.Errorf("metric %s: %w", m, err)
		}
	}

	return nil
}

func getMetricValue(b []byte, metricName string) (float64, error) {
	parser := textparse.NewPromParser(b, labels.NewSymbolTable())

	for {
		entry, err := parser.Next()
		if err != nil {
			return 0, err
		}

		if entry == textparse.EntryInvalid {
			return 0, fmt.Errorf("invalid prometheus metric entry")
		}

		if entry != textparse.EntrySeries {
			continue
		}

		seriesLabels := labels.Labels{}
		parser.Labels(&seriesLabels)

		if seriesLabels.Get("__name__") != metricName {
			continue
		}

		_, _, val := parser.Series()

		return val, nil
	}
}
