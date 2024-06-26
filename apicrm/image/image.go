package image

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func (svc *APIImageService) Get(imageID string) ([]byte, error) {
	response, err := svc.client.CRMHandlerImage(imageID)

	return response, err
}

func (svc *APIImageService) GetFile(id, entryPointType string) ([]byte, io.ReadCloser, error) {
	response, io, err := svc.client.CRMHandlerFile(id, entryPointType)

	return response, io, err
}

func (svc *APIImageService) ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (svc *APIImageService) BytesFileToBase64(file []byte) string {
	if len(file) <= 0 {
		return ""
	}

	return fmt.Sprintf("data:%v;base64,", http.DetectContentType(file)) + svc.ToBase64(file)
}
