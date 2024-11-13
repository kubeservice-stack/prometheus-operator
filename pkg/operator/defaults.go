// Copyright 2020 The prometheus-operator Authors
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

package operator

import "github.com/prometheus/common/version"

const (
	// DefaultAlertmanagerVersion is a default image tag for the prometheus alertmanager.
	DefaultAlertmanagerVersion = "v0.27.0"
	// DefaultAlertmanagerBaseImage is a base container registry address for the prometheus alertmanager.
	DefaultAlertmanagerBaseImage = "quay.io/prometheus/alertmanager"
	// DefaultAlertmanagerImage is a default image pulling address for the prometheus alertmanager.
	DefaultAlertmanagerImage = DefaultAlertmanagerBaseImage + ":" + DefaultAlertmanagerVersion

	// DefaultThanosVersion is a default image tag for the Thanos long-term prometheus storage collector.
	DefaultThanosVersion = "v0.36.1"
	// DefaultThanosBaseImage is a base container registry address for the Thanos long-term prometheus
	// storage collector.
	DefaultThanosBaseImage = "quay.io/thanos/thanos"
	// DefaultThanosImage is a default image pulling address for the Thanos long-term prometheus storage collector.
	DefaultThanosImage = DefaultThanosBaseImage + ":" + DefaultThanosVersion
)

var (
	// DefaultPrometheusVersion is a default image tag for the prometheus.
	DefaultPrometheusVersion = PrometheusCompatibilityMatrix[len(PrometheusCompatibilityMatrix)-1]
	// DefaultPrometheusExperimentalVersion is a default image tag for the prometheus experimental version (like Prometheus 3 beta).
	DefaultPrometheusExperimentalVersion = PrometheusExperimentalVersions[len(PrometheusExperimentalVersions)-1]
	// DefaultPrometheusBaseImage is a base container registry address for the prometheus.
	DefaultPrometheusBaseImage = "quay.io/prometheus/prometheus"
	// DefaultPrometheusImage is a default image pulling address for the prometheus.
	DefaultPrometheusImage = DefaultPrometheusBaseImage + ":" + DefaultPrometheusVersion

	// DefaultPrometheusConfigReloaderImage is an image that will be used as a sidecar to provide dynamic prometheus
	// configuration reloading.
	DefaultPrometheusConfigReloaderImage = "quay.io/prometheus-operator/prometheus-config-reloader:v" + version.Version

	// PrometheusCompatibilityMatrix is a list of supported prometheus versions.
	// prometheus-operator end-to-end tests verify that the operator can deploy from the current LTS version to the latest stable release.
	// This list should be updated every time a new LTS is released.
	PrometheusCompatibilityMatrix = []string{
		"v2.45.0",
		"v2.46.0",
		"v2.47.0",
		"v2.47.1",
		"v2.47.2",
		"v2.48.0",
		"v2.48.1",
		"v2.49.0",
		"v2.49.1",
		"v2.50.0",
		"v2.50.1",
		"v2.51.0",
		"v2.51.1",
		"v2.51.2",
		"v2.52.0",
		// The v2.52.1 image tag is missing from docker.io and quay.io registries.
		// "v2.52.1",
		"v2.53.0",
		"v2.53.1",
		"v2.53.2",
		"v2.53.3",
		"v2.54.0",
		"v2.54.1",
		"v2.55.0",
		"v2.55.1",
	}

	// Note: Issues in this version won't be supported by operator till its stable
	// This is only added for users to try the unstable versions.
	PrometheusExperimentalVersions = []string{
		"v3.0.0-beta.0",
		"v3.0.0-beta.1",
		"v3.0.0-rc.0",
		"v3.0.0-rc.1",
	}
)
