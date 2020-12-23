package libs_http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	@return
	string http status
	[]byte response data
	error
*/
func ReqPostRawJson(reqURL string, body interface{}, header map[string]string, timeout time.Duration) (*http.Response, []byte, error) {
	// encode data
	reqBody := new(bytes.Buffer)
	if err := json.NewEncoder(reqBody).Encode(body); err != nil {
		return nil, nil, err
	}

	// new request
	request, err := http.NewRequest("POST", reqURL, reqBody)
	if err != nil {
		return nil, nil, err
	}

	// set header
	request.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		request.Header.Add(k, v)
	}

	// request
	client := &http.Client{Timeout: timeout}
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()

	// read response data from the cache
	rspBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	// return
	return response, rspBody, nil
}
