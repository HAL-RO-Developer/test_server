FROM golang:latest
RUN go get github.com/HAL-RO-Developer/test_server

EXPOSE 8080

ENTRYPOINT ["test_server"]