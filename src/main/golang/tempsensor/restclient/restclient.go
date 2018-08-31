package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type restclient struct {
	url    string
	client http.Client
}

func NewRestClient(url string) *restclient {
	return &restclient{
		url,
		http.Client{
			Timeout: time.Duration(5 * time.Second),
		},
	}
}

func (rc *restclient) PostAsJSON(content interface{}) (resp *http.Response, err error) {
	jsonValue, err := json.Marshal(content)
	if err != nil {
		return
	}
	resp, err = rc.client.Post(rc.url, "application/json", bytes.NewBuffer(jsonValue))
	return
}
