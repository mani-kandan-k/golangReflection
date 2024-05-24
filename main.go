package main

import (
	"fmt"
	"reflect"

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

	var lUser User
	lUser.Name = "FT123"
	lUser.Email = "FT123@gmail.com"

	MyValidator(lUser)

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

		reflect.TypeOf
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
