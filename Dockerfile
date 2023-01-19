FROM golang:1.19-alpine
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . /code
WORKDIR /code
RUN go build -o /sample-app .
ENTRYPOINT /sample-app
