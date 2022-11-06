package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	HostURL    string
	Username   string
	Password   string
	Insecure   bool
	HTTPClient *http.Client
	UserAgent  string
}

type errorResponse struct {
	Detail  string `json:"detail"`
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func NewClient(hosturl string, username string, password string, insecure bool, useragent string) *Client {
	return &Client{
		HostURL:   hosturl,
		Username:  username,
		Password:  password,
		Insecure:  insecure,
		UserAgent: useragent,
		HTTPClient: &http.Client{
			Timeout: time.Minute,

			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecure,
				},
			},
		},
	}
}

func (c *Client) Get(method string, path string, r interface{}) error {

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.HostURL, path), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			fmt.Printf(errRes.Detail)
			return errors.New(errRes.Detail)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	body, _ := ioutil.ReadAll(res.Body)
	if len(body) != 0 {
		if err = json.Unmarshal(body, &r); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Create(method string, path string, b interface{}, r interface{}) error {
	reqBody, err := json.Marshal(b)
	if err != nil {
		return err
	}

	reqbytes := bytes.NewBuffer(reqBody).Bytes()
	reqString := string(reqbytes)
	log.Printf("req string: %s", reqString)

	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.HostURL, path), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			fmt.Printf(errRes.Detail)
			return errors.New(errRes.Detail)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	body, _ := ioutil.ReadAll(res.Body)
	myString := string(body)
	log.Printf("body: %s", myString)
	if len(body) != 0 {
		if err = json.Unmarshal(body, &r); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			fmt.Printf(errRes.Detail)
			return errors.New(errRes.Detail)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}
	body, _ := ioutil.ReadAll(res.Body)
	if len(body) != 0 {
		if err = json.Unmarshal(body, &v); err != nil {
			return err
		}
	}
	return nil
}
