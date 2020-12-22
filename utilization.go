package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	consul "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"

	"github.com/burizz/team-utilization/consulkv"
	"github.com/burizz/team-utilization/teams"
)

// Configure logging
func init() {
	// TODO: configure dotenv file with variables
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{}) // can be &log.JSONFormatter

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// log.SetLevel(log.WarnLevel)
	log.SetLevel(log.DebugLevel)
}

const (
	consulAddr = "127.0.0.1:8500"
)

var kv *consul.KV

// Configure Consul Key/Value store
func init() {
	// TODO: Change address with const
	// TODO: move consul connection Init to consul.go package
	// Consul Client
	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Errorf("Err: %v", err)
	}

	// Consul Key/Value store alias
	kv = client.KV()
}

// TODO: Project / Package structure refactoring
func main() {
	var itgixTeams teams.ItgixTeams
	var teamVar string = "Team: "

	var seedDataJSON string = "seed/sample_input_data.json"

	// Parse JSON file into team struct
	if jsonParseErr := parseJSON(seedDataJSON, &itgixTeams); jsonParseErr != nil {
		log.Fatalf("Err: %v", jsonParseErr) // exit if json cannot be parsed
	}

	teamValue, err := json.Marshal(itgixTeams)
	if err != nil {
		log.Errorf("Err : %v", err)
	}

	// TODO: Parse each team into separate consul kv pair
	// TODO: Parse each team member into separate kv / pair
	for _, value := range itgixTeams.Teams {
		fmt.Println(value)
	}

	// Put a new KV pair
	if setKvPairErr := consulkv.SetConsulKV(kv, teamVar, teamValue); setKvPairErr != nil {
		log.Fatalf("Err: %v", setKvPairErr) // exit if Consul KV pair cannot be set
	}

	// Lookup KV pair in Consul
	if kvPair, getKvPairErr := consulkv.GetConsulKV(kv, "team"); getKvPairErr != nil {
		fmt.Println(kvPair)
		log.Errorf("Err: %v", getKvPairErr)
	}
}

func parseJSON(filePath string, team *teams.ItgixTeams) error {
	// TODO: Refactor this to take the JSON from URL where it can be uploaded automatically
	// Read json file and convert to byte slice
	byteValue, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("parseJSON: Cannot open file: %v - %v", filePath, err)
		return err
	}

	// Unmarshal byte array into team struct
	json.Unmarshal(byteValue, &team)
	if err != nil {
		log.Errorf("parseJSON: Cannot unmarshal JSON to team struct: %v", err)
		return err
	}

	log.Infof("JSON parsed successuflly")
	return nil
}
