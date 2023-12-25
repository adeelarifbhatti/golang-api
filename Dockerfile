# syntax=docker/dockerfile:1
FROM golang:1.21
# Set destination for COPY
WORKDIR /app

COPY go.mod go.sum ./
#RUN go mod download
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8080
#RUN
CMD [ "/bin/sh", "-c" , "sleep 10 && /docker-gs-ping"]



