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

// Resource is documented here http://hl7.org/fhir/StructureDefinition/Resource
type Resource struct {
	Id            *string `json:"id,omitempty"`
	Meta          *Meta   `json:"meta,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Language      *string `json:"language,omitempty"`
}
type OtherResource Resource

// MarshalJSON marshals the given Resource as JSON into a byte slice
func (r Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResource
		ResourceType string `json:"resourceType"`
	}{
		OtherResource: OtherResource(r),
		ResourceType:  "Resource",
	})
}

// UnmarshalResource unmarshals a Resource.
func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}
