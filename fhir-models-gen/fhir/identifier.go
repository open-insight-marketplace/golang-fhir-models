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

// Identifier is documented here http://hl7.org/fhir/StructureDefinition/Identifier
type Identifier struct {
	Id        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Use       *IdentifierUse   `json:"use,omitempty"`
	Type      *CodeableConcept `json:"type,omitempty"`
	System    *string          `json:"system,omitempty"`
	Value     *string          `json:"value,omitempty"`
	Period    *Period          `json:"period,omitempty"`
	Assigner  *Reference       `json:"assigner,omitempty"`
}
