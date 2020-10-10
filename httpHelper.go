package emailServiceHelper

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostRequestJSON(url string, jsonBody []byte, header map[string]string) (statusCode int,
	response []byte, err error) {

	clientHTTP := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))

	if err != nil {
		return 0, response, err
	}

	//add header
	req.Header.Add("Content-Type", "application/json")

	//populate header
	for headerKey, headerVal := range header {
		req.Header.Add(headerKey, headerVal)
	}

	//execute http post
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return 0, response, nil
	}
	defer resp.Body.Close()

	//read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, response, err
	}

	return resp.StatusCode, responseBody, nil
}
