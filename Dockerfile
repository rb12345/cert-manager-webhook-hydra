FROM golang:1.24-trixie AS build_deps

RUN apt-get update && apt-get install -y git && apt-get clean

WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_deps AS build

COPY . .

RUN CGO_ENABLED=0 go build -o webhook -ldflags '-w -extldflags "-static"' .

FROM debian:trixie-slim

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

COPY --from=build /workspace/webhook /usr/local/bin/webhook

ENTRYPOINT ["webhook"]
