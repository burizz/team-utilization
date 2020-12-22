package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	consul "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"

	"github.com/burizz/team-utilization/teams"
)

// Configure logging
func init() {
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
	// Consul Client
	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Errorf("Err: %v", err)
	}

	// Consul Key/Value store alias
	kv = client.KV()
}

func main() {
	var itgixTeams teams.ItgixTeams
	//var teamVar string = "Team: "

	var seedDataJSON string = "seed/sample_input_data.json"

	// Parse JSON file into team struct
	if err := parseJSON(seedDataJSON, &itgixTeams); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	//teamValue, err := json.Marshal(itgixTeams)
	//if err != nil {
	//log.Errorf("Err : %v", err)
	//}

	for _, value := range itgixTeams.Teams {
		fmt.Println(value)
	}

	// Put a new KV pair
	//if err := setConsulKV(kv, teamVar, teamValue); err != nil {
	//log.Errorf("setConsulKV: %v", err)
	//}

	// Lookup KV pair in Consul
	if kvPair, err := getConsulKV(kv, "team"); err != nil {
		fmt.Println(kvPair)
		log.Errorf("Err: %v", err)
	}
}

func parseJSON(filePath string, team *teams.ItgixTeams) error {
	// Maybe refactor this to take the JSON from URL where it can be uploaded automatically
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

func setConsulKV(kv *consul.KV, consulKey string, consulValue []byte) error {
	// TODO: add check if key already exists to run the update, otherwise skip it

	// Put a new KV pair
	kp := &consul.KVPair{Key: consulKey, Value: consulValue}

	// Consul CAS used for Check and Set operation; returns true if successful
	success, meta, err := kv.CAS(kp, nil)
	if err != nil {
		log.Errorf("Consul Set: %v : %v", kv, err)
		return err
	}

	if !success {
		setConsulKV(kv, consulKey, consulValue) // retry setting value
	} else {
		log.Debugf("Consul Set: Set Request time: %v", meta.RequestTime)
		log.Infof("Consul Set: updated key ' %v ' to ' %v '", consulKey, string(consulValue))
	}

	log.Debugf("Consul: Set key: [%v] / value: [%v]", kp.Key, string(kp.Value))
	return nil
}

func getConsulKV(kv *consul.KV, consulKey string) (kValue string, err error) {
	// Lookup KV pair in Consul
	kp, meta, err := kv.Get(consulKey, nil)

	log.Debugf("Consul: Get Request time: %v", meta.RequestTime)
	if err != nil {
		log.Errorf("Consul Get: %v : %v", kv, err)
		return "", err
	}
	// TODO: Handle if kp is nil

	log.Debugf("Consul: Get key: [%v] value: [%s]\n", kp.Key, kp.Value)
	return string(kp.Value), nil
}
