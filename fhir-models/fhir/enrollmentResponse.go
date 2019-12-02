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

// EnrollmentResponse is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
	Id                *string                       `json:"id,omitempty"`
	Meta              *Meta                         `json:"meta,omitempty"`
	ImplicitRules     *string                       `json:"implicitRules,omitempty"`
	Language          *string                       `json:"language,omitempty"`
	Text              *Narrative                    `json:"text,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                  `json:"identifier,omitempty"`
	Status            *FinancialResourceStatusCodes `json:"status,omitempty"`
	Request           *Reference                    `json:"request,omitempty"`
	Outcome           *ClaimProcessingCodes         `json:"outcome,omitempty"`
	Disposition       *string                       `json:"disposition,omitempty"`
	Created           *string                       `json:"created,omitempty"`
	Organization      *Reference                    `json:"organization,omitempty"`
	RequestProvider   *Reference                    `json:"requestProvider,omitempty"`
}
type OtherEnrollmentResponse EnrollmentResponse

// MarshalJSON marshals the given EnrollmentResponse as JSON into a byte slice
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentResponse: OtherEnrollmentResponse(r),
		ResourceType:            "EnrollmentResponse",
	})
}

// UnmarshalEnrollmentResponse unmarshals a EnrollmentResponse.
func UnmarshalEnrollmentResponse(b []byte) (EnrollmentResponse, error) {
	var enrollmentResponse EnrollmentResponse
	if err := json.Unmarshal(b, &enrollmentResponse); err != nil {
		return enrollmentResponse, err
	}
	return enrollmentResponse, nil
}
