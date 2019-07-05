package shipper_validator

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"strings"
	"regexp"
)

//Register validations for shipper.
func RegisterShipperValidations(v *validator.Validate) {
	v.RegisterValidation("validname", validShipperName)
	v.RegisterValidation("validpin", validPincode)
	v.RegisterValidation("validtransportmode", validTransportMode)
}

//Validation for shipper's name.
func validShipperName(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var regex = "^[A-Za-z0-9]*$"
	var nameregex = regexp.MustCompile(regex)
	return nameregex.MatchString(field.String())
}

//Validation for shipper's pincode.
func validPincode(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	var regex = "^[1-9][0-9]{2,6}$"
	var nameregex = regexp.MustCompile(regex)
	return nameregex.MatchString(field.String()) || field.String() == ""
}

//Validation for shipper's transport mode.
func validTransportMode(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	val := strings.ToLower(field.String())
	if val == "air" || val == "road" || val == "ship" || val == "" {
		return true
	}
	return false
}