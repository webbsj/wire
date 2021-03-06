package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockBeneficiary creates a Beneficiary
func mockBeneficiary() *Beneficiary {
	ben := NewBeneficiary()
	ben.Personal.IdentificationCode = DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	return ben
}

// TestMockBeneficiary validates mockBeneficiary
func TestMockBeneficiary(t *testing.T) {
	ben := mockBeneficiary()
	if err := ben.Validate(); err != nil {
		t.Error("mockBeneficiary does not validate and will break other tests")
	}
}

// TestBeneficiaryIdentificationCodeValid validates Beneficiary IdentificationCode
func TestBeneficiaryIdentificationCodeValid(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = "Baseball Card ID"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIdentifierAlphaNumeric validates Beneficiary Identifier is alphanumeric
func TestBeneficiaryIdentifierAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Identifier = "®"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryNameAlphaNumeric validates Beneficiary Name is alphanumeric
func TestBeneficiaryNameAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Name = "®"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryAddressLineOneAlphaNumeric validates Beneficiary AddressLineOne is alphanumeric
func TestBeneficiaryAddressLineOneAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Address.AddressLineOne = "®"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryAddressLineTwoAlphaNumeric validates Beneficiary AddressLineTwo is alphanumeric
func TestBeneficiaryAddressLineTwoAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Address.AddressLineTwo = "®"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryAddressLineThreeAlphaNumeric validates Beneficiary AddressLineThree is alphanumeric
func TestBeneficiaryAddressLineThreeAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Address.AddressLineThree = "®"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIdentificationCodeRequired validates Beneficiary IdentificationCode is required
func TestBeneficiaryIdentificationCodeRequired(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = ""
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIdentifierRequired validates Beneficiary Identifier is required
func TestBeneficiaryIdentifierRequired(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Identifier = ""
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryWrongLength parses a wrong Beneficiary record length
func TestParseBeneficiaryWrongLength(t *testing.T) {
	var line = "{4200}31234                              Name                               Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	err := r.parseBeneficiary()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryReaderParseError parses a wrong Beneficiary reader parse error
func TestParseBeneficiaryReaderParseError(t *testing.T) {
	var line = "{4200}31234                              Na®e                               Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	err := r.parseBeneficiary()
	if err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryTagError validates Beneficiary tag
func TestBeneficiaryTagError(t *testing.T) {
	ben := mockBeneficiary()
	ben.tag = "{9999}"
	if err := ben.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
