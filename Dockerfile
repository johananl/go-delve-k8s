FROM golang:1.19-alpine
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY . /code
WORKDIR /code
RUN go build -o /sample-app .
ENTRYPOINT [ "/go/bin/dlv" , "--listen=0.0.0.0:50100", "--headless=true", "--only-same-user=false", "--accept-multiclient", "--check-go-version=false", "exec", "/sample-app"]
