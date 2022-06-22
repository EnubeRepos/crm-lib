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
		ID:          "Comissionamento:" + generatedId,
		Name:        "",
		Deleted:     false,
		Description: "",
		CreatedAt:   "idk",
		ModifiedAt:  "idk",
		Status:      "Cancelada",
		Amount:      3000,
		// ProductName:                "Casa",
		SalesDate: "18.06.2022",
		// Release:                    true,
		// ParcelCode:                 "1224",
		// ProductDivision:            "",
		// ParcelCreatedAt:            "home",
		// ProductUnity:               "2006",
		// ProductAddress:             "Rua 123",
		// Payment:                    "10+Bank",
		// DueDate:                    "18.06.2022",
		// NetAmount:                  3456,
		// InvoiceFee:                 66,
		// ExemptAccount:              false,
		// IntegratorFee:              99,
		// AdministrationFee:          54,
		ParcelNumber:   1,
		AmountCurrency: "BRL", //
		//	NetAmountCurrency:          "",
		//	AdministrationFeeCurrency:  "5",
		//	AdministrationFeeConverted: 67,
		//	InvoiceFeeConverted:        68,
		//	IntegratorFeeCurrency:      "5",
		//	CreatedByID:                "555555555",
		//	CreatedByName:              "Thomas",
		//	ModifiedByID:               "",
		//	ModifiedByName:             "",
		//	AssignedUserID:             generatedId,
		//	AssignedUserName:           "Thomas",
		//	AmountConverted:            55,
		SalesID: "9",
		//SalesName: "PV test",
		//	InvoiceFeeCurrency:         "BRL",
		ParcelsID:   "34",
		ParcelsName: "Parcel 1",
		//	NetAmountConverted:         55,
		//	AccountID:                  "5643",
		//	AccountName:                "aa",
		//	IntegratorFeeConverted:     45,
		IsFollowed:              true,
		RecipientDocumentNumber: "67",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	expectedName := "Comissionamento: " + res.ID

	if res.Name != expectedName {
		t.Errorf("Error POST Account %q, wanted %q", res.Name, expectedName)
	}
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
