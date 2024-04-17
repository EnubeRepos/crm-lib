package image

import (
	"encoding/csv"
	"fmt"
	"testing"

	"github.com/EnubeRepos/crm-lib/client/crmapi"
)

const (
	HOST  = "https://host_espocrm/api/v1/"
	TOKEN = "_"
)

func TestNewAPIImageService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig(HOST, TOKEN))

	srvImage := NewAPIImageService(client)
	_, err := srvImage.Get("62c34a3a955723e4a")

	if err != nil {
		t.Errorf("Error GET Image:: error: %v", err)
		return
	}
}

func TestNewAPIFileService(t *testing.T) {
	client := crmapi.NewCRMAPIClient(crmapi.NewCRMAPIConfig("https://cdp.homolog.cloudanalytics.me/api/v1/", "bmljazpUdWJhckAwMQ=="))
	client.Headers = []crmapi.Headers{
		crmapi.Headers{
			Key:   "Content-type",
			Value: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		},
	}

	srvImage := NewAPIImageService(client)
	data, r, err := srvImage.GetFile("66197ad176b821d23", "download")

	fmt.Println(data)
	// Lê o corpo da resposta como um arquivo CSV
	csvReader := csv.NewReader(r)

	// Define o caractere delimitador do CSV (vírgula por padrão)
	csvReader.Comma = ';'

	// Lê todas as linhas do arquivo CSV
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo CSV:", err)
		return
	}

	// Exibe as linhas do arquivo CSV
	for _, row := range records {
		fmt.Println(row)
	}
}
