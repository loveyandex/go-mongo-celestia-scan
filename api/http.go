package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	http   *http.Client
	ApiUrl string
	Auth   string
}

func NewApi(url, token string) *Api {
	return &Api{
		http:   &http.Client{},
		ApiUrl: url,
		Auth:   token,
	}

}

func (api *Api) JsonPost(path string, body interface{}) (interface{}, error) {

	url := api.ApiUrl + path
	method := "POST"

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	if err != nil {

		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+api.Auth)

	res, err := client.Do(req)
	if err != nil {

		return nil, err
	}
	defer res.Body.Close()

	bodyresp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var reps interface{}
	err2 := json.Unmarshal(bodyresp, &reps)
	// fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	fmt.Printf("reps: %v\n", reps)

	return reps, err2

}
func (api *Api) JsonPut(path string, body interface{}) (interface{}, error) {

	url := api.ApiUrl + path
	method := "PUT"

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))

	if err != nil {

		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+api.Auth)

	res, err := client.Do(req)
	if err != nil {

		return nil, err
	}
	defer res.Body.Close()

	bodyresp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var reps interface{}
	err2 := json.Unmarshal(bodyresp, &reps)
	// fmt.Printf("res.StatusCode: %v\n", res.StatusCode)
	fmt.Printf("reps: %v\n", reps)

	return reps, err2

}

func (api *Api) JsonGet(path string) ([]byte, error) {

	url := api.ApiUrl + path
	method := "GET"
	client := api.http
	req, err := http.NewRequest(method, url, nil)

	if err != nil {

		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+api.Auth)

	res, err := client.Do(req)
	if err != nil {

		return nil, err
	}
	defer res.Body.Close()

	bodyresp, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		return nil, err
	}
	var reps interface{}
	err2 := json.Unmarshal(bodyresp, &reps)
	// fmt.Printf("string(bodyresp): %v\n", string(bodyresp))

	return bodyresp, err2

}
