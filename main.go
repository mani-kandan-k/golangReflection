package main

import (
	"fmt"
	validatefields "govalidator/validateFields"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name    string `validate:"checkUser" FT:""`
	OtpMode string
	PhoneNo int
	Email   string
	Profile UserProfile
}

type UserProfile struct {
	UserName  string
	Password  string
	Followers []UserProfile
	Contacts  []User
}

func main() {

	// ReflectDemo()
	// ArrayLiteralWithIndices()

	// var lUser User
	// lUser.Name = "FT123"
	// lUser.Email = "FT123@gmail.com"

	// MyValidator(lUser)
	NomineeAdditionValidator(validatefields.AssignNomineeDetails())

}

func RequiredIfOneOf(pField validator.FieldLevel) (rFieldRequired bool) {

	// fmt.Println("pField :", pField.Field())
	// fmt.Println("FieldName :", pField.FieldName())
	// fmt.Println("GetTag :", pField.GetTag())
	// fmt.Println("Param :", pField.Param())

	lParamValue := pField.Param()

	// split the param to get each dependent fileds & values
	lParamArr := strings.Split(lParamValue, ";")

	// fmt.Println("lParamArr :", lParamArr)

	// iterate each dependent-fieldName and their values
	for _, val := range lParamArr {

		lParamValues := strings.Split(val, " ")

		// FieldName which the validationField is dependent on
		lDependentFieldValue := pField.Parent().FieldByName(lParamValues[0])

		for i := 1; i < len(lParamValues); i++ {

			// value of dependent field
			val = lParamValues[i]

			// fmt.Println("val :", val)
			// fmt.Println("lDependentFieldValue :", lDependentFieldValue.String())

			fmt.Println("lDependentFieldValue.Type() :", lDependentFieldValue.Type())

			// check if DataType is string
			if lDependentFieldValue.Kind() == reflect.String {

				// check if value matches
				if val == lDependentFieldValue.Interface().(string) {

					// fmt.Println("val == lDependentFieldValue.String() :", val == lDependentFieldValue.String())

					rFieldRequired = true

					rFieldRequired = pField.Parent().FieldByName(pField.FieldName()).Interface().(string) != ""

					fmt.Println("rFieldRequired :", rFieldRequired)
					// if the value matches break the loop
					break
				}

			}

			if lDependentFieldValue.Kind() == reflect.Int {

				lIntval, err := strconv.Atoi(val)
				if err != nil {
					rFieldRequired = false
					// if error break the loop
					break
				}

				if lIntval == lDependentFieldValue.Interface().(int) {
					rFieldRequired = true

					rFieldRequired = pField.Parent().FieldByName(pField.FieldName()).Interface().(int) != 0

					fmt.Println("rFieldRequired int :", rFieldRequired)

					// if the value matches break the loop
					break
				}

			}

		}

	}

	// if rFieldRequired {
	// 	// fmt.Printf("pField~~ %T,|-%v-|", pField.Parent().FieldByName(pField.FieldName()), pField.Parent().FieldByName(pField.FieldName()))
	// 	rFieldRequired = pField.Parent().FieldByName(pField.FieldName()).Interface().(string) != ""
	// } else {
	// 	rFieldRequired = true
	// }

	// fmt.Println("StructFieldName :", pField.StructFieldName())
	// fmt.Println("\nrFieldRequired :", rFieldRequired)

	return rFieldRequired
}

// register custom validation method
func NomineeAdditionValidator(pStructValue interface{}) {

	validate := validator.New()

	// adding custom validation function
	validate.RegisterValidation("required_if_oneof", RequiredIfOneOf)

	lErr := validate.Struct(pStructValue)
	if lErr != nil {
		fmt.Println("lErr :", lErr)
	}
}

func GetFieldTags(pReflectStructField reflect.StructField) (rTag string) {
	return ""
}

func MyValidator(pStructValue interface{}) {

	lStructType := reflect.TypeOf(pStructValue)
	lField := lStructType.Field(0)

	// lFieldName := lField.Name

	lFieldTag := lField.Tag.Get("FT")

	fmt.Println("lFieldTag :", lFieldTag)

	validate := validator.New()
	validate.RegisterValidation("checkUser", ValidUser)

	lErr := validate.Struct(pStructValue)
	if lErr != nil {
		fmt.Println("lErr :", lErr)
	}

}

func ValidUser(pField validator.FieldLevel) bool {

	lReflectValue := pField.Param()
	fmt.Println("pField :", lReflectValue)

	return true
}

func ReflectDemo() {

	var lUser User
	lUser.Name = "Manikandan K"
	lUser.Email = "mani@gmail.com"

	/*

		reflect.TypeOf(Struct)
		|_
		  Type(reflect.Type)
		  |_
		    Field(reflect.StructField) ---> Field.Type(reflect.Type)


	*/

	/* access the struct --> reflect.TypeOf(struct_instance) */
	lUserType := reflect.TypeOf(lUser)

	// lUserType.Align()
	// lUserType.AssignableTo(lUserType)
	// lUserType.Bits()
	// lUserType.ChanDir()
	// lUserType.Comparable()
	// lUserType.ConvertibleTo(lUserType)
	// lUserType.Elem()
	// lUserType.Field(0)
	// lUserType.FieldAlign()
	// lUserType.FieldByIndex([]int{})
	// lUserType.FieldByName("")
	// // lUserType.FieldByNameFunc(lUserType.FieldByName(""))
	// lUserType.Implements(lUserType)
	// lUserType.In(0)
	// lUserType.IsVariadic()
	// lUserType.Key()
	// lUserType.Kind()
	// lUserType.Len()
	// lUserType.Method(0)
	// lUserType.MethodByName("")
	// lUserType.Name()
	// lUserType.NumField()
	// lUserType.NumIn()
	// lUserType.NumMethod()
	// lUserType.NumOut()
	// lUserType.Out(0)
	// lUserType.PkgPath()
	// lUserType.Size()
	// fmt.Println("lUserType.String() :", lUserType.String())

	lField, fieldExist := lUserType.FieldByName("Name") // return --> (reflect.StructField, bool)
	fmt.Println("lField :", lField, "fieldExist :", fieldExist)

	// lField := lUserType.Field(0)

	/* access the Fields Info --> lField */
	// lAnonymous := lField.Anonymous
	// fmt.Println("lAnonymous : ", lAnonymous)
	// lIndex := lField.Index
	// fmt.Println("lIndex : ", lIndex)
	// lName := lField.Name
	// fmt.Println("lName : ", lName).Get("Name")
	// lOffset := lField.Offset
	// fmt.Println("lOffset : ", lOffset)
	// lPkgPath := lField.PkgPath
	// fmt.Println("lPkgPath : ", lPkgPath)
	lTag := lField.Tag
	fmt.Printf("lTag.Get() %T \n", lTag.Get("maxLen"))
	// fmt.Println("lTag : ", lTag)
	// lType := lField.Type
	// fmt.Println("lType : ", lType)

}

func ArrayLiteralWithIndices() {

	// const Name = 20

	var StringArr = []string{5: "Mani"}

	fmt.Println(StringArr)

}
