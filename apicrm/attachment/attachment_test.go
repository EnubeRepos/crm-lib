package attachment

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

//base64 invalida
const filetest = "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAoHBwgHBgoICAgLCgoLDhgQDg0NDh0VFhEYIx8lJCIfIiEmKzcvJik0KSEiMEExNDk7Pj4+JS5ESUM8SDc9Pjv/2wBDAQoLCw4NDhwQEBw7KCIoOzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozv/wgARCAMqBLADASIAAhEBAxEB/8QAGwAAAgMBAQEAAAAAAAAAAAAAAAECAwQFBgf/xAAZAQEBAQEBAQAAAAAAAAAAAAAAAQIDBAX/2gAMAwEAAhADEAAAAfZgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACClLjEo3GFm0wyNpgDcYQ3mBm0wTNphRueGJvMKN5iDaYom8xxNxhDcYQ3GKB0DEG450zcc8OgYkbjCjcYQ3GKJvMSN5gZuMAbjCzcYom4wM3GKJ0DAjeY4G98+RuOfI2mFm0ws3HOsNpgDe"

func TestGetAPIAttachmentService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAttachment := NewAPIAttachmentService(client)
	res, err := srvAttachment.GetById("62153f3f7877ee35d")

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error GET Attachment:: error: %v", err)
		return
	}
}

func TestNewAPIAttachmentService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAttachment := NewAPIAttachmentService(client)
	res, err := srvAttachment.Post(DomainAttachment{
		Field:       "idDocumentBack",
		Name:        "6ca89870-c34c-11e9-ae50-089faf9e7db1.jpg",
		Type:        "image/jpeg",
		Size:        283881,
		Role:        "Attachment",
		RelatedType: "Registration",
		File:        filetest,
	})

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error GET Attachment:: error: %v", err)
		return
	}
}
