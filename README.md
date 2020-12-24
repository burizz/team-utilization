## team-utilization

### Consul Server
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=consul-server \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0

Application build needs to be ran from the root directory, to be able to find the .env and seed file
```
```