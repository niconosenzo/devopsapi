FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o api .

WORKDIR /dest

# Copy binary from app/api to dest folder
RUN cp /app/api .

EXPOSE 3000

# Run the binary
CMD ["/dest/api"]