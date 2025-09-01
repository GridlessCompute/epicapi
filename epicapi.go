package epicapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EpicApi struct {
	IP   string
	Port string
	Pass string
	Get  Get
	Post Post
}

type Get struct {
	api *EpicApi
}

type Post struct {
	api *EpicApi
}

func New(ip, port, password string) *EpicApi {
	e := &EpicApi{
		IP:   ip,
		Port: port,
		Pass: password,
	}
	e.Get.api = e
	e.Post.api = e
	return e
}

func (e *EpicApi) GET(endpoint string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("http://%s:%s/%s", e.IP, e.Port, endpoint))
	if err != nil {
		return nil, fmt.Errorf("error getting 'http://%s:%s/%s' (%w)", e.IP, e.Port, endpoint, err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error getting 'http://%s:%s/%s' incorrect status code %s", e.IP, e.Port, endpoint, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading res body for 'http://%s:%s/%s' (%w)", e.IP, e.Port, endpoint, err)
	}

	return body, nil
}

func (e *EpicApi) POST(endpoint string, payload map[string]any) error {
	var r PostRes

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload for 'http://%s:%s/%s' (%w)", e.IP, e.Port, endpoint, err)
	}

	res, err := http.Post(
		fmt.Sprintf("http://%s:%s/%s", e.IP, e.Port, endpoint),
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return fmt.Errorf("error posting to 'http://%s:%s/%s' (%w)", e.IP, e.Port, endpoint, err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("error posting to 'http://%s:%s/%s' incorrect status code %s", e.IP, e.Port, endpoint, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading res body for 'https://%s:%s/%s' (%w)", e.IP, e.Port, endpoint, err)
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return fmt.Errorf("error unmarshaling res body for 'https://%s:%s/%s' (%w)", e.IP, e.Port, endpoint, err)
	}

	if !r.Result {
		return fmt.Errorf("post returned command failure for 'https://%s:%s/%s' Result: %s", e.IP, e.Port, endpoint, r.Error)
	}

	return nil
}

