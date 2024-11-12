package jackett

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (c *Client) DownloadBase64EncodedTorrentContentBase64(trackerId, link string) (string, error) {
	var res *http.Response
	var err error

	params := url.Values{}
	params.Add("jackett_apikey", c.apikey)

	request := fmt.Sprintf("%s/dl/%s/?%s&%s", c.jackettUrl, trackerId, params.Encode(), link)

	res, err = http.Get(request)
	if err != nil {
		return "", fmt.Errorf("http get error: %w", err)
	}
	defer res.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return "", fmt.Errorf("http read body error: %w", err)
	}

	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)

	if _, err = io.Copy(encoder, bytes.NewReader(body)); err != nil {
		return "", fmt.Errorf("can't copy file content into the base64 encoder: %w", err)
	}

	if err = encoder.Close(); err != nil {
		return "", fmt.Errorf("can't flush last bytes of the base64 encoder: %w", err)
	}

	return buffer.String(), nil
}
