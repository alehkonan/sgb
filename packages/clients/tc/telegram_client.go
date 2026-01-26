package tc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

// Creates new telegram client
func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func (c *Client) Updates(offset, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var r UpdatesResponse
	if err = json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("can't parse updates: %w", err)
	}

	return r.Result, err
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return fmt.Errorf("can't send message: %w", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	const errMessage = "can't do request: "
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf(errMessage+"%w", err)
	}

	req.URL.RawQuery = query.Encode()
	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(errMessage+"%w", err)
	}
	defer func() { _ = res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(errMessage+"%w", err)
	}

	return body, err
}

func newBasePath(token string) string {
	return "bot" + token
}
