FROM golang:1.21
# Set destination for COPY
WORKDIR /app
EXPOSE 8080
COPY * /app/
# Download Go modules
#COPY * ./app/
CMD ["go", "run", "main.go"]



