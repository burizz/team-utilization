package config

import (
	"os"

	consul "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func ConsulConfig() (consulClient *consul.KV, err error) {
	consulAddr := os.Getenv("SERVER_ADRESS")

	// If not provided set default localhost value
	if len(consulAddr) == 0 {
		consulAddr = "127.0.0.1:8500"
	}

	config := &consul.Config{
		Address: consulAddr,
	}

	// Initialize Consul Client
	//client, consulInitClientErr := consul.NewClient(consul.DefaultConfig())
	client, consulInitClientErr := consul.NewClient(config)
	if consulInitClientErr != nil {
		log.Errorf("Error cannot create new consul client: %v", consulInitClientErr)
		return nil, consulInitClientErr
	}

	// Consul Key/Value store alias
	kv := client.KV()
	return kv, nil
}
