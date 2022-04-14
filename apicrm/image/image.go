package image

import (
	"encoding/base64"
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
	bytes := file

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	return base64Encoding
}
