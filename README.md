## team-utilization

### Consul Server
docker run \
    -d \
    -p 8500:8500 \
    -p 8600:8600/udp \
    --name=badger \
    consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0


### Get consul server IP
docker exec badger consul members

```
:41
Node      Address          Status  Type    Build  Protocol  DC   Segment
server-1  172.17.0.2:8301  alive   server  1.9.1  2         dc1  <all>
```

### Consul client

docker run \
   --name=fox \
   consul agent -node=client-1 -join=172.17.0.2

```docker exec badger consul members                                                      root@DESKTOP-NG0SFLN 22:35:35
Node      Address          Status  Type    Build  Protocol  DC   Segment
server-1  172.17.0.2:8301  alive   server  1.9.1  2         dc1  <all>
client-1  172.17.0.3:8301  alive   client  1.9.1  2         dc1  <default>
```