package jackett

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	jackettUrl string
	apikey     string
	id         int
}

func New(url, apiKey string) *Client {
	return &Client{
		jackettUrl: url,
		apikey:     apiKey,
		id:         1670395402779,
	}
}

func (c *Client) Search(search string) (Response, error) {
	var res *http.Response
	var err error
	var r Response

	params := url.Values{}
	params.Add("apikey", c.apikey)
	params.Add("Query", search)
	params.Add("_", strconv.Itoa(c.id))

	request := fmt.Sprintf("%s/api/v2.0/indexers/all/results?%s", c.jackettUrl, params.Encode())
	c.id++

	if res, err = http.Get(request); err != nil {
		return r, errors.New(fmt.Sprintf("http get error: %s", err))
	}
	defer res.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return r, errors.New(fmt.Sprintf("http read body error: %s", err))
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return r, errors.New(fmt.Sprintf("http unmarshal body error:%s", err))
	}

	return r, nil
}

func (c *Client) Download(trackerId, downloadLink string) error {
	var res *http.Response
	var err error
	var r Download

	params := url.Values{}
	params.Add("jackett_apikey", c.apikey)

	request := fmt.Sprintf("%s/bh/%s/?%s&%s", c.jackettUrl, trackerId, params.Encode(), downloadLink)
	c.id++

	if res, err = http.Get(request); err != nil {
		return errors.New(fmt.Sprintf("http get error: %s", err))
	}
	defer res.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return errors.New(fmt.Sprintf("http read body error: %s", err))
	}

	if err = json.Unmarshal(body, &r); err != nil {
		return errors.New(fmt.Sprintf("http unmarshal body error:%s", err))
	}

	if r.Result == "success" {
		return nil
	}

	return errors.New("не удалось поставить на закачку")
}
