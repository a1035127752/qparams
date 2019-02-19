package qparams

import (
	"net/url"
	"errors"
	"reflect"
	"fmt"
	"strconv"
)

const tag  = "json"

var (
	errNonPointerTarget = errors.New("invalid Unmarshal target. must be a pointer")
	errInvalidURL = errors.New("invalid url provided")
	errInvalidType = errors.New("invalid field type")
)


func unmarshalField(v reflect.Value, t reflect.Type, i int, qs url.Values) error {
	var (
		field    reflect.StructField
		paramVal string
		tagVal   string
	)

	field = t.Field(i)

	tagVal = field.Tag.Get(tag)
	if tagVal == "" {
		return nil
	}
	paramVal = qs.Get(tagVal)
	if len(paramVal) == 0 {
		return nil
	}
	switch {
	case field.Type.Kind() == reflect.String:
		v.Field(i).SetString(paramVal)
	case field.Type.Kind() == reflect.Int:
		vint,_ := strconv.ParseInt(paramVal,10,64)
		v.Field(i).SetInt(vint)
	case field.Type.Kind() == reflect.Bool:
		v.Field(i).SetBool(paramVal == "true")
	case field.Type.Kind() == reflect.Float64 || field.Type.Kind() == reflect.Float32:
		vfloat,_ := strconv.ParseFloat(paramVal,64)
		v.Field(i).SetFloat(vfloat)
	default:
		return fmt.Errorf(errInvalidType.Error(), field.Name)
	}

	return nil
}

func Unmarshal(u *url.URL, i interface{}) error{
	if u == nil {
		return errInvalidURL
	}

	iVal := reflect.ValueOf(i)
	if iVal.Kind() != reflect.Ptr || iVal.IsNil() {
		return errNonPointerTarget
	}

	v := iVal.Elem()
	t := v.Type()

	qs := u.Query()

	for i := 0; i < t.NumField(); i++ {
		if err := unmarshalField(v, t, i, qs); err != nil {
			return err
		}
	}
	return nil
}


