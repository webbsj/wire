package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockBeneficiaryFI creates a BeneficiaryFI
func mockBeneficiaryFI() *BeneficiaryFI {
	bfi := NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return bfi
}

// TestMockBeneficiaryFI validates mockBeneficiaryFI
func TestMockBeneficiaryFI(t *testing.T) {
	bfi := mockBeneficiaryFI()
	if err := bfi.Validate(); err != nil {
		t.Error("mockBeneficiaryFI does not validate and will break other tests")
	}
}

// TestBeneficiaryFIIdentificationCodeValid validates BeneficiaryFI IdentificationCode
func TestBeneficiaryFIIdentificationCodeValid(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = "Football Card ID"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIIdentificationCodeFI validates BeneficiaryFI IdentificationCode is an FI code
func TestBeneficiaryFIIdentificationCodeFI(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = "1"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIIdentifierAlphaNumeric validates BeneficiaryFI Identifier is alphanumeric
func TestBeneficiaryFIIdentifierAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Identifier = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFINameAlphaNumeric validates BeneficiaryFI Name is alphanumeric
func TestBeneficiaryFINameAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Name = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIAddressLineOneAlphaNumeric validates BeneficiaryFI AddressLineOne is alphanumeric
func TestBeneficiaryFIAddressLineOneAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Address.AddressLineOne = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIAddressLineTwoAlphaNumeric validates BeneficiaryFI AddressLineTwo is alphanumeric
func TestBeneficiaryFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Address.AddressLineTwo = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIAddressLineThreeAlphaNumeric validates BeneficiaryFI AddressLineThree is alphanumeric
func TestBeneficiaryFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Address.AddressLineThree = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIIdentificationCodeRequired validates BeneficiaryFI IdentificationCode is required
func TestBeneficiaryFIIdentificationCodeRequired(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = ""
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryFIIdentifierRequired validates BeneficiaryFI Identifier is required
func TestBeneficiaryFIIdentifierRequired(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.FinancialInstitution.Identifier = ""
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryFIWrongLength parses a wrong BeneficiaryFI record length
func TestParseBeneficiaryFIWrongLength(t *testing.T) {
	var line = "{4100}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	err := r.parseBeneficiaryFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryFIReaderParseError parses a wrong BeneficiaryFI reader parse error
func TestParseBeneficiaryFIReaderParseError(t *testing.T) {
	var line = "{4100}D123456789                         F® Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	err := r.parseBeneficiaryFI()
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

// TestBeneficiaryFITagError validates a BeneficiaryFI tag
func TestBeneficiaryFITagError(t *testing.T) {
	bfi := mockBeneficiaryFI()
	bfi.tag = "{9999}"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
