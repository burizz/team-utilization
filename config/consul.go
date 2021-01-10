package config

import (
	"os"

	"github.com/hashicorp/consul/api"
	consul "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func ConsulConfig() (consulClient *consul.KV, err error) {
	consulAddr := os.Getenv("CONSUL_ADDRESS")

	// If consul host not provided set default localhost value
	if len(consulAddr) == 0 {
		consulAddr = "127.0.0.1:8500"
	}

	envType := os.Getenv("ENV_TYPE")
	config := api.DefaultConfig()

	// If running in Docker change consul host
	if envType == "Docker" || envType == "DOCKER" {
		config.Address = "consul:8500"
		consulAddr = "consul:8500"
	} else {
		config.Address = consulAddr
	}

	log.Debugf("Consul Address is set to %v", consulAddr)

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
