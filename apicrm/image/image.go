package image

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func (svc *APIImageService) Get(imageID string) ([]byte, error) {
	response, err := svc.client.CRMHandlerImage(imageID)

	return response, err
}

func (svc *APIImageService) toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (svc *APIImageService) bytesFileToBase64(file []byte) string {
	if len(file) <= 0 {
		return ""
	}

	return fmt.Sprintf("data:%v;base64,", http.DetectContentType(file)) + svc.toBase64(file)
}
