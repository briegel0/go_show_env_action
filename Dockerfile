# Dockerfile
FROM golang:1.14.3-alpine
WORKDIR /src
COPY . .
RUN go build -o /bin/app .
ENTRYPOINT ["/bin/app"]
