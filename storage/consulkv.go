package storage

import (
	"errors"

	consul "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

// GetConsulKV - get Key/Value pair from Consul KV Store
func GetConsulKV(kv *consul.KV, consulKey string) (consulValue string, err error) {
	// Lookup KV pair in Consul
	kp, meta, kvGetErr := kv.Get(consulKey, nil)

	if kvGetErr != nil {
		log.Errorf("Consul Get: %v : %v", kv, kvGetErr)
		return "", kvGetErr
	}

	if kp == nil {
		missingKeyErr := errors.New("Consul Get: key does not exist")
		return "", missingKeyErr
	}

	log.Debugf("Consul: Get key: [%v] value: [%s]\n", consulKey, string(kp.Value))
	log.Debugf("Consul: Get Request time: %v", meta.RequestTime)
	return string(kp.Value), nil
}

// SetConsulKV - set Key/Value pair in Consul KV Store
func SetConsulKV(kv *consul.KV, consulKey string, consulValue string) error {
	byteFmtConsulValue := []byte(consulValue)
	// Get current value from KV store
	pair, _, getKvPairErr := kv.Get(consulKey, nil)
	if getKvPairErr != nil {
		log.Errorf("Consul Set: Failed reading Consul key %v : %v", consulKey, getKvPairErr)
		return getKvPairErr
	}

	// If current value is empty or different set it now
	if pair == nil || string(pair.Value) != consulValue {
		var kp *consul.KVPair
		if pair == nil {
			kp = &consul.KVPair{Key: consulKey, Value: byteFmtConsulValue}
		} else {
			// modifying index used for additional verification, will update key only if
			// it matches the last index that modified this key - https://www.consul.io/api/kv.html#modifyindex
			kp = &consul.KVPair{Key: consulKey, Value: byteFmtConsulValue, ModifyIndex: pair.ModifyIndex}
		}

		// Consul kv.CAS used for Check and Set operation; returns true if successful
		success, meta, setKvPairErr := kv.CAS(kp, nil)
		if setKvPairErr != nil {
			log.Errorf("Consul Set: %v : %v", kp.Key, setKvPairErr)
			return setKvPairErr
		}

		// Retry setting value if previous attempt was not successful
		if !success {
			SetConsulKV(kv, consulKey, consulValue)
		} else {
			log.Debugf("Consul Set: Set Request time: %v", meta.RequestTime)
			log.Debugf("Consul Set: key: [%v] / value: [%v]", kp.Key, string(kp.Value))
		}
		return nil
	} else {
		// if value is not changed, skip
		log.Infof("Consul Set: key: [%v] : skipping, value has not changed", consulKey)
		return nil
	}
}
