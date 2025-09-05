package main

import (
	"time"
	vm "victoria-vs-graphite/internal/victoria-metrics"

	"victoria-vs-graphite/internal/graphite"
)

type Client interface {
	Connect() error
	Send(metric string, value float64) error
	Close() error
}

func main() {

	vmClient := vm.NewClient(
		"localhost",
		2003,
		"TESTING.",
		5*time.Second,
	)

	err := vmClient.Connect()
	if err != nil {
		panic(err)
	}

	err = vmClient.Send("METRICA.TESTANDO.VALUE", 42.0)
	if err != nil {
		panic(err)
	}

	defer vmClient.Close()

	graphiteClient := graphite.NewClient(
		"localhost",
		2005,
		"test.",
		5*time.Second,
	)

	defer graphiteClient.Close()

	err = graphiteClient.Connect()
	if err != nil {
		panic(err)
	}
	err = graphiteClient.Send("test.value", 42.0)
	if err != nil {
		panic(err)
	}

}
