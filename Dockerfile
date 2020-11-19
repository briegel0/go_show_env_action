# Dockerfile
FROM golang:1.13
WORKDIR /src
COPY . .
RUN go get -u github.com/sethvargo/go-githubactions/...
RUN go build -o /bin/app .
ENTRYPOINT ["/bin/app"]
