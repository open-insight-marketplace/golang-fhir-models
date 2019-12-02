// Copyright 2019 The Samply Development Community
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

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Subscription is documented here http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
	Id                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Status            SubscriptionStatus  `json:"status"`
	Contact           []ContactPoint      `json:"contact,omitempty"`
	End               *string             `json:"end,omitempty"`
	Reason            string              `json:"reason"`
	Criteria          string              `json:"criteria"`
	Error             *string             `json:"error,omitempty"`
	Channel           SubscriptionChannel `json:"channel"`
}
type SubscriptionChannel struct {
	Id                *string                 `json:"id,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Type              SubscriptionChannelType `json:"type"`
	Endpoint          *string                 `json:"endpoint,omitempty"`
	Payload           *string                 `json:"payload,omitempty"`
	Header            []string                `json:"header,omitempty"`
}
type OtherSubscription Subscription

// MarshalJSON marshals the given Subscription as JSON into a byte slice
func (r Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType"`
	}{
		OtherSubscription: OtherSubscription(r),
		ResourceType:      "Subscription",
	})
}

// UnmarshalSubscription unmarshals a Subscription.
func UnmarshalSubscription(b []byte) (Subscription, error) {
	var subscription Subscription
	if err := json.Unmarshal(b, &subscription); err != nil {
		return subscription, err
	}
	return subscription, nil
}
