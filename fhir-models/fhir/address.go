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

// Address is documented here http://hl7.org/fhir/StructureDefinition/Address
type Address struct {
	Id         *string      `json:"id,omitempty"`
	Extension  []Extension  `json:"extension,omitempty"`
	Use        *AddressUse  `json:"use,omitempty"`
	Type       *AddressType `json:"type,omitempty"`
	Text       *string      `json:"text,omitempty"`
	Line       []string     `json:"line,omitempty"`
	City       *string      `json:"city,omitempty"`
	District   *string      `json:"district,omitempty"`
	State      *string      `json:"state,omitempty"`
	PostalCode *string      `json:"postalCode,omitempty"`
	Country    *string      `json:"country,omitempty"`
	Period     *Period      `json:"period,omitempty"`
}
