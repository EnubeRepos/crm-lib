package invoiceissuance

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "_"
	TOKEN = "_"
)

func TestGetInvoiceIssuance(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)
	res, err := srvInvoiceIssuance.Get()

	if err != nil {
		t.Errorf("Error GET InvoiceIssuance:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET InvoiceIssuance %q, wanted %q", res.Total, expected)
	}
}

func TestGetInvoiceIssuanceById(t *testing.T) {
	expectedId := "628e7c113fdf10966"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)
	res, err := srvInvoiceIssuance.GetById(expectedId)

	if err != nil {
		t.Errorf("Error GET InvoiceIssuance:: error: %v", err)
		return
	}

	if res.ID != expectedId {
		t.Errorf("Error GetId InvoiceIssuance %s, wanted %s", res.ID, expectedId)
	}
}

func TestGetInvoiceIssuanceByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=in&where%5B0%5D%5Battribute%5D=status&where%5B0%5D%5Bvalue%5D%5B%5D=Distratada"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)
	res, err := srvInvoiceIssuance.GetByFilter(filter)

	if err != nil {
		t.Errorf("Error GETBYFILTER InvoiceIssuance:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GETBYFILTER InvoiceIssuance %q, wanted %q", res.Total, 1)
	}
}

func TestPostParcel(t *testing.T) {

	// expectedUserId := "55"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)
	res, err := srvInvoiceIssuance.Post(DomainInvoiceIssuance{
		Name: "Test",
		// AssignedUserId:   expectedUserId,
		DueDate: "18.06.2022",
		Amount:  100,
		// BillingContactId: "1234",
	})

	if err != nil {
		t.Errorf("Error POST InvoiceIssuance:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestPutAuthCode(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)
	res, err := srvInvoiceIssuance.PutAuthCode(DomainInvoiceIssuancePutAuthCode{
		ID:           "628be41ecd213fbfb",
		StatusDetail: "Distratada",
		Status:       "oi",
		// ParcelPaid:   "sim",
	})

	if err != nil {
		t.Errorf("Error PUT InvoiceIssuance:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestPutStatus(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)
	res, err := srvInvoiceIssuance.PutStatus(DomainInvoiceIssuancePutStatus{
		ID:           "628be41ecd213fbfb",
		StatusDetail: "Transferiada",
		Status:       "Transferida",
	})

	if err != nil {
		t.Errorf("Error PUT InvoiceIssuance:: error: %v", err)
		return
	}

	fmt.Println(res)
}

func TestDeleteParcel(t *testing.T) {
	id := "628be41ecd213fbfb"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvInvoiceIssuance := New(client)

	res, err := srvInvoiceIssuance.Delete(id)

	if err != nil {
		t.Errorf("Error DELETE InvoiceIssuance:: error: %v", err)
		return
	}

	if res == false {
		t.Errorf("Error DELETE: InvoiceIssuance not deleted")
		return
	}
}
