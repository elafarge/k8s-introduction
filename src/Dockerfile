# Build container
FROM golang:latest AS build

COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service

# Target container
FROM alpine:latest
RUN apk add --update --no-cache curl
COPY --from=build /src/service /service
RUN chmod +x /service
ENTRYPOINT ["/service"]
