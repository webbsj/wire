/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// FinancialInstitution - FinancialInstitution is financial institution demographic information
type FinancialInstitution struct {

	// Identification Code:  * `B` - SWIFT Bank Identifier Code (BIC) * `C` - CHIPS Participant * `D` - Demand Deposit Account (DDA) Number * `F` - Fed Routing Number * `T` - SWIFT BIC or Bank Entity Identifier (BEI) and Account Number * `U` - CHIPS Identifier 
	IdentificationCode string `json:"identificationCode"`

	// Identifier
	Identifier string `json:"identifier"`

	// Name
	Name string `json:"name"`

	Address Address `json:"address"`
}
