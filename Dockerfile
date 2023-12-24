FROM golang:1.21
# Set destination for COPY
WORKDIR /app
EXPOSE 8080
RUN apt-get update
RUN apt-get install -y default-mysql-client
COPY * /app/
# Download Go modules
#COPY * ./app/
CMD ["go", "run", "main.go"]



