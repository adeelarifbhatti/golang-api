# golang-api
API in golang
# it will be running on a Docker container,
docker build . --tag golang-api
docker run -it --rm -p 8080:8080 -v /home/adeel/Desktop/git/golang-api/:/app/ golang-api
