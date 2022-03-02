package crmapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// CRMHandlerGetService make request for CRM.
func (api *CRMAPIClient) CRMHandlerGetService(resource string, params string) ([]byte, error) {
	url := api.BaseURL + resource + params
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+api.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = isValidResponse(res)

	return body, err
}

func (api *CRMAPIClient) CRMHandlerPutService(resource string, payload []byte) ([]byte, error) {
	url := api.BaseURL + resource
	method := "PUT"

	// ctx.Logger.Infoln(url)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+api.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body = bytes.TrimPrefix(body, []byte("239, 187, 191, 26"))

	err = isValidResponse(res)

	return body, err
}

func (api *CRMAPIClient) CRMHandlerPostService(resource string, payload []byte) ([]byte, error) {
	url := api.BaseURL + resource
	method := "POST"

	// fmt.Println(string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic "+api.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("Status Code: " + strconv.Itoa(res.StatusCode))

	if err != nil {
		return nil, err
	}

	err = isValidResponse(res)

	return body, err
}

// CRMHandlerDeleteService delete request in the CRM API
func (api *CRMAPIClient) CRMHandlerDeleteService(resource, params string) ([]byte, error) {
	url := api.BaseURL + resource
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+api.Token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = isValidResponse(res)

	return body, err
}

func isValidResponse(res *http.Response) error {
	if res.StatusCode == http.StatusBadRequest || res.StatusCode == http.StatusConflict || res.StatusCode == http.StatusNotFound {
		errorHeader := res.Header.Get("X-Status-Reason")

		if errorHeader != "" {
			return fmt.Errorf(errorHeader)
		}
	}

	return nil
}