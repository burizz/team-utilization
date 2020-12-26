## team-utilization

### Consul Server
```shell
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=consul-server \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

Application build needs to be ran from the root directory, to be able to find the .env and seed file
```
```


Consul sample
```go
// Put a new KV pair
if setKvPairErr := storage.SetConsulKV(kv, teamVar, teamValue); setKvPairErr != nil {
    log.Fatalf("Err: %v", setKvPairErr) // exit if Consul KV pair cannot be set
}

// Lookup KV pair in Consul
if teamValue, getKvPairErr := storage.GetConsulKV(kv, teamVar); getKvPairErr != nil {
    log.Errorf("Err: %v", getKvPairErr)
}
```
