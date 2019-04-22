# Use the Golang image to build the binary
FROM golang:1.12 as builder

WORKDIR /go/src/marvin

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o marvin

# Multi-stage build to create a lean production image
FROM alpine

COPY --from=builder /go/src/marvin /marvin

EXPOSE 8080

WORKDIR /marvin

CMD ["./marvin"]
