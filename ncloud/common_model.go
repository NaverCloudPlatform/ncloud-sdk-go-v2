package ncloud

type CommonResponse struct {
	RequestId     *string `xml:"requestId"`
	ReturnCode    *string `xml:"returnCode"`
	ReturnMessage *string `xml:"returnMessage"`
}
