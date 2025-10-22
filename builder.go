package quickwit

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type client struct {
	log          logrus.FieldLogger
	endpoint     string
	interceptors []reqModifier
	httpClient   *http.Client
}

type clientOption func(*client)
type reqModifier func(req *http.Request)

func New(opts ...clientOption) Client {
	c := client{
		endpoint:     DefaultEndpoint,
		interceptors: []reqModifier{},
		httpClient:   http.DefaultClient,
		log:          logrus.New(),
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

// Set endpoint
// must have `http://localhost:7280` pattern
func WithEndpoint(endpoint string) func(*client) {
	return func(c *client) { c.endpoint = endpoint }
}

func WithBearerToken(token string) func(*client) {
	return func(c *client) {
		if token == "" {
			return
		}

		c.interceptors = append(c.interceptors, func(req *http.Request) {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		})
	}
}

func WithLogger(logger logrus.FieldLogger) func(*client) {
	return func(c *client) { c.log = logger }
}

func WithBasicAuth(user, password string) func(*client) {
	return func(c *client) {
		c.interceptors = append(c.interceptors, func(req *http.Request) {
			req.SetBasicAuth(user, password)
		})
	}
}

func WithHttpClient(httpClient *http.Client) func(*client) {
	return func(c *client) { c.httpClient = httpClient }
}
