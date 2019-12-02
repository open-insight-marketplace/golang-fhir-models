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

// SubstanceSpecification is documented here http://hl7.org/fhir/StructureDefinition/SubstanceSpecification
type SubstanceSpecification struct {
	Id                   *string                                                 `json:"id,omitempty"`
	Meta                 *Meta                                                   `json:"meta,omitempty"`
	ImplicitRules        *string                                                 `json:"implicitRules,omitempty"`
	Language             *string                                                 `json:"language,omitempty"`
	Text                 *Narrative                                              `json:"text,omitempty"`
	Extension            []Extension                                             `json:"extension,omitempty"`
	ModifierExtension    []Extension                                             `json:"modifierExtension,omitempty"`
	Identifier           *Identifier                                             `json:"identifier,omitempty"`
	Type                 *CodeableConcept                                        `json:"type,omitempty"`
	Status               *CodeableConcept                                        `json:"status,omitempty"`
	Domain               *CodeableConcept                                        `json:"domain,omitempty"`
	Description          *string                                                 `json:"description,omitempty"`
	Source               []Reference                                             `json:"source,omitempty"`
	Comment              *string                                                 `json:"comment,omitempty"`
	Moiety               []SubstanceSpecificationMoiety                          `json:"moiety,omitempty"`
	Property             []SubstanceSpecificationProperty                        `json:"property,omitempty"`
	ReferenceInformation *Reference                                              `json:"referenceInformation,omitempty"`
	Structure            *SubstanceSpecificationStructure                        `json:"structure,omitempty"`
	Code                 []SubstanceSpecificationCode                            `json:"code,omitempty"`
	Name                 []SubstanceSpecificationName                            `json:"name,omitempty"`
	MolecularWeight      []SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
	Relationship         []SubstanceSpecificationRelationship                    `json:"relationship,omitempty"`
	NucleicAcid          *Reference                                              `json:"nucleicAcid,omitempty"`
	Polymer              *Reference                                              `json:"polymer,omitempty"`
	Protein              *Reference                                              `json:"protein,omitempty"`
	SourceMaterial       *Reference                                              `json:"sourceMaterial,omitempty"`
}
type SubstanceSpecificationMoiety struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Identifier        *Identifier      `json:"identifier,omitempty"`
	Name              *string          `json:"name,omitempty"`
	Stereochemistry   *CodeableConcept `json:"stereochemistry,omitempty"`
	OpticalActivity   *CodeableConcept `json:"opticalActivity,omitempty"`
	MolecularFormula  *string          `json:"molecularFormula,omitempty"`
}
type SubstanceSpecificationProperty struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `json:"category,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Parameters        *string          `json:"parameters,omitempty"`
}
type SubstanceSpecificationStructure struct {
	Id                       *string                                                `json:"id,omitempty"`
	Extension                []Extension                                            `json:"extension,omitempty"`
	ModifierExtension        []Extension                                            `json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                                       `json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                                       `json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                                `json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                                `json:"molecularFormulaByMoiety,omitempty"`
	Isotope                  []SubstanceSpecificationStructureIsotope               `json:"isotope,omitempty"`
	MolecularWeight          *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
	Source                   []Reference                                            `json:"source,omitempty"`
	Representation           []SubstanceSpecificationStructureRepresentation        `json:"representation,omitempty"`
}
type SubstanceSpecificationStructureIsotope struct {
	Id                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Identifier        *Identifier                                            `json:"identifier,omitempty"`
	Name              *CodeableConcept                                       `json:"name,omitempty"`
	Substitution      *CodeableConcept                                       `json:"substitution,omitempty"`
	HalfLife          *Quantity                                              `json:"halfLife,omitempty"`
	MolecularWeight   *SubstanceSpecificationStructureIsotopeMolecularWeight `json:"molecularWeight,omitempty"`
}
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `json:"method,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}
type SubstanceSpecificationStructureRepresentation struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Representation    *string          `json:"representation,omitempty"`
	Attachment        *Attachment      `json:"attachment,omitempty"`
}
type SubstanceSpecificationCode struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	StatusDate        *string          `json:"statusDate,omitempty"`
	Comment           *string          `json:"comment,omitempty"`
	Source            []Reference      `json:"source,omitempty"`
}
type SubstanceSpecificationName struct {
	Id                *string                              `json:"id,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Name              string                               `json:"name"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	Status            *CodeableConcept                     `json:"status,omitempty"`
	Preferred         *bool                                `json:"preferred,omitempty"`
	Language          []CodeableConcept                    `json:"language,omitempty"`
	Domain            []CodeableConcept                    `json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                    `json:"jurisdiction,omitempty"`
	Synonym           []SubstanceSpecificationName         `json:"synonym,omitempty"`
	Translation       []SubstanceSpecificationName         `json:"translation,omitempty"`
	Official          []SubstanceSpecificationNameOfficial `json:"official,omitempty"`
	Source            []Reference                          `json:"source,omitempty"`
}
type SubstanceSpecificationNameOfficial struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `json:"authority,omitempty"`
	Status            *CodeableConcept `json:"status,omitempty"`
	Date              *string          `json:"date,omitempty"`
}
type SubstanceSpecificationRelationship struct {
	Id                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	Relationship        *CodeableConcept `json:"relationship,omitempty"`
	IsDefining          *bool            `json:"isDefining,omitempty"`
	AmountRatioLowLimit *Ratio           `json:"amountRatioLowLimit,omitempty"`
	AmountType          *CodeableConcept `json:"amountType,omitempty"`
	Source              []Reference      `json:"source,omitempty"`
}
type OtherSubstanceSpecification SubstanceSpecification

// MarshalJSON marshals the given SubstanceSpecification as JSON into a byte slice
func (r SubstanceSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceSpecification
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceSpecification: OtherSubstanceSpecification(r),
		ResourceType:                "SubstanceSpecification",
	})
}

// UnmarshalSubstanceSpecification unmarshals a SubstanceSpecification.
func UnmarshalSubstanceSpecification(b []byte) (SubstanceSpecification, error) {
	var substanceSpecification SubstanceSpecification
	if err := json.Unmarshal(b, &substanceSpecification); err != nil {
		return substanceSpecification, err
	}
	return substanceSpecification, nil
}
