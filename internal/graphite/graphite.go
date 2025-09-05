package graphite

import (
	"fmt"
	"time"

	"github.com/marpaia/graphite-golang"
)

type Graphite struct {
	Client   *graphite.Graphite
	Host     string
	Port     int
	Protocol string
	Timeout  time.Duration
}

func NewClient(host string, port int, prefix string, timeout time.Duration) *Graphite {

	client, err := graphite.NewGraphite(host, port)
	if err != nil {
		panic(err)
	}

	return &Graphite{
		Host:     host,
		Port:     port,
		Protocol: "tcp",
		Timeout:  timeout,
		Client:   client,
	}
}

func (c *Graphite) Connect() error {
	err := c.Client.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (c *Graphite) Send(metric string, value float64) error {
	timestamp := time.Now().Unix()

	graphiteMetric := graphite.Metric{
		Name:      metric,
		Value:     fmt.Sprintf("%f", value),
		Timestamp: timestamp,
	}

	err := c.Client.SendMetric(graphiteMetric)
	if err != nil {
		return err
	}
	return nil
}

func (c *Graphite) Close() error {
	if c.Client != nil {
		return c.Client.Disconnect()
	}
	return nil
}
