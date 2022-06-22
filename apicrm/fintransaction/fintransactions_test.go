package fintransaction

import (
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://app.10maisbank.com.br/api/v1/"
	TOKEN = "Y29ubmVjdF91c2VyX3dvcmtlcnM6R21YZTg4MXR0Ug=="
)

func TestGetTransactions(t *testing.T) {
	expected := 1
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Get()

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GET Account %q, wanted %q", res.Total, expected)
	}

}

func TestGetTransactionById(t *testing.T) {
	expectedId := "629e3a774eaaef34f"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.GetById(expectedId)

	if err != nil {
		t.Errorf("Error GET Account:: error: %v", err)
		return
	}

	if res.ID != expectedId {
		t.Errorf("Error GetId Account %s, wanted %s", res.ID, expectedId)
	}
}

func TestGetTransactionByFilter(t *testing.T) {
	filter := "where%5B0%5D%5Btype%5D=in&where%5B0%5D%5Battribute%5D=transactionType&where%5B0%5D%5Bvalue%5D%5B%5D=Pagamento%20PIX"
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.GetByFilter(filter)

	if err != nil {
		t.Errorf("Error GETBYFILTER Account:: error: %v", err)
		return
	}

	if res.Total == 0 {
		t.Errorf("Error GETBYFILTER Account %q, wanted %q", res.Total, 1)
	}
}

func TestPostTransaction(t *testing.T) {
	expectedValue := 32.0
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvAccount := New(client)
	res, err := srvAccount.Post(DomainFinTransaction{
		ID:                               "",
		Name:                             "Transaction",
		Deleted:                          false,
		Status:                           string(StatusWaiting),
		DateStart:                        "",
		DateEnd:                          "",
		IsAllDay:                         true,
		Duration:                         3,
		Description:                      "",
		CreatedAt:                        "",
		ModifiedAt:                       "",
		TransactionType:                  string(TransactionTypeTED),
		Value:                            float32(expectedValue),
		AccountTransaction:               string(AccountTransactionSent),
		Amount:                           9,
		ConfirmTransaction:               true,
		ScheduleTransaction:              "",
		ExternalOrigin:                   "",
		EntityID:                         "",
		IdEmpotencyKey:                   "",
		CompanyKey:                       "",
		Context:                          "",
		Timestamp:                        "",
		CorrelationID:                    "",
		Version:                          "",
		AuthenticationCode:               "",
		RecipientDocumentValue:           "",
		RecipientDocumentType:            "",
		RecipientType:                    "",
		RecipientName:                    "",
		RecipientAccountBranch:           "",
		RecipientAccountNumber:           "",
		RecipientAccountBankIspb:         "",
		RecipientAccountBankCode:         "",
		RecipientAccountBankName:         "",
		RecipientAccountBalanceValue:     22,
		RecipientAccountBalanceCurrrency: "",
		RecipientStatus:                  "",
		ChannelName:                      "",
		ChannelControlNumber:             "",
		ChannelControlNumberOriginal:     "",
		ChannelOwnerNumber:               "",
		SenderDocumentValue:              "",
		SenderDocumentType:               "",
		SenderType:                       "",
		SenderName:                       "",
		SenderAccountBranch:              "",
		SenderAccountNumber:              "",
		SenderAccountBankIspb:            "",
		SenderAccountBankCode:            "",
		SenderAccountBankName:            "",
		SenderStatus:                     "",
		DateStartDate:                    "06.06.2022",
		DateEndDate:                      "08.08.2022",
		ValueCurrency:                    "",
		ParentType:                       "",
		ParentName:                       "",
		Commission:                       true,
		TransactionCategory:              "",
		ParentID:                         "",
		AssignedUserID:                   "",
		AssignedUserName:                 "",
		//TeamsIds                :          []string `json:"teamsIds"`
		StatusTrackingID:                  "",
		StatusTrackingName:                "",
		ValueConverted:                    3,
		SalesOrderID:                      "",
		SalesOrderName:                    "",
		BankAccountID:                     "",
		BankAccountName:                   "",
		CommissionsID:                     "",
		CommissionsName:                   "",
		DocumentID:                        "",
		DocumentName:                      "",
		ParcelID:                          "",
		ParcelName:                        "",
		ContactID:                         "",
		ContactName:                       "",
		TedID:                             "",
		TedName:                           "",
		UserID:                            "",
		UserName:                          "",
		PixID:                             "",
		PixName:                           "",
		PixKeyID:                          "",
		PixKeyName:                        "",
		InstitutionFinancialRecipientID:   "",
		InstitutionFinancialRecipientName: "",
		InstitutionFinancialSenderID:      "",
		InstitutionFinancialSenderName:    "",
	})

	if err != nil {
		t.Errorf("Error POST Account:: error: %v", err)
		return
	}

	fmt.Println(res)

	// if res.Value != float32(expectedValue) {
	// 	t.Errorf("Error POST Account %q, wanted %f", res.ID, expectedValue)
	// }
}

func TestDeleteTransaction(t *testing.T) {
	id := "62a39735d5083047c"
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
