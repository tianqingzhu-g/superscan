FROM golang:1.20-alpine AS build
RUN apk add --no-cache git nodejs npm python3
WORKDIR /src
COPY . .
RUN go build -o /bin/superscan ./cmd/superscan

FROM alpine:3.18
RUN apk add --no-cache nodejs npm python3
COPY --from=build /bin/superscan /usr/local/bin/superscan
COPY parsers /usr/local/bin/parsers
ENTRYPOINT ["/usr/local/bin/superscan"]