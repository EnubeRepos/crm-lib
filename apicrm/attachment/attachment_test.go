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

func TestGetAttachment(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAttachmentService(client)
	res, err := srvAccount.Get()

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}
}

func TestGetAPIAttachmentService(t *testing.T) {
	id := "6285348b2b9dfc2bb"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAttachment := NewAPIAttachmentService(client)
	res, err := srvAttachment.GetById(id)

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error GET Attachment:: error: %v", err)
		return
	}
}

func TestGetByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=in&where%5B0%5D%5Battribute%5D=role&where%5B0%5D%5Bvalue%5D%5B%5D=Import%20File"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAttachmentService(client)
	res, err := srvAccount.GetByFilter(filter)

	if err != nil {
		t.Errorf("Error GETBYFILTER Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GETBYFILTER Account %q, wanted %q", res.Total, 1)
	}
}

//test post
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

func TestDelete(t *testing.T) {
	id := "62ac9a51b490827b0"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := NewAPIAttachmentService(client)

	res, err := srvAccount.Delete(id)

	if err != nil {
		t.Errorf("Error DELETE Account:: error: %v", err)
		return
	}

	if res == false {
		t.Errorf("Error DELETE: Account not deleted")
		return
	}
}
