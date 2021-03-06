// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
	"unicode/utf8"
)

// BeneficiaryCustomer is the beneficiary customer
type BeneficiaryCustomer struct {
	// tag
	tag string
	// CoverPayment is CoverPayment
	CoverPayment CoverPayment `json:"coverPayment,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryCustomer returns a new BeneficiaryCustomer
func NewBeneficiaryCustomer() *BeneficiaryCustomer {
	bc := &BeneficiaryCustomer{
		tag: TagBeneficiaryCustomer,
	}
	return bc
}

// Parse takes the input string and parses the BeneficiaryCustomer values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bc *BeneficiaryCustomer) Parse(record string) error {
	if utf8.RuneCountInString(record) != 186 {
		return NewTagWrongLengthErr(186, len(record))
	}
	bc.tag = record[:6]
	bc.CoverPayment.SwiftFieldTag = bc.parseStringField(record[6:11])
	bc.CoverPayment.SwiftLineOne = bc.parseStringField(record[11:46])
	bc.CoverPayment.SwiftLineTwo = bc.parseStringField(record[46:81])
	bc.CoverPayment.SwiftLineThree = bc.parseStringField(record[81:116])
	bc.CoverPayment.SwiftLineFour = bc.parseStringField(record[116:151])
	bc.CoverPayment.SwiftLineFive = bc.parseStringField(record[151:186])
	return nil
}

// String writes BeneficiaryCustomer
func (bc *BeneficiaryCustomer) String() string {
	var buf strings.Builder
	buf.Grow(186)
	buf.WriteString(bc.tag)
	buf.WriteString(bc.SwiftFieldTagField())
	buf.WriteString(bc.SwiftLineOneField())
	buf.WriteString(bc.SwiftLineTwoField())
	buf.WriteString(bc.SwiftLineThreeField())
	buf.WriteString(bc.SwiftLineFourField())
	buf.WriteString(bc.SwiftLineFiveField())
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryCustomer and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BeneficiaryCustomer) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	if bc.tag != TagBeneficiaryCustomer {
		return fieldError("tag", ErrValidTagForType, bc.tag)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, bc.CoverPayment.SwiftFieldTag)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineOne); err != nil {
		return fieldError("SwiftLineOne", err, bc.CoverPayment.SwiftLineOne)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineTwo); err != nil {
		return fieldError("SwiftLineTwo", err, bc.CoverPayment.SwiftLineTwo)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineThree); err != nil {
		return fieldError("SwiftLineThree", err, bc.CoverPayment.SwiftLineThree)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineFour); err != nil {
		return fieldError("SwiftLineFour", err, bc.CoverPayment.SwiftLineFour)
	}
	if err := bc.isAlphanumeric(bc.CoverPayment.SwiftLineFive); err != nil {
		return fieldError("SwiftLineFive", err, bc.CoverPayment.SwiftLineFive)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bc *BeneficiaryCustomer) fieldInclusion() error {
	if bc.CoverPayment.SwiftLineSix != "" {
		return fieldError("SwiftLineSix", ErrInvalidProperty, bc.CoverPayment.SwiftLineSix)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (bc *BeneficiaryCustomer) SwiftFieldTagField() string {
	return bc.alphaField(bc.CoverPayment.SwiftFieldTag, 5)
}

// SwiftLineOneField gets a string of the SwiftLineOne field
func (bc *BeneficiaryCustomer) SwiftLineOneField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineOne, 35)
}

// SwiftLineTwoField gets a string of the SwiftLineTwo field
func (bc *BeneficiaryCustomer) SwiftLineTwoField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineTwo, 35)
}

// SwiftLineThreeField gets a string of the SwiftLineThree field
func (bc *BeneficiaryCustomer) SwiftLineThreeField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineThree, 35)
}

// SwiftLineFourField gets a string of the SwiftLineFour field
func (bc *BeneficiaryCustomer) SwiftLineFourField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineFour, 35)
}

// SwiftLineFiveField gets a string of the SwiftLineFive field
func (bc *BeneficiaryCustomer) SwiftLineFiveField() string {
	return bc.alphaField(bc.CoverPayment.SwiftLineFive, 35)
}
