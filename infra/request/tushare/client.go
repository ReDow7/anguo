package tushare

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	endpoint        string = "https://api.tushare.pro"
	maxRetriesTimes        = 5
	timeLayout             = "20060102"
)

type Response struct {
	Code      int
	Message   string
	RequestId string       `json:"request_id"`
	Data      ResponseData `json:"data"`
}

type ResponseData struct {
	Fields  []string
	Items   [][]interface{}
	HasMore bool `json:"has_more"`
}

type Client struct {
	token string
}

func (r *Response) anyError() error {
	if r == nil {
		return fmt.Errorf("encounter a nil response from tushare without any error")
	}
	if r.Code != 0 {
		return fmt.Errorf("tushare response with status not ok, message : %s", r.Message)
	}
	lengthOfFields := len(r.Data.Fields)
	if lengthOfFields == 0 {
		fmt.Errorf("tushare response with no fields")
	}
	if len(r.Data.Items) == 0 {
		return fmt.Errorf("encouter an empty data.Items fetched from tushare")
	}
	for _, val := range r.Data.Items {
		if len(val) != lengthOfFields {
			fmt.Errorf("tushare response with fields and values size mismatch %d vs %d",
				len(val), lengthOfFields)
		}
	}
	return nil
}

var HttpClient *Client

func InitClient(token string) {
	HttpClient = &Client{token: token}
}

func (c *Client) makeBody(apiName, fields string, params map[string]interface{}) (io.Reader, error) {
	m := map[string]interface{}{
		"api_name": apiName,
		"token":    c.token,
		"fields":   fields,
	}
	if params == nil {
		m["params"] = ""
	} else {
		m["params"] = params
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

func (c *Client) request(body io.Reader) (*Response, error) {
	return c.requestWithRetries(body, maxRetriesTimes)
}

func (c *Client) requestWithRetries(body io.Reader, retries int) (*Response, error) {
	resp, err := http.Post(endpoint, "application/json", body)
	for err != nil {
		log.Printf("error when send a request to tushare, cur times: %d, remain retries: %d",
			retries+1, maxRetriesTimes-retries)
		if retries == maxRetriesTimes {
			return nil, err
		}
		return c.requestWithRetries(body, retries-1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res Response
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	if res.Code != 0 {
		return nil, errors.New(res.Message)
	}
	return &res, nil
}

func fetchTushareRawData(apiName, fields string, params map[string]any) (*Response, error) {
	reqBody, err := HttpClient.makeBody(apiName, fields, params)
	if err != nil {
		return nil, err
	}
	return HttpClient.request(reqBody)
}
