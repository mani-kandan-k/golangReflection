package validatefields

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// NewRecord represents the data structure for records with the new field descriptions
type Eightline struct {
	LineNumber             int     `validate:"min=7,max=99"`
	PurposeCode            int     `validate:"max=99"`
	Name                   string  `validate:"omitempty,max=100"`
	MiddleName             string  `validate:"required_if_oneof=Name luffy zoro,max=20"`
	LastSearchName         string  `validate:"omitempty,max=20"`
	Title                  string  `validate:"omitempty,max=10"`
	Suffix                 string  `validate:"omitempty,max=10"`
	FatherHusbandName      string  `validate:"omitempty,max=50"`
	Address1               string  `validate:"max=55"`
	Address2               string  `validate:"omitempty,max=55"`
	Address3               string  `validate:"omitempty,max=55"`
	CountryCode            string  `validate:"omitempty,iso3166_1_alpha2,max=2"`
	ZipCode                string  `validate:"omitempty,max=10"`
	StateCode              string  `validate:"required_if=CountryCode IN,max=6"` // ,$$
	State                  string  `validate:"omitempty,max=25"`
	City                   string  `validate:"omitempty,max=25"`
	CitySequenceNo         int     `validate:"required_if=CountryCode IN,max=99"`
	MobileTelephoneISDCode string  `validate:"required_with=MobileTelephoneNo,max=6"`
	MobileTelephoneNo      string  `validate:"required_with=MobileTelephoneISDCode,max=17"`
	DateOfBirth            string  `validate:"omitempty,datetime=2006-01-02"`
	Fax                    string  `validate:"omitempty,max=17"`
	IncomeTaxPan           string  `validate:"omitempty,max=10"`
	UID                    string  `validate:"omitempty,max=16"`
	UIDVerificationFlag    string  `validate:"omitempty,max=1"`
	NameChangeReasonCode   string  `validate:"omitempty,max=2"`
	ITCircle               string  `validate:"omitempty,max=15"`
	PrimaryEmail           string  `validate:"omitempty,email,max=100"`
	UserText1              string  `validate:"omitempty,max=50"`
	UserText2              string  `validate:"omitempty,max=50"`
	UserField3             int     `validate:"omitempty,max=9999"`
	UserField4             string  `validate:"omitempty,max=4"`
	UserField5             int     `validate:"omitempty,max=9999"`
	NomineeSerialNumber    int     `validate:"required_if_oneof=PurposeCode 8 10,max=99"` // $$ oneof=1 2 3
	RelationshipWithBO     int     `validate:"required_if=PurposeCode 6,oneof=1 2 3 4 5 6 7 8 9 10 11 12 13,max=99"`
	PercentageOfShares     float64 `validate:"required_if=PurposeCode 6,max=99999"` // $$
	ResidualSecuritiesFlag string  `validate:"required_if=PurposeCode 6,max=1"`
	Filler1                string  `validate:"omitempty,max=16"`
	Filler2                string  `validate:"omitempty,max=72"`
	Filler3                string  `validate:"omitempty,max=1"`
	Filler4                string  `validate:"omitempty,max=1"`
	Filler5                string  `validate:"omitempty,max=10"`
}

// NomineeAddStrConstruction constructs a record string based on the provided NewRecord
func NomineeAddStrConstruction(pEightline Eightline) string {

	return fmt.Sprintf("%02d%02d%-100s%-20s%-20s%-10s%-10s%-50s%-55s%-55s%-55s%-2s%-10s%-6s%-25s%-25s%02d%-6s%-17s%-8s%-17s%-10s%-16s%-1s%-2s%-15s%-100s%-50s%-50s%04d%-4s%04d%02d%02d%5.2f%-1s%-16s%-72s%-1s%-1s%-10s",
		pEightline.LineNumber,
		pEightline.PurposeCode,
		pEightline.Name,
		pEightline.MiddleName,
		pEightline.LastSearchName,
		pEightline.Title,
		pEightline.Suffix,
		pEightline.FatherHusbandName,
		pEightline.Address1,
		pEightline.Address2,
		pEightline.Address3,
		pEightline.CountryCode,
		pEightline.ZipCode,
		pEightline.StateCode,
		pEightline.State,
		pEightline.City,
		pEightline.CitySequenceNo,
		pEightline.MobileTelephoneISDCode,
		pEightline.MobileTelephoneNo,
		pEightline.DateOfBirth,
		pEightline.Fax,
		pEightline.IncomeTaxPan,
		pEightline.UID,
		pEightline.UIDVerificationFlag,
		pEightline.NameChangeReasonCode,
		pEightline.ITCircle,
		pEightline.PrimaryEmail,
		pEightline.UserText1,
		pEightline.UserText2,
		pEightline.UserField3,
		pEightline.UserField4,
		pEightline.UserField5,
		pEightline.NomineeSerialNumber,
		pEightline.RelationshipWithBO,
		pEightline.PercentageOfShares,
		pEightline.ResidualSecuritiesFlag,
		pEightline.Filler1,
		pEightline.Filler2,
		pEightline.Filler3,
		pEightline.Filler4,
		pEightline.Filler5,
	)
}

func AssignNomineeDetails() Eightline {

	var lEighthLineClientRec Eightline

	lEighthLineClientRec.LineNumber = 7
	lEighthLineClientRec.PurposeCode = 9
	lEighthLineClientRec.Name = "zoro"
	lEighthLineClientRec.MiddleName = "mmm"
	lEighthLineClientRec.LastSearchName = ""
	lEighthLineClientRec.Title = ""
	lEighthLineClientRec.Suffix = ""
	lEighthLineClientRec.FatherHusbandName = ""
	lEighthLineClientRec.Address1 = ""
	lEighthLineClientRec.Address2 = ""
	lEighthLineClientRec.Address3 = ""
	lEighthLineClientRec.CountryCode = "IN"
	lEighthLineClientRec.ZipCode = ""
	lEighthLineClientRec.StateCode = "TN"
	lEighthLineClientRec.State = ""
	lEighthLineClientRec.City = ""
	lEighthLineClientRec.CitySequenceNo = 6
	lEighthLineClientRec.MobileTelephoneISDCode = "+61"
	lEighthLineClientRec.MobileTelephoneNo = "12345"
	lEighthLineClientRec.DateOfBirth = "2024-01-01"
	lEighthLineClientRec.Fax = ""
	lEighthLineClientRec.IncomeTaxPan = ""
	lEighthLineClientRec.UID = ""
	lEighthLineClientRec.UIDVerificationFlag = ""
	lEighthLineClientRec.NameChangeReasonCode = ""
	lEighthLineClientRec.ITCircle = ""
	lEighthLineClientRec.PrimaryEmail = "1@1.a"
	lEighthLineClientRec.UserText1 = ""
	lEighthLineClientRec.UserText2 = ""
	lEighthLineClientRec.UserField3 = 0
	lEighthLineClientRec.UserField4 = ""
	lEighthLineClientRec.UserField5 = 0
	lEighthLineClientRec.NomineeSerialNumber = 0
	lEighthLineClientRec.RelationshipWithBO = 02
	lEighthLineClientRec.PercentageOfShares = 528.08
	lEighthLineClientRec.ResidualSecuritiesFlag = "N"
	lEighthLineClientRec.Filler1 = ""
	lEighthLineClientRec.Filler2 = ""
	lEighthLineClientRec.Filler3 = ""
	lEighthLineClientRec.Filler4 = ""
	lEighthLineClientRec.Filler5 = ""

	return lEighthLineClientRec
}

func ValidateNomineeDetails(pValidateReq Eightline) string {
	var lResult string

	if pValidateReq.LineNumber < 7 { // value should be 0 to 7, required
		lResult += "Line Number should be greater than 8 | "
	} else if (pValidateReq.LineNumber / 10) > 10 { // value len(no of digit) == 2
		lResult += "Line Number should be less than 2 digit | "
	}

	if pValidateReq.PurposeCode == 0 { // required --> != 0
		lResult += "Purpose Code is Mandatory | "
	} else if (pValidateReq.PurposeCode / 10) > 10 { // value len(no of digit) == 2
		lResult += "Purpose Code should be less than 2 digit | "
	}

	if pValidateReq.Name != "" {
		if len(pValidateReq.Name) > 100 { // min-len=100
			lResult += "BO Name Length should be less than 100 | "
		}
		if pValidateReq.NameChangeReasonCode == "" { // dependent required
			lResult += "Name Change Reason Code is Mandatory while changing the Name | "
		} else if len(pValidateReq.NameChangeReasonCode) > 2 { //
			lResult += "Name Change Reason Code length should be less than 2 | "
		}
	} else {
		if len(pValidateReq.NameChangeReasonCode) > 2 {
			lResult += "Name Change Reason Code length should be less than 2 | "
		}
	}

	if pValidateReq.MiddleName != "" {
		if len(pValidateReq.MiddleName) > 20 {
			lResult += "Middle Name length should be less than 20 | "
		}
	}

	if pValidateReq.LastSearchName != "" {
		if len(pValidateReq.LastSearchName) > 20 {
			lResult += "Last Search Name length should be less than 20 | "
		}
	}

	if pValidateReq.Title != "" {
		if len(pValidateReq.Title) > 10 {
			lResult += "Title length should be less than 10 | "
		}
	}

	if pValidateReq.Suffix != "" {
		if len(pValidateReq.Suffix) > 10 {
			lResult += "Suffix length should be less than 10 | "
		}
	}

	if pValidateReq.FatherHusbandName != "" {
		if len(pValidateReq.FatherHusbandName) > 50 {
			lResult += "Father / Husband Name length should be less than 50 | "
		}
	}

	if pValidateReq.Address1 != "" || pValidateReq.Address2 != "" || pValidateReq.Address3 != "" {
		if len(pValidateReq.Address1) > 55 || len(pValidateReq.Address2) > 55 || len(pValidateReq.Address3) > 55 {
			lResult += "Address length should be less than 55 | "
		}
	}

	// If Country Code is IN
	if pValidateReq.CountryCode == "IN" {
		if pValidateReq.StateCode == "" {
			lResult += "State Code is Mandatory when the Country Code is IN | "
		} else if len(pValidateReq.StateCode) > 6 {
			lResult += "State Code length should be less than 6 | "
		}
		// -------- Check this validation --------
		if pValidateReq.CitySequenceNo >= 1 {
			if (pValidateReq.CitySequenceNo / 10) > 10 {
				lResult += "City Sequence Number should be two digit | "
			}
		} else {
			lResult += "City Sequence Number is mandatory when the Country Code is IN | "
		}
	} else {
		if len(pValidateReq.CountryCode) > 2 {
			lResult += "Country Code length should be less than 2 | "
		}
		if len(pValidateReq.StateCode) > 6 {
			lResult += "State Code length should be less than 6 | "
		}
	}

	// ZIP Code
	if len(pValidateReq.ZipCode) > 10 {
		lResult += "ZIP Code less than 10 | "
	}

	// City
	if len(pValidateReq.City) > 25 {
		lResult += "City length should be less than 25 | "
	}

	if pValidateReq.State != "" {
		if len(pValidateReq.State) > 25 {
			lResult += "State length should less than 25 | "
		}
	}

	// Primary Mobile Number and Mobile ISD Code Validation
	PrimeMobileNumber := true
	if pValidateReq.MobileTelephoneNo != "" {
		PrimeMobileNumber = false
		if len(pValidateReq.MobileTelephoneNo) > 17 {
			lResult += "Secondary Mobile Number length Should be less than 17 | "

		}
		if pValidateReq.MobileTelephoneISDCode != "" {
			lResult += "Secondary Mobile ISD Code is Mandatory when Secondary Mobile Number is Present | "
		} else {
			if len(pValidateReq.MobileTelephoneISDCode) > 6 {
				lResult += "Secondary Mobile ISD Code is Should be less than 6 digit | "

			}
		}
	}
	if pValidateReq.MobileTelephoneISDCode != "" && PrimeMobileNumber {
		if len(pValidateReq.MobileTelephoneISDCode) > 6 {
			lResult += "Secondary Mobile ISD Code is Should be less than 6 digit | "
		}
		if pValidateReq.MobileTelephoneNo == "" {
			lResult += "Secondary Mobile Number is Mandatory when Secondary Mobile ISD Code is Present | "
		} else {
			if len(pValidateReq.MobileTelephoneNo) > 17 {
				lResult += "Secondary Mobile Number length Should be less than 17 | "
			}
		}
	}

	// Date of Birth Origin
	if pValidateReq.DateOfBirth != "" {
		if len(pValidateReq.DateOfBirth) > 8 {
			lResult += "Date of Birth length should be less than 8 | "
		}
	}

	// Fax Length
	if len(pValidateReq.Fax) > 17 {
		lResult += "Fax length Should be less 17 | "
	}

	// Income Tax Pan
	if len(pValidateReq.IncomeTaxPan) > 10 {
		lResult += "PAN Length should be less than 10 | "
	}

	// UID Length
	if len(pValidateReq.UID) > 16 {
		lResult += "UID Length should be less 16 | "
	}

	// UID Verification Flag Length
	if len(pValidateReq.UIDVerificationFlag) > 1 {
		lResult += "UID Verification Flag length Should be less 1 | "
	}

	// ITCircle length
	if len(pValidateReq.ITCircle) > 15 {
		lResult += "IT Circle length Should be less than 15 | "
	}

	// Primary Email
	if pValidateReq.PrimaryEmail != "" {
		if utf8.RuneCountInString(pValidateReq.PrimaryEmail) > 100 {
			lResult += "Primary Email Id should contain less than 100 Character | "
		}
		if EmailValidation(pValidateReq.PrimaryEmail) {
			lResult += "Invalid Primary Email Id | "
		}
	}

	if len(pValidateReq.UserText2) > 50 {
		lResult += "User Text 2 length should be less than 50 | "
	}

	if (pValidateReq.UserField3 / 1000) > 10 {
		lResult += "User Field 3 should be less than 4 digit | "
	}

	if pValidateReq.UserField4 == "" {
		lResult += "User Field 4 Field is Mandatory | "
	}

	if (pValidateReq.UserField5 / 1000) > 10 {
		lResult += "User Field 5 should be less 4 digit | "
	}

	if pValidateReq.PurposeCode == 6 || pValidateReq.PurposeCode == 8 {
		if pValidateReq.NomineeSerialNumber == 0 {
			lResult += "Nominee Serial Number is Mandatory when purpose code is 6 or purpose code 8 | "
		} else if pValidateReq.NomineeSerialNumber > 3 {
			lResult += "Nominee Serial Number Value should be less 3 | "
		}
	} else if (pValidateReq.NomineeSerialNumber / 10) > 10 {
		lResult += "Nominee Serial Number is should be less 2 digit | "
	}

	if pValidateReq.PurposeCode == 6 || pValidateReq.PurposeCode == 7 || pValidateReq.PurposeCode == 8 {
		if pValidateReq.RelationshipWithBO == 0 {
			lResult += "Relationship with BO is Mandatory when Purpose code is 6 or 7 or 8 | "
		} else if pValidateReq.RelationshipWithBO > 13 {
			lResult += "Relationship with BO Value should be less than 13 | "
		}
	} else if pValidateReq.RelationshipWithBO > 13 {
		lResult += "Relationship with BO Value shoule be less than 13 | "
	}

	if pValidateReq.PurposeCode == 6 {
		if pValidateReq.PercentageOfShares == 0 {
			lResult += "Percentage of Shares is Mandatory when purpose code is 6 | "
		}
		if pValidateReq.ResidualSecuritiesFlag == "" {
			lResult += "Residual Securities Flag is Mandatory when purpose code is 6 | "
		}
	}
	return lResult
}

func EmailValidation(pEmail string) bool {
	lIndex := strings.Index(pEmail, "@")
	left := pEmail[:lIndex]
	right := pEmail[lIndex+1:]
	if !strings.Contains(right, ".") {
		return true
	} else {
		lFirstFlag, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+$`, left)
		lSecflag, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, right[:strings.Index(right, ".")])
		var lThirdFlag bool
		Count := strings.Count(right, ".")
		if Count == 1 {
			lDot := right[strings.Index(right, ".")+1:]
			lThirdFlag, _ = regexp.MatchString(`^[a-z]{2,4}$`, lDot)
		} else if Count == 2 {
			lDot := right[strings.Index(right, ".")+1 : strings.LastIndex(right, ".")]
			lSecDot := right[strings.LastIndex(right, ".")+1:]
			lDotFlag, _ := regexp.MatchString(`^[a-z]{2,4}$`, lDot)
			lSecDotflag, _ := regexp.MatchString(`^[a-z]{2,3}$`, lSecDot)
			lThirdFlag = lDotFlag && lSecDotflag
		} else {
			lThirdFlag = false
		}
		return !lFirstFlag || !lSecflag || !lThirdFlag
	}
}
