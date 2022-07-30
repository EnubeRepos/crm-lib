package commissions

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/common"
	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetAPICommissionService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srv := New(client)
	res, err := srv.GetById("621e7cbe89b53ea11")

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error GET Attachment:: error: %v", err)
		return
	}
}

func TestGetCommisions(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srv := New(client)
	res, err := srv.Get()

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error GET Attachment:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}
}

func TestGetCommisionsByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=in&where%5B0%5D%5Battribute%5D=status&where%5B0%5D%5Bvalue%5D%5B%5D=Cancelada"

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srv := New(client)
	res, err := srv.GetByFilter(filter)

	fmt.Println(res)

	if err != nil {
		t.Errorf("Error GET Attachment:: error: %v", err)
		return
	}

}

func TestPostCommision(t *testing.T) {

	generatedId := common.GenerateUUID()

	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srv := New(client)

	res, err := srv.Post(DomainCommissions{
		ID:                      "Comissionamento:" + generatedId,
		Name:                    "",
		Deleted:                 false,
		Description:             "",
		CreatedAt:               "idk",
		ModifiedAt:              "idk",
		Status:                  "Cancelada",
		Amount:                  3000,
		SalesDate:               "18.06.2022",
		ParcelNumber:            1,
		AmountCurrency:          "BRL",
		SalesID:                 "9",
		ParcelsID:               "34",
		ParcelsName:             "Parcel 1",
		IsFollowed:              true,
		RecipientDocumentNumber: "67",
		ProductName:             "Test",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)

}

func TestPutCommission(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srv := New(client)

	res, err := srv.Put(DomainCommissionsBase{
		ID:                      "62b5ca25af9c03c6c",
		Sales:                   "Thomas",
		Name:                    "Thomas Test",
		Deleted:                 false,
		Description:             "",
		CreatedAt:               "idk",
		ModifiedAt:              "idk",
		Status:                  "Transferida",
		Amount:                  7000,
		SalesDate:               "18.06.2022",
		ParcelNumber:            1,
		AmountCurrency:          "BRL",
		SalesID:                 "9",
		ParcelsID:               "34",
		ParcelsName:             "Parcel 1",
		IsFollowed:              true,
		RecipientDocumentNumber: "67",
		ProductName:             "Test",
	})

	if err != nil {
		t.Errorf("Error PUT Account:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestDeleteCommission(t *testing.T) {
	id := "62acbbf5e285e0de6"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)

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
