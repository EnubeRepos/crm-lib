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

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func bytesFileToBase64(file []byte) string {
	if len(file) <= 0 {
		return ""
	}
	
	return fmt.Sprintf("data:%v;base64,", http.DetectContentType(file)) + toBase64(file)
}
