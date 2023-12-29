package errorcode

import "net/http"

var (
	InternalServerError                   = NewInternalServerError(500000, "the request failed due to an internal error")
	InternalServerErrorIdGeneration       = NewInternalServerError(500001, "ID generation error")
	InternalServerErrorZipFileCompression = NewInternalServerError(500002, "zip file compression error")
	InternalServerErrorFileUpload         = NewInternalServerError(500003, "file upload error")
	InternalServerErrorFileExport         = NewInternalServerError(500004, "file export error")
	InternalServerErrorFileDownload       = NewInternalServerError(500005, "file download error")
)

type InternalServerErrorErr struct {
	Err
}

func NewInternalServerError(code int32, message string) *InternalServerErrorErr {
	return &InternalServerErrorErr{Err{Code: code, Message: message, HttpStatus: http.StatusInternalServerError}}
}
