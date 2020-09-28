package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	baseURL string = "https://api.46elks.com"
	sendURL string = baseURL + "/a1/sms"
)

func request(method, url string, elk *Elks46, form url.Values) *http.Request {
	var (
		req *http.Request
		err error
	)
	if form != nil {
		req, err = http.NewRequest(method, url, strings.NewReader(form.Encode()))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(elk.username, elk.password)
	return req
}

func call(elk *Elks46, r *http.Request) (map[string]interface{}, error) {
	var (
		response map[string]interface{}
	)
	res, err := elk.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	return response, json.Unmarshal(data, &response)
}
