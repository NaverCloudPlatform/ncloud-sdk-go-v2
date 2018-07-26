package ncloud

import (
	"strconv"
	"reflect"
)

func String(v string) *string {
	return &v
}

func IntString(n int) *string {
	return String(strconv.Itoa(n))
}

func StringList(input []interface{}) []*string {
	vs := make([]*string, 0, len(input))
	for _, v := range input {
		vs = append(vs, v.(*string))
	}
	return vs
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func Bool(v bool) *bool {
	return &v
}

func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

func Int(v int) *int {
	return &v
}

func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

func Int32(v int32) *int32 {
	return &v
}

func Int32Value(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

func Int64(v int64) *int64 {
	return &v
}

func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

func Float32(v float32) *float32 {
	return &v
}

func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}


func StringField(f reflect.Value) *string {
	if f.Kind() == reflect.Ptr && f.Type().String() == "*string" {
		return f.Interface().(*string)
	} else if f.Kind() == reflect.Slice && f.Type().String() == "string" {
		return String(f.Interface().(string))
	}
	return nil
}

func GetCommonResponse(i interface{}) (*CommonResponse) {
	var requestId *string
	var returnCode *string
	var returnMessage *string
	if f := reflect.ValueOf(i).Elem().FieldByName("RequestId"); !f.IsNil() && f.IsValid() {
		requestId = StringField(f)
	}
	if f := reflect.ValueOf(i).Elem().FieldByName("ReturnCode"); !f.IsNil() && f.IsValid() {
		returnCode = StringField(f)
	}
	if f := reflect.ValueOf(i).Elem().FieldByName("ReturnMessage"); !f.IsNil() && f.IsValid() {
		returnMessage = StringField(f)
	}

	return &CommonResponse{
		RequestId:     requestId,
		ReturnCode:    returnCode,
		ReturnMessage: returnMessage,
	}
}

func GetCommonCode(i interface{}) (*CommonCode) {
	var code *string
	var codeName *string
	if f := reflect.ValueOf(i).Elem().FieldByName("Code"); !f.IsNil() && f.IsValid() {
		code = StringField(f)
	}
	if f := reflect.ValueOf(i).Elem().FieldByName("CodeName"); !f.IsNil() && f.IsValid() {
		codeName = StringField(f)
	}

	return &CommonCode{
		Code:    code,
		CodeName: codeName,
	}
}
