package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/burizz/team-utilization/teams"
	"github.com/hashicorp/consul/api"
)

const (
	consulAddr = "127.0.0.1:8500"
)

func main() {
	var sourceData string = "testKey"
	var inputValues string = "testValue"

	// TODO: Change address with const
	// Consul Client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// Consul Key/Value store
	kv := client.KV()

	// Put a new KV pair
	p := &api.KVPair{Key: sourceData, Value: []byte(inputValues)}
	_, err = kv.Put(p, nil)
	if err != nil {
		panic(err)
	}

	// Lookup KV pair in Consul
	pair, _, err := kv.Get(sourceData, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)

	// Init Team struct and parse JSON file into it
	var itgixTeam teams.ItgixTeam
	parseJSON("seed/sample_input_data.json", itgixTeam)
}

// func parseJSON(filePath string, target interface{}) error {
func parseJSON(filePath string, team teams.ItgixTeam) error {
	// Maybe refactor this to take the JSON from URL where it can be uploaded automatically
	// Read json file and convert to byte slice
	byteValue, _ := ioutil.ReadFile(filePath)

	// Unmarshal byte array into team struct
	err := json.Unmarshal(byteValue, &team)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(team.Engineers); i++ {
		fmt.Println("Engineer Firstname: ", team.Engineers[i].Firstname)
		fmt.Println("Engineer Lastname: ", team.Engineers[i].Firstname)
	}

	return nil
	// return json.NewDecoder(r.Body).Decode(target)
}
