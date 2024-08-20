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

package v1beta1

import (
	v1beta1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1beta1"
)

// MatcherApplyConfiguration represents a declarative configuration of the Matcher type for use
// with apply.
type MatcherApplyConfiguration struct {
	Name      *string            `json:"name,omitempty"`
	Value     *string            `json:"value,omitempty"`
	MatchType *v1beta1.MatchType `json:"matchType,omitempty"`
}

// MatcherApplyConfiguration constructs a declarative configuration of the Matcher type for use with
// apply.
func Matcher() *MatcherApplyConfiguration {
	return &MatcherApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MatcherApplyConfiguration) WithName(value string) *MatcherApplyConfiguration {
	b.Name = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *MatcherApplyConfiguration) WithValue(value string) *MatcherApplyConfiguration {
	b.Value = &value
	return b
}

// WithMatchType sets the MatchType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MatchType field is set to the value of the last call.
func (b *MatcherApplyConfiguration) WithMatchType(value v1beta1.MatchType) *MatcherApplyConfiguration {
	b.MatchType = &value
	return b
}
