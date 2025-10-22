package quickwit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Client interface {
	Search(ctx context.Context, indexID, query string) (*SearchResponse, error)
	// StreamSearchIndex(ctx context.Context, indexID string) error
	ListIndexes(ctx context.Context) ([]Index, error)
	GetIndex(ctx context.Context, indexID string) (*Index, error)
	CreateIndex(ctx context.Context, idx IndexConfig) (*Index, error)
	DeleteIndex(ctx context.Context, indexID string) error
	ClearIndex(ctx context.Context, indexID string) error
	DescribeIndex(ctx context.Context, indexID string) (*Describe, error)
	ListSplits(ctx context.Context, indexID string) (*SplitsRes, error)

	CreateSource(ctx context.Context, idx string, src SourceConfig) (*SourceConfig, error)
	DeleteSource(ctx context.Context, indexID, sourceID string) error

	GetElastic(ctx context.Context) (*Cluster, error)
	GetCluster(ctx context.Context) (*Cluster, error)
}

func (c *client) Search(ctx context.Context, indexID, query string) (*SearchResponse, error) {
	var req *http.Request
	var err error

	req, err = http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/%s/search?query=%s", c.endpoint, indexID, query), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	return Request[SearchResponse](c.log, req)
}

// func (c *client) StreamSearchIndex(ctx context.Context, indexID string) error {}

func (c *client) ListIndexes(ctx context.Context) ([]Index, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/indexes", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	indexes, err := GetList[Index](c.log, req)
	if err != nil {
		return nil, err
	}

	return indexes, nil
}

func (c *client) GetIndex(ctx context.Context, indexID string) (*Index, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/indexes/%s", c.endpoint, indexID), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	index, err := Request[Index](c.log, req)
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (c *client) CreateIndex(ctx context.Context, idx IndexConfig) (*Index, error) {
	body := MustMarshall(idx)

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/v1/indexes", c.endpoint),
		body,
	)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	index, err := Request[Index](c.log, req)
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (c *client) DeleteIndex(ctx context.Context, indexID string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/api/v1/indexes/%s", c.endpoint, indexID), nil)
	if err != nil {
		return err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	return RequestNoContent(c.log, req)
}

func (c *client) ClearIndex(ctx context.Context, indexID string) error {
	req, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf("%s/api/v1/indexes/%s/clear", c.endpoint, indexID), nil)
	if err != nil {
		return err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	return RequestNoContent(c.log, req)
}

func (c *client) DescribeIndex(ctx context.Context, indexID string) (*Describe, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/indexes/%s/describe", c.endpoint, indexID), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	index, err := Request[Describe](c.log, req)
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (c *client) ListSplits(ctx context.Context, indexID string) (*SplitsRes, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/indexes/%s/splits", c.endpoint, indexID), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	index, err := Request[SplitsRes](c.log, req)
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (c *client) CreateSource(ctx context.Context, idx string, src SourceConfig) (*SourceConfig, error) {
	body := MustMarshall(src)

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/api/v1/indexes/%s/sources", c.endpoint, idx),
		body,
	)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	s, err := Request[SourceConfig](c.log, req)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func WithBucket(bucketName string) func(cfg *IndexConfig) {
	return func(cfg *IndexConfig) {
		cfg.URI = fmt.Sprintf("s3://%s", bucketName)
	}
}

func WithRetention(d time.Duration) func(cfg *IndexConfig) {
	return func(cfg *IndexConfig) {
		cfg.Retention.Period = fmt.Sprintf("%f seconds", math.Round(d.Seconds()))
	}
}

func (c *client) DeleteSource(ctx context.Context, indexID, sourceId string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/api/v1/indexes/%s/sources/%s", c.endpoint, indexID, sourceId), nil)
	if err != nil {
		return err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	return RequestNoContent(c.log, req)
}

func (c *client) GetElastic(ctx context.Context) (*Cluster, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/_elastic", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	cluster, err := Request[Cluster](c.log, req)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c *client) GetCluster(ctx context.Context) (*Cluster, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/cluster?format=pretty_json", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	for _, interceptor := range c.interceptors {
		interceptor(req)
	}

	//logrus.Debugf("REQ: %+v", req)

	cluster, err := Request[Cluster](c.log, req)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func Request[T any](log logrus.FieldLogger, req *http.Request) (*T, error) {
	t := new(T)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.WithError(err).Error("failed to close response body")
		}
	}()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.WithFields(logrus.Fields{
			"headers": res.Header,
		}).Debugf("QW request failed: %s %s %d", req.Method, req.URL.String(), res.StatusCode)

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.WithError(err).Warn("cannot read error body")
		}
		msg := string(body)

		m := &ErrorMsg{}
		if err := json.Unmarshal(body, &m); err == nil {
			msg = m.Message + m.Error
		}

		return nil, fmt.Errorf("quickwit error: %d - %s", res.StatusCode, msg)
	}

	//payload, _ := io.ReadAll(res.Body)
	//fmt.Printf("%+v\n", string(payload))

	return t, json.NewDecoder(res.Body).Decode(t)
}

func RequestNoContent(log logrus.FieldLogger, req *http.Request) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.WithError(err).Error("failed to close response body")
		}
	}()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.WithError(err).Warn("cannot read error body")
		}
		msg := string(body)

		m := &ErrorMsg{}
		if err := json.Unmarshal(body, &m); err == nil {
			msg = m.Message + m.Error
		}

		return fmt.Errorf("quickwit error: %d - %s", res.StatusCode, msg)
	}

	return nil
}

func GetList[T any](log logrus.FieldLogger, req *http.Request) ([]T, error) {
	t := []T{}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.WithError(err).Error("failed to close response body")
		}
	}()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("quickwit error: %d", res.StatusCode)
	}

	return t, json.NewDecoder(res.Body).Decode(&t)
}

func MustMarshall(i any) *bytes.Buffer {
	// converted := convertYAMLMapToJSONMap(i)

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		panic(fmt.Errorf("cannot JSON encode %+v: %w", i, err))
	}

	return buf
}
