FROM        golang:1.20.1
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download
RUN         go build -o app
ENTRYPOINT  ["./app"]