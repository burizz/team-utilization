package config

import (
	"net/http"
	"os"

	consul "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func ConsulConfig() (consulClient *consul.KV, err error) {
	consulAddr := os.Getenv("CONSUL_ADDRESS")

	// If not provided set default localhost value
	if len(consulAddr) == 0 {
		consulAddr = "127.0.0.1:8500"
	}

	log.Debugf("Consul Address is set to %v", consulAddr)

	// TODO: fix this for Docker - https://didil.medium.com/building-a-simple-distributed-system-with-go-consul-39b08ffc5d2c
	//config := api.DefaultConfig()
	//config.Address = "consul:8500"

	config := &consul.Config{
		Address:    consulAddr,
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	}

	// Initialize Consul Client
	client, consulInitClientErr := consul.NewClient(config)
	if consulInitClientErr != nil {
		log.Errorf("Error cannot create new consul client: %v", consulInitClientErr)
		return nil, consulInitClientErr
	}

	// Consul Key/Value store alias
	kv := client.KV()
	return kv, nil
}
