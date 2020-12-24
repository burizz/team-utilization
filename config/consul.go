package config

import (
	"log"
	"os"

	log "github.com/sirupsen/logrus"

	consul "github.com/hashicorp/consul/api"
)

func ConsulConfig() (consulClient *consul.KV, err error) {
	consulAddr := os.Getenv("SERVER_ADRESS")

	// If not provided set default localhost value
	if len(consulAddr) == 0 {
		consulAddr = "127.0.0.1:8500"
	}

	// TODO: Change address with const
	// Initialize Consul Client
	//client, consulInitClientErr := consul.NewClient(consul.DefaultConfig())
	client, consulInitClientErr := consul.NewClient(consul.Config(consulAddr))

	if consulInitClientErr != nil {
		log.Errorf("ConsulConfig: Error cannot create new consul client: %v", consulInitClientErr)
		return nil, consulInitClientErr
	}

	// Consul Key/Value store alias
	kv := client.KV()
	return kv, nil
}
