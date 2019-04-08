// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FIIntermediaryFIAdvice is the financial institution intermediary financial institution
type FIIntermediaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitEmpty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIIntermediaryFIAdvice returns a new FIIntermediaryFIAdvice
func NewFIIntermediaryFIAdvice() FIIntermediaryFIAdvice {
	fiifia := FIIntermediaryFIAdvice{
		tag: TagFIIntermediaryFIAdvice,
	}
	return fiifia
}

// Parse takes the input string and parses the FIIntermediaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fiifia *FIIntermediaryFIAdvice) Parse(record string) {
	fiifia.tag = record[:6]
	fiifia.Advice.AdviceCode = fiifia.parseStringField(record[6:9])
	fiifia.Advice.LineOne = fiifia.parseStringField(record[9:35])
	fiifia.Advice.LineTwo = fiifia.parseStringField(record[35:68])
	fiifia.Advice.LineThree = fiifia.parseStringField(record[68:101])
	fiifia.Advice.LineFour = fiifia.parseStringField(record[101:134])
	fiifia.Advice.LineFive = fiifia.parseStringField(record[134:167])
	fiifia.Advice.LineSix = fiifia.parseStringField(record[167:200])
}

// String writes FIIntermediaryFIAdvice
func (fiifia *FIIntermediaryFIAdvice) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(200)
	buf.WriteString(fiifia.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on FIIntermediaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fiifia *FIIntermediaryFIAdvice) Validate() error {
	if err := fiifia.fieldInclusion(); err != nil {
		return err
	}
	if err := fiifia.isAdviceCode(fiifia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fiifia.Advice.AdviceCode)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fiifia.Advice.LineOne)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fiifia.Advice.LineTwo)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fiifia.Advice.LineThree)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fiifia.Advice.LineFour)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fiifia.Advice.LineFive)
	}
	if err := fiifia.isAlphanumeric(fiifia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fiifia.Advice.LineSix)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (fiifia *FIIntermediaryFIAdvice) fieldInclusion() error {
	return nil
}
