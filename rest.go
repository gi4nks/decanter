package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type RestClient struct {
}

func (r *RestClient) Send(vote Vote) error {
	var err error

	buf, err := xml.Marshal(vote)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", settings.RestUrl(), bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	parrot.Debug("response Status:" + resp.Status)
	//parrot.Debug(resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	parrot.Debug("response Body:" + string(body))

	if resp.StatusCode != 201 {
		return errors.New("response Status:" + resp.Status)
	}

	return nil
}
