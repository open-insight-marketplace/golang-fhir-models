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

// BiologicallyDerivedProduct is documented here http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	Id                *string                                 `json:"id,omitempty"`
	Meta              *Meta                                   `json:"meta,omitempty"`
	ImplicitRules     *string                                 `json:"implicitRules,omitempty"`
	Language          *string                                 `json:"language,omitempty"`
	Text              *Narrative                              `json:"text,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                            `json:"identifier,omitempty"`
	ProductCategory   *BiologicallyDerivedProductCategory     `json:"productCategory,omitempty"`
	ProductCode       *CodeableConcept                        `json:"productCode,omitempty"`
	Status            *BiologicallyDerivedProductStatus       `json:"status,omitempty"`
	Request           []Reference                             `json:"request,omitempty"`
	Quantity          *int                                    `json:"quantity,omitempty"`
	Parent            []Reference                             `json:"parent,omitempty"`
	Collection        *BiologicallyDerivedProductCollection   `json:"collection,omitempty"`
	Processing        []BiologicallyDerivedProductProcessing  `json:"processing,omitempty"`
	Manipulation      *BiologicallyDerivedProductManipulation `json:"manipulation,omitempty"`
	Storage           []BiologicallyDerivedProductStorage     `json:"storage,omitempty"`
}
type BiologicallyDerivedProductCollection struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Collector         *Reference  `json:"collector,omitempty"`
	Source            *Reference  `json:"source,omitempty"`
}
type BiologicallyDerivedProductProcessing struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          *Reference       `json:"additive,omitempty"`
}
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
}
type BiologicallyDerivedProductStorage struct {
	Id                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Description       *string                                 `json:"description,omitempty"`
	Temperature       *string                                 `json:"temperature,omitempty"`
	Scale             *BiologicallyDerivedProductStorageScale `json:"scale,omitempty"`
	Duration          *Period                                 `json:"duration,omitempty"`
}
type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

// MarshalJSON marshals the given BiologicallyDerivedProduct as JSON into a byte slice
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(r),
		ResourceType:                    "BiologicallyDerivedProduct",
	})
}

// UnmarshalBiologicallyDerivedProduct unmarshals a BiologicallyDerivedProduct.
func UnmarshalBiologicallyDerivedProduct(b []byte) (BiologicallyDerivedProduct, error) {
	var biologicallyDerivedProduct BiologicallyDerivedProduct
	if err := json.Unmarshal(b, &biologicallyDerivedProduct); err != nil {
		return biologicallyDerivedProduct, err
	}
	return biologicallyDerivedProduct, nil
}
