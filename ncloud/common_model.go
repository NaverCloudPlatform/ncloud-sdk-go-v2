package ncloud

type CommonResponse struct {
	RequestId     *string `xml:"requestId"`
	ReturnCode    *string `xml:"returnCode"`
	ReturnMessage *string `xml:"returnMessage"`
}

type CommonCode struct {
	Code *string `json:"code,omitempty"`
	CodeName *string `json:"codeName,omitempty"`
}
