# Use official and latest Golang image
FROM golang:latest

# Set working directory
WORKDIR /app

# Copy the source code of this folder inside the container
COPY . .

# Download and install dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

#Expose the port
EXPOSE 8008

#Run the executable
CMD ["./api"]