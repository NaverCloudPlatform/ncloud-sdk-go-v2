package ncloud

import (
	"reflect"
	"testing"
)

type TestResponse struct {
	RequestId     *string `json:"requestId,omitempty"`
	ReturnCode    *string `json:"returnCode,omitempty"`
	ReturnMessage *string `json:"returnMessage,omitempty"`
	TotalRows     *int32  `json:"totalRows,omitempty"`
}

type TestCommonCode struct {
	Code     *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}

func TestGetCommonResponse(t *testing.T) {
	requestId := String("RequestId")
	returnCode := String("ReturnCode")
	returnMessage := String("ReturnMessage")
	resp := &TestResponse{
		RequestId:     requestId,
		ReturnCode:    returnCode,
		ReturnMessage: returnMessage,
	}
	commonResponse := GetCommonResponse(resp)

	if !reflect.DeepEqual(requestId, commonResponse.RequestId) {
		t.Fatalf("Expected: %s, Actual: %s", StringValue(requestId), StringValue(commonResponse.RequestId))
	}
}

func TestGetCommonResponse_convertNil(t *testing.T) {
	returnCode := String("ReturnCode")
	returnMessage := String("ReturnMessage")
	resp := &TestResponse{
		RequestId:     nil,
		ReturnCode:    returnCode,
		ReturnMessage: returnMessage,
	}
	commonResponse := GetCommonResponse(resp)

	if commonResponse.RequestId != nil {
		t.Fatalf("Expected: nil, Actual: %s", StringValue(commonResponse.RequestId))
	}
}

func TestGetCommonCode(t *testing.T) {
	code := String("code")
	codeName := String("codeName")
	resp := &TestCommonCode{
		Code:     code,
		CodeName: codeName,
	}
	commonCode := GetCommonCode(resp)

	if !reflect.DeepEqual(code, commonCode.Code) {
		t.Fatalf("Expected: %s, Actual: %s", StringValue(code), StringValue(commonCode.Code))
	}
}

func TestGetCommonCode_convertNil(t *testing.T) {
	codeName := String("codeName")
	resp := &TestCommonCode{
		Code:     nil,
		CodeName: codeName,
	}
	commonCode := GetCommonCode(resp)

	if commonCode.Code != nil {
		t.Fatalf("Expected: nil, Actual: %s", StringValue(commonCode.Code))
	}
}
