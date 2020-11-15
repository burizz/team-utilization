**Initial release - v0.1** :

**Main** :

API
- [ ] take JSON with monthly utilization of whole team
- [ ] parse JSON into team data structure
- [ ] store in MongoDB
- [ ] fetch from DB and calculate utilization % for each team member for the past month, than store it back in DB
- [ ] fetch from DB and show all historic utilization data of the team
- [ ] expose API endpoints for each of operation
- [ ] unit testing
- [ ] extend to all teams

FrontEnd
- [ ] research and test WebAssembly
- [ ] connect to FE to API
- [ ] import JSON from FE
- [ ] drop-down of teams and engineers
- [ ] nice visualization / graphs of historic data

**Infra** :

- [ ] MongoDB - Docker persistent mounted volume
- [ ] Dockerize applications
    - [ ] API
    - [ ] Frontend
- [ ] Makefile
    - [ ] build/test
    - [ ] deploy dev env: automatically start api/fe and mongo containers
    - [ ] deploy prod env: waypoint into EKS
- [ ] Waypoint deploy of both App and Mongo containers and connecting them together in K8s

 