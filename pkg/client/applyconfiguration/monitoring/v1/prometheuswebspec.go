// Copyright The prometheus-operator Authors
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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PrometheusWebSpecApplyConfiguration represents a declarative configuration of the PrometheusWebSpec type for use
// with apply.
type PrometheusWebSpecApplyConfiguration struct {
	WebConfigFileFieldsApplyConfiguration `json:",inline"`
	PageTitle                             *string `json:"pageTitle,omitempty"`
	MaxConnections                        *int32  `json:"maxConnections,omitempty"`
}

// PrometheusWebSpecApplyConfiguration constructs a declarative configuration of the PrometheusWebSpec type for use with
// apply.
func PrometheusWebSpec() *PrometheusWebSpecApplyConfiguration {
	return &PrometheusWebSpecApplyConfiguration{}
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *PrometheusWebSpecApplyConfiguration) WithTLSConfig(value *WebTLSConfigApplyConfiguration) *PrometheusWebSpecApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *PrometheusWebSpecApplyConfiguration) WithHTTPConfig(value *WebHTTPConfigApplyConfiguration) *PrometheusWebSpecApplyConfiguration {
	b.HTTPConfig = value
	return b
}

// WithPageTitle sets the PageTitle field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PageTitle field is set to the value of the last call.
func (b *PrometheusWebSpecApplyConfiguration) WithPageTitle(value string) *PrometheusWebSpecApplyConfiguration {
	b.PageTitle = &value
	return b
}

// WithMaxConnections sets the MaxConnections field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxConnections field is set to the value of the last call.
func (b *PrometheusWebSpecApplyConfiguration) WithMaxConnections(value int32) *PrometheusWebSpecApplyConfiguration {
	b.MaxConnections = &value
	return b
}
