FROM golang:1.23-alpine

WORKDIR /backend

# Install air
RUN go install github.com/air-verse/air@latest
RUN apk add --no-cache postgresql-client

# Copy go.mod and go.sum
COPY go.* ./

RUN go mod download

# Copy the rest of the application code
COPY . .

# Install dependencies and build the application
RUN go build -o main .

# Expose the port
EXPOSE ${PORT}

# Use air for live reloading
CMD ["air"]
