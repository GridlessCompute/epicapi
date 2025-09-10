package epicapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// EpicApi is the main struct for the ePIC API wrapper
// It contains the IP, Port, and Password for the device
// as well as the Get and Post structs for making API calls
type EpicApi struct {
	ip       string
	port     string
	password string
	Get      Get
	Post     Post
}

// Get is a struct for making GET requests to the ePIC API
type Get struct {
	api *EpicApi
}

// Post is a struct for making POST requests to the ePIC API
type Post struct {
	api *EpicApi
}

// New creates a new EpicApi struct
// It takes the IP, Port, and Password for the device
// and returns a pointer to the new EpicApi struct
func New(ip, port, password string) *EpicApi {
	e := &EpicApi{
		ip:       ip,
		port:     port,
		password: password,
	}
	e.Get.api = e
	e.Post.api = e
	return e
}

// get is a helper function for making get requests to the ePIC API
func (e *EpicApi) get(endpoint string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("http://%s:%s/%s", e.ip, e.port, endpoint))
	if err != nil {
		return nil, fmt.Errorf("error getting 'http://%s:%s/%s' (%w)", e.ip, e.port, endpoint, err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting 'http://%s:%s/%s' incorrect status code %s", e.ip, e.port, endpoint, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading res body for 'http://%s:%s/%s' (%w)", e.ip, e.port, endpoint, err)
	}

	return body, nil
}

// post is a helper function for making post requests to the ePIC API
func (e *EpicApi) post(endpoint string, payload map[string]any) error {
	var r PostRes

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload for 'http://%s:%s/%s' (%w)", e.ip, e.port, endpoint, err)
	}

	res, err := http.Post(
		fmt.Sprintf("http://%s:%s/%s", e.ip, e.port, endpoint),
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return fmt.Errorf("error posting to 'http://%s:%s/%s' (%w)", e.ip, e.port, endpoint, err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("error posting to 'http://%s:%s/%s' incorrect status code %s", e.ip, e.port, endpoint, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading res body for 'https://%s:%s/%s' (%w)", e.ip, e.port, endpoint, err)
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return fmt.Errorf("error unmarshaling res body for 'https://%s:%s/%s' (%w)", e.ip, e.port, endpoint, err)
	}

	if !r.Result {
		return fmt.Errorf("post returned command failure for 'https://%s:%s/%s' Result: %s", e.ip, e.port, endpoint, r.Error)
	}

	return nil
}
