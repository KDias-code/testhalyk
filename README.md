# Service of Library with using framework Fiber

# In this project used technologys:
* Docker
* Prometheus
* Healthz
* DB Postgres

# Descriptions:
* cmd/main.go - main file where we start our project(server start, database start and etc.)
* config/config.go - structures of the configs with parametres with path's , name's of db, host and etc.
* internal - directory to buisness logic
* internal/domain/entity/models - models of our project, so here we see structures of the our entity's.
* internal/handler - in this directory we can find all hadnle's of our project.
* internal/server/server.go - settings of server.
* internal/service/mock - file which generated by Mockgen, to source of our service.
* internal/service - place where we save our interfaces, business logic and etc.
* internal/store - here we can see our repository, where logic of the work with database.
* pkg/db/postgres - our project work by Docker, but if you don't want to work with docker or have any troubles, you can start with this file without docker.
* pkg/db/schema - migrations files with sql side code to create db or drop it.
* pkg/error/error.go - file with done errors to use in a all project.
* prometheus/prometheus.yml - file with the settings

# Request - Response
## Create author(request): 
### curl --location 'http://localhost:8888/author' \
### --header 'Content-Type: application/json' \
### --data '{
###    "full_name": "Cristiano Ronaldo",
###   "nick": "CR7",
###    "speciality": "Football"
### }'
## Create author(response):
### {
###    "message": "created"
### }

## Get author(request):
### curl --location --request GET 'http://localhost:8888/author' \
### --header 'Content-Type: application/json' \
### --data '{
###     "full_name": "Cristiano Ronaldo",
###     "nick": "CR7",
###    "speciality": "Football"
### }'
## Get author(response):
### {
###        "author_id": 1,
###        "full_name": "Cristiano Ronaldo",
###        "nick": "CR7",
###        "speciality": "Football"
### }

## Update author(request):
### curl --location --request PATCH 'http://localhost:8888/author/1' \
### --header 'Content-Type: application/json' \
### --data '{
###     "nick": "SUUUUUUUIIII"
### }'
## Update author(response):
### {
###     "message": "updated"
### }

## Delete author(request):
### curl --location --request DELETE 'http://localhost:8888/author/2' \
### --header 'Content-Type: application/json' \
### --data '{
###     "nick": "SUUUUUUUIIII"
### }'
## Delete author(response):
### {
###     "message": "deleted"
### }
