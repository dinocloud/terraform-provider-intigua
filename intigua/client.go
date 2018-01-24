package intigua

import (
	"net/http"
	"reflect"
	"encoding/json"
	"bytes"
	"context"
	"io"
	"github.com/google/go-querystring/query"
	"fmt"
	"strings"
)

type Service struct {
	client 	*http.Client
	host 	string
}

func NewService(c *http.Client, host string) *Service {
	if c == nil {
		c = http.DefaultClient
	}
	return &Service{
		client:	c,
		host:	"http://" + host,
	}
}

// NewRequest generates an HTTP request, but does not perform the request.
func (s *Service) NewRequest(ctx context.Context, method, path string, body interface{}, q interface{}) (*http.Request, error) {
	var ctype string
	var rbody io.Reader
	switch t := body.(type) {
	case nil:
	case string:
		rbody = bytes.NewBufferString(t)
	case io.Reader:
		rbody = t
	default:
		v := reflect.ValueOf(body)
		if !v.IsValid() {
			break
		}
		if v.Type().Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
			if !v.IsValid() {
				break
			}
		}
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rbody = bytes.NewReader(j)
		ctype = "application/json"
	}
	req, err := http.NewRequest(method, s.host+path, rbody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if q != nil {
		v, err := query.Values(q)
		if err != nil {
			return nil, err
		}
		query := v.Encode()
		if req.URL.RawQuery != "" && query != "" {
			req.URL.RawQuery += "&"
		}
		req.URL.RawQuery += query
	}
	req.Header.Set("Accept", "application/json")
	if ctype != "" {
		req.Header.Set("Content-Type", ctype)
	}
	return req, nil
}

// Do sends a request and decodes the response into v.
func (s *Service) Do(ctx context.Context, v interface{}, method, path string, body interface{}, q interface{}, lr *ListRange) error {
	req, err := s.NewRequest(ctx, method, path, body, q)
	if err != nil {
		return err
	}
	if lr != nil {
		lr.SetHeader(req)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch t := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(t, resp.Body)
	default:
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return err
}

// Get sends a GET request and decodes the response into v.
func (s *Service) Get(ctx context.Context, v interface{}, path string, query interface{}, lr *ListRange) error {
	return s.Do(ctx, v, "GET", path, nil, query, lr)
}

// Patch sends a Path request and decodes the response into v.
func (s *Service) Patch(ctx context.Context, v interface{}, path string, body interface{}) error {
	return s.Do(ctx, v, "PATCH", path, body, nil, nil)
}

// Post sends a POST request and decodes the response into v.
func (s *Service) Post(ctx context.Context, v interface{}, path string, body interface{}) error {
	return s.Do(ctx, v, "POST", path, body, nil, nil)
}

// Put sends a PUT request and decodes the response into v.
func (s *Service) Put(ctx context.Context, v interface{}, path string, body interface{}) error {
	return s.Do(ctx, v, "PUT", path, body, nil, nil)
}

// Delete sends a DELETE request.
func (s *Service) Delete(ctx context.Context, v interface{}, path string) error {
	return s.Do(ctx, v, "DELETE", path, nil, nil, nil)
}

// ListRange describes a range.
type ListRange struct {
	Field      string
	Max        int
	Descending bool
	FirstID    string
	LastID     string
}

// SetHeader set headers on the given Request.
func (lr *ListRange) SetHeader(req *http.Request) {
	var hdrval string
	if lr.Field != "" {
		hdrval += lr.Field + " "
	}
	hdrval += lr.FirstID + ".." + lr.LastID
	params := make([]string, 0, 2)
	if lr.Max != 0 {
		params = append(params, fmt.Sprintf("max=%d", lr.Max))
	}
	if lr.Descending {
		params = append(params, "order=desc")
	}
	if len(params) > 0 {
		hdrval += fmt.Sprintf("; %s", strings.Join(params, ","))
	}
	req.Header.Set("Range", hdrval)
	return
}