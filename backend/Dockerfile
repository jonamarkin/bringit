FROM golang:1.26-alpine AS build

WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build -o /bin/bringit-api ./cmd/server

FROM alpine:3.21

RUN adduser -D -H bringit
USER bringit
COPY --from=build /bin/bringit-api /bin/bringit-api
EXPOSE 8080
ENTRYPOINT ["/bin/bringit-api"]
