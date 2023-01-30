FROM golang:1.19-alpine
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . /code
WORKDIR /code
RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /sample-app .
ENTRYPOINT ["/go/bin/dlv", "--listen=0.0.0.0:2345", "--headless=true", "--only-same-user=false", "--accept-multiclient", "--check-go-version=false", "--api-version=2", "exec", "/sample-app"]
