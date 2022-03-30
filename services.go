package services

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type Service struct {
	endpoint string
}

func CreateService(endpoint string) Service {
	return Service{
		endpoint,
	}
}

func (s *Service) Get(context context.Context, path string, cookie string) (*http.Response, error) {
	return s.request(context, "GET", path, nil, cookie)
}

func (s *Service) Post(context context.Context, path string, cookie string, body map[string]string) (*http.Response, error) {
	return s.request(context, "POST", path, body, cookie)
}

func (s *Service) Put(context context.Context, path string, cookie string, body map[string]string) (*http.Response, error) {
	return s.request(context, "PUT", path, body, cookie)
}

func (s *Service) request(context context.Context, method string, path string, body map[string]string, cookie string) (*http.Response, error) {
	var data io.Reader
	if body != nil {
		jsonD, err := json.Marshal(body)
		data = bytes.NewBuffer(jsonD)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(context, method, s.endpoint+path, data)

	if err != nil {
		return nil, err
	}

	if cookie != "" {
		req.Header.Add("Cookie", "jwt="+cookie)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
