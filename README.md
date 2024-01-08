# golang-api (CRUD application)
API in golang, it has functions for GET, Put, Create and Delete operations.

docker-compose.yaml >> has two services, one for mysql and other for the golang.


# it will be running on a Docker container,

docker-compose up

# Tests and github actions

The tests for the app will be run via github actions. The tests could be run locally as well by "docker-compose up" and then going in to the golang-api docker container and running the "go test" command from there.