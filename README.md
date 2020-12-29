## team-utilization

### Consul Server
```console
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=consul-server \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

Find consul server IP : 
```
docker exec consul-server consul members

Node      Address          Status  Type    Build  Protocol  DC   Segment
server-1  172.17.0.2:8301  alive   server  1.9.1  2         dc1  <all>
```

Application build 
```console
go build -o bin/utilization-server server/utilization.go
```

Docker
```console
docker build -t utilization-server .
docker run -it --rm utilization-server
```

Consul sample
```go
// Put a new KV pair in Consul
if setKvPairErr := storage.SetConsulKV(kv, teamVar, teamValue); setKvPairErr != nil {
    log.Fatalf("Err: %v", setKvPairErr) // exit if Consul KV pair cannot be set
}

// Lookup KV pair in Consul
if teamValue, getKvPairErr := storage.GetConsulKV(kv, teamVar); getKvPairErr != nil {
    log.Errorf("Err: %v", getKvPairErr)
}
```
