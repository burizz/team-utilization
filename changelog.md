**Initial release - v0.1** :

**Main** :

App
- [ ] take JSON with monthly utilization of whole team
- [ ] parse JSON into team data structure
- [ ] parse teams, individual members and tracked hours in KV store
- [ ] store in Consul 
- [ ] fetch from Consul and calculate utilization % for each team member for the past month, than store it back in DB
- [ ] fetch from Consul and show all historic utilization data of the team
- [ ] expose API endpoints for each of operation
- [ ] unit testing
- [ ] HTTP handler and API calls

**Infra** :

- [ ] Consul - Docker persistent mounted volume
- [ ] Dockerize applications
    - Dockerfile
    - Docker-compose
        - App
        - Consul server
- [ ] Makefile
    - [ ] build/test
    - [ ] deploy dev env: automatically start - app and consul
    - [ ] deploy prod env: waypoint into EKS
- [ ] Traefik - https://github.com/traefik/traefik
- [ ] Waypoint deploy of both App and Mongo containers and connecting them together in K8s
 
FrontEnd
- [ ] Research and test WebAssembly
- [ ] connect to FE to API
- [ ] import JSON from FE
- [ ] drop-down of teams and engineers
- [ ] nice visualization / graphs of historic data
