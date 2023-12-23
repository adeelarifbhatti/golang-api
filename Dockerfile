# syntax=docker/dockerfile:1

FROM golang:1.19
# Set destination for COPY
WORKDIR /app
EXPOSE 8080
# Download Go modules
#COPY * ./app/
CMD ["go", "run", "main.go"]



