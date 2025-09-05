package vm

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	Host    string
	Port    int
	Prefix  string
	Timeout time.Duration
	Conn    net.Conn
}

func NewClient(host string, port int, prefix string, timeout time.Duration) *Client {
	return &Client{
		Host:    host,
		Port:    port,
		Prefix:  prefix,
		Timeout: timeout,
	}
}

func (c *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), c.Timeout)
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

func (c *Client) Send(metric string, value float64) error {
	timestamp := time.Now().Unix()
	fullMetric := fmt.Sprintf("%s%s %f %d\n", c.Prefix, metric, value, timestamp)
	_, err := fmt.Fprint(c.Conn, fullMetric)
	return err
}

func (c *Client) Close() error {
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}
