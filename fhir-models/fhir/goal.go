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

// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
	Id                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Identifier        []Identifier        `json:"identifier,omitempty"`
	LifecycleStatus   GoalLifecycleStatus `json:"lifecycleStatus"`
	AchievementStatus *CodeableConcept    `json:"achievementStatus,omitempty"`
	Category          []CodeableConcept   `json:"category,omitempty"`
	Priority          *CodeableConcept    `json:"priority,omitempty"`
	Description       CodeableConcept     `json:"description"`
	Subject           Reference           `json:"subject"`
	Target            []GoalTarget        `json:"target,omitempty"`
	StatusDate        *string             `json:"statusDate,omitempty"`
	StatusReason      *string             `json:"statusReason,omitempty"`
	ExpressedBy       *Reference          `json:"expressedBy,omitempty"`
	Addresses         []Reference         `json:"addresses,omitempty"`
	Note              []Annotation        `json:"note,omitempty"`
	OutcomeCode       []CodeableConcept   `json:"outcomeCode,omitempty"`
	OutcomeReference  []Reference         `json:"outcomeReference,omitempty"`
}
type GoalTarget struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Measure           *CodeableConcept `json:"measure,omitempty"`
}
type OtherGoal Goal

// MarshalJSON marshals the given Goal as JSON into a byte slice
func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}

// UnmarshalGoal unmarshals a Goal.
func UnmarshalGoal(b []byte) (Goal, error) {
	var goal Goal
	if err := json.Unmarshal(b, &goal); err != nil {
		return goal, err
	}
	return goal, nil
}
