# syntax=docker/dockerfile:1

###############################################################
# Builder
###############################################################
FROM golang:1.22 as build

WORKDIR /app

# Copy dependencies
COPY . .

# Build binary
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o main main.go

###############################################################
# Final
###############################################################

# Copy artifacts to a clean image
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

COPY --from=build /app/main ./main

EXPOSE 8000

ENTRYPOINT ["./main"]