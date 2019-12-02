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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// TriggerDefinition is documented here http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	Id        *string           `json:"id,omitempty"`
	Extension []Extension       `json:"extension,omitempty"`
	Type      TriggerType       `json:"type"`
	Name      *string           `json:"name,omitempty"`
	Data      []DataRequirement `json:"data,omitempty"`
	Condition *Expression       `json:"condition,omitempty"`
}
