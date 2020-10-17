package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
)

// registerServiceWithConsul registers a new service to consul
// and also enables a health check using a simple small webserver
// which gets automatically started.
// To configure the port you can pass
// PRODUCT_SERVICE_PORT and PRODUCT_HEALTH_PORT as environment variable.
// The default ports are 8100 and 8101.
func registerServiceWithConsul(serviceName string) {
	// connect to consul
	config := api.DefaultConfig()
	consulHost := os.Getenv("CONSUL_HOST")
	if consulHost != "" {
		config.Address = consulHost
	}

	consul, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("could not create consul client %v", err)
	}

	// setup registration
	registration := new(api.AgentServiceRegistration)
	registration.ID = hostname()
	registration.Name = serviceName
	address := hostname()
	registration.Address = address
	port, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalf("wrong port format %v", err)
	}
	registration.Port = port

	// setup simple health detection using a small webserver
	registration.Check = new(api.AgentServiceCheck)
	healthPortNr, err := strconv.Atoi(healthPort()[1:len(healthPort())])
	if err != nil {
		log.Fatalf("wrong healt port format %v", err)
	}
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, healthPortNr)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `I am alive!`)
	})

	go func() {
		err := http.ListenAndServe(healthPort(), nil)
		log.Fatalf("healthcheck webserver failed %v", err)
	}()

	// finally register the service
	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("registering to consul failed %v", err)
	}
}

func port() string {
	p := os.Getenv("PRODUCT_SERVICE_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8100"
	}
	return fmt.Sprintf(":%s", p)
}

func healthPort() string {
	p := os.Getenv("PRODUCT_HEALTH_PORT")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8101"
	}
	return fmt.Sprintf(":%s", p)
}

func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("retrieving hostname failed %v", err)
	}
	return hostname
}
