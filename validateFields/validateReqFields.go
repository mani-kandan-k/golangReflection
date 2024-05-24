package validatefields

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateFields(pReqRec interface{}) (rErrMsg string, rErr error) {

	values := reflect.ValueOf(pReqRec)
	types := values.Type()
	// for i := 0; i < values.NumField(); i++ {
	// 	fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
	// }

	for i := 0; i < values.NumField(); i++ {

		lField := types.Field(i).Name
		lFieldType := types.Field(i).Type.String()
		// fmt.Println("field type :", lFieldType)

		// check if field type is STRUCT
		if len(strings.Split(lFieldType, ".")) > 1 {

			fmt.Println(values.Field(i))

			// lErrMsg, lErr := ValidateFields(values.Field(i))
			// if lErr != nil {
			// 	return lErrMsg, lErr
			// }
			// rErrMsg += lErrMsg

		} else {

			if len(strings.Split(lField, "__")) > 1 {

				getValidations := strings.Split(lField, "__")[1]

				required := strings.Split(getValidations, "_")[0] == "m"

				if len(strings.Split(getValidations, "_")) > 1 {

					maxLen, err := strconv.Atoi(strings.Split(getValidations, "_")[1])

					if err != nil {
						return "", err
					}

					// check if value is required
					if required {
						lFieldValue := values.Field(i).String()
						// check if field is Empty
						if lFieldValue == "" {
							rErrMsg += lField + ` is missing | `
							// fmt.Println("errMsg :", rErrMsg)
						} else {
							// check Max-Length is Valid
							if len(lFieldValue) > maxLen {
								rErrMsg += lField + ` is longer than expected | `
							}
						}
					}

				}
			}
		}

	} // for
	return rErrMsg, nil
}

func ValidateMandatoryFields(pReqRec interface{}) (rErrMsg string, rErr error) {

	values := reflect.ValueOf(pReqRec)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {

		lField := types.Field(i).Name
		lFieldType := types.Field(i).Type.String()

		// fmt.Println("field type :", lFieldType)

		// check if field type is STRUCT
		if len(strings.Split(lFieldType, ".")) > 1 {

			fmt.Println(values.Field(i))
			lMsg, lErr := ValidateMandatoryFields2(values.Field)
			if lErr != nil {
				return "error", lErr
			}

			rErrMsg += lMsg

		} else {

			// fmt.Println(" lField[:len(lField)-1] :", lField[len(lField)-1:])

			if lField[len(lField)-1:] == "_" {

				lFieldValue := values.Field(i).String()

				fmt.Println(" lFieldValue :", lFieldValue)

				if lFieldValue == "" {
					rErrMsg += lField + `is missing | `
				}
			}
		}

	} // for

	return rErrMsg, nil
}

func ValidateMandatoryFields2(pReqRec interface{}) (rErrMsg string, rErr error) {

	values := reflect.ValueOf(pReqRec)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {

		lField := types.Field(i).Name
		lFieldType := types.Field(i).Type.String()

		// fmt.Println("field type :", lFieldType)

		// check if field type is STRUCT
		if len(strings.Split(lFieldType, ".")) > 1 {

			fmt.Println(values.Field(i).Type())

		} else {

			// fmt.Println(" lField[:len(lField)-1] :", lField[len(lField)-1:])

			if lField[len(lField)-1:] == "_" {

				lFieldValue := values.Field(i).String()

				fmt.Println(" lFieldValue :", lFieldValue)

				if lFieldValue == "" {
					rErrMsg += lField + ` is missing | `
				}
			}
		}

	} // for

	return rErrMsg[:len(rErrMsg)-2], nil
}
