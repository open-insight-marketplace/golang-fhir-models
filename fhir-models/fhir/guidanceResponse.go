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

// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
	Id                 *string                `json:"id,omitempty"`
	Meta               *Meta                  `json:"meta,omitempty"`
	ImplicitRules      *string                `json:"implicitRules,omitempty"`
	Language           *string                `json:"language,omitempty"`
	Text               *Narrative             `json:"text,omitempty"`
	Extension          []Extension            `json:"extension,omitempty"`
	ModifierExtension  []Extension            `json:"modifierExtension,omitempty"`
	RequestIdentifier  *Identifier            `json:"requestIdentifier,omitempty"`
	Identifier         []Identifier           `json:"identifier,omitempty"`
	Status             GuidanceResponseStatus `json:"status"`
	Subject            *Reference             `json:"subject,omitempty"`
	Encounter          *Reference             `json:"encounter,omitempty"`
	OccurrenceDateTime *string                `json:"occurrenceDateTime,omitempty"`
	Performer          *Reference             `json:"performer,omitempty"`
	ReasonCode         []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference    []Reference            `json:"reasonReference,omitempty"`
	Note               []Annotation           `json:"note,omitempty"`
	EvaluationMessage  []Reference            `json:"evaluationMessage,omitempty"`
	OutputParameters   *Reference             `json:"outputParameters,omitempty"`
	Result             *Reference             `json:"result,omitempty"`
	DataRequirement    []DataRequirement      `json:"dataRequirement,omitempty"`
}
type OtherGuidanceResponse GuidanceResponse

// MarshalJSON marshals the given GuidanceResponse as JSON into a byte slice
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}

// UnmarshalGuidanceResponse unmarshals a GuidanceResponse.
func UnmarshalGuidanceResponse(b []byte) (GuidanceResponse, error) {
	var guidanceResponse GuidanceResponse
	if err := json.Unmarshal(b, &guidanceResponse); err != nil {
		return guidanceResponse, err
	}
	return guidanceResponse, nil
}
