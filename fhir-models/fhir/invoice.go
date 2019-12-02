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

// Invoice is documented here http://hl7.org/fhir/StructureDefinition/Invoice
type Invoice struct {
	Id                  *string                         `json:"id,omitempty"`
	Meta                *Meta                           `json:"meta,omitempty"`
	ImplicitRules       *string                         `json:"implicitRules,omitempty"`
	Language            *string                         `json:"language,omitempty"`
	Text                *Narrative                      `json:"text,omitempty"`
	Extension           []Extension                     `json:"extension,omitempty"`
	ModifierExtension   []Extension                     `json:"modifierExtension,omitempty"`
	Identifier          []Identifier                    `json:"identifier,omitempty"`
	Status              InvoiceStatus                   `json:"status"`
	CancelledReason     *string                         `json:"cancelledReason,omitempty"`
	Type                *CodeableConcept                `json:"type,omitempty"`
	Subject             *Reference                      `json:"subject,omitempty"`
	Recipient           *Reference                      `json:"recipient,omitempty"`
	Date                *string                         `json:"date,omitempty"`
	Participant         []InvoiceParticipant            `json:"participant,omitempty"`
	Issuer              *Reference                      `json:"issuer,omitempty"`
	Account             *Reference                      `json:"account,omitempty"`
	LineItem            []InvoiceLineItem               `json:"lineItem,omitempty"`
	TotalPriceComponent []InvoiceLineItemPriceComponent `json:"totalPriceComponent,omitempty"`
	TotalNet            *Money                          `json:"totalNet,omitempty"`
	TotalGross          *Money                          `json:"totalGross,omitempty"`
	PaymentTerms        *string                         `json:"paymentTerms,omitempty"`
	Note                []Annotation                    `json:"note,omitempty"`
}
type InvoiceParticipant struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Actor             Reference        `json:"actor"`
}
type InvoiceLineItem struct {
	Id                *string                         `json:"id,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Sequence          *int                            `json:"sequence,omitempty"`
	PriceComponent    []InvoiceLineItemPriceComponent `json:"priceComponent,omitempty"`
}
type InvoiceLineItemPriceComponent struct {
	Id                *string                   `json:"id,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Type              InvoicePriceComponentType `json:"type"`
	Code              *CodeableConcept          `json:"code,omitempty"`
	Factor            *string                   `json:"factor,omitempty"`
	Amount            *Money                    `json:"amount,omitempty"`
}
type OtherInvoice Invoice

// MarshalJSON marshals the given Invoice as JSON into a byte slice
func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType"`
	}{
		OtherInvoice: OtherInvoice(r),
		ResourceType: "Invoice",
	})
}

// UnmarshalInvoice unmarshals a Invoice.
func UnmarshalInvoice(b []byte) (Invoice, error) {
	var invoice Invoice
	if err := json.Unmarshal(b, &invoice); err != nil {
		return invoice, err
	}
	return invoice, nil
}
