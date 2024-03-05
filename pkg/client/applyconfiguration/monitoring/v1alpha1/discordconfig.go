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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// DiscordConfigApplyConfiguration represents an declarative configuration of the DiscordConfig type for use
// with apply.
type DiscordConfigApplyConfiguration struct {
	SendResolved   *bool                         `json:"sendResolved,omitempty"`
	WebhookURL     *v1.SecretKeySelector         `json:"apiURL,omitempty"`
	WebhookURLFile *string                       `json:"webhookURLFile,omitempty"`
	Title          *string                       `json:"title,omitempty"`
	Message        *string                       `json:"message,omitempty"`
	HTTPConfig     *HTTPConfigApplyConfiguration `json:"httpConfig,omitempty"`
}

// DiscordConfigApplyConfiguration constructs an declarative configuration of the DiscordConfig type for use with
// apply.
func DiscordConfig() *DiscordConfigApplyConfiguration {
	return &DiscordConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *DiscordConfigApplyConfiguration) WithSendResolved(value bool) *DiscordConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithWebhookURL sets the WebhookURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookURL field is set to the value of the last call.
func (b *DiscordConfigApplyConfiguration) WithWebhookURL(value v1.SecretKeySelector) *DiscordConfigApplyConfiguration {
	b.WebhookURL = &value
	return b
}

// WithWebhookURLFile sets the WebhookURLFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookURLFile field is set to the value of the last call.
func (b *DiscordConfigApplyConfiguration) WithWebhookURLFile(value string) *DiscordConfigApplyConfiguration {
	b.WebhookURLFile = &value
	return b
}

// WithTitle sets the Title field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Title field is set to the value of the last call.
func (b *DiscordConfigApplyConfiguration) WithTitle(value string) *DiscordConfigApplyConfiguration {
	b.Title = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *DiscordConfigApplyConfiguration) WithMessage(value string) *DiscordConfigApplyConfiguration {
	b.Message = &value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *DiscordConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *DiscordConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}
