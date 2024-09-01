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

// TelegramConfigApplyConfiguration represents a declarative configuration of the TelegramConfig type for use
// with apply.
type TelegramConfigApplyConfiguration struct {
	SendResolved         *bool                                `json:"sendResolved,omitempty"`
	APIURL               *string                              `json:"apiURL,omitempty"`
	BotToken             *SecretKeySelectorApplyConfiguration `json:"botToken,omitempty"`
	BotTokenFile         *string                              `json:"botTokenFile,omitempty"`
	ChatID               *int64                               `json:"chatID,omitempty"`
	Message              *string                              `json:"message,omitempty"`
	DisableNotifications *bool                                `json:"disableNotifications,omitempty"`
	ParseMode            *string                              `json:"parseMode,omitempty"`
	HTTPConfig           *HTTPConfigApplyConfiguration        `json:"httpConfig,omitempty"`
}

// TelegramConfigApplyConfiguration constructs a declarative configuration of the TelegramConfig type for use with
// apply.
func TelegramConfig() *TelegramConfigApplyConfiguration {
	return &TelegramConfigApplyConfiguration{}
}

// WithSendResolved sets the SendResolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SendResolved field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithSendResolved(value bool) *TelegramConfigApplyConfiguration {
	b.SendResolved = &value
	return b
}

// WithAPIURL sets the APIURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIURL field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithAPIURL(value string) *TelegramConfigApplyConfiguration {
	b.APIURL = &value
	return b
}

// WithBotToken sets the BotToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BotToken field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithBotToken(value *SecretKeySelectorApplyConfiguration) *TelegramConfigApplyConfiguration {
	b.BotToken = value
	return b
}

// WithBotTokenFile sets the BotTokenFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BotTokenFile field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithBotTokenFile(value string) *TelegramConfigApplyConfiguration {
	b.BotTokenFile = &value
	return b
}

// WithChatID sets the ChatID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ChatID field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithChatID(value int64) *TelegramConfigApplyConfiguration {
	b.ChatID = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithMessage(value string) *TelegramConfigApplyConfiguration {
	b.Message = &value
	return b
}

// WithDisableNotifications sets the DisableNotifications field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableNotifications field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithDisableNotifications(value bool) *TelegramConfigApplyConfiguration {
	b.DisableNotifications = &value
	return b
}

// WithParseMode sets the ParseMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ParseMode field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithParseMode(value string) *TelegramConfigApplyConfiguration {
	b.ParseMode = &value
	return b
}

// WithHTTPConfig sets the HTTPConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPConfig field is set to the value of the last call.
func (b *TelegramConfigApplyConfiguration) WithHTTPConfig(value *HTTPConfigApplyConfiguration) *TelegramConfigApplyConfiguration {
	b.HTTPConfig = value
	return b
}
