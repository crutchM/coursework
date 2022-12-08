FROM golang:latest

RUN go version
ENV GOPATH=/

COPY . ./

RUN go build -o main main.go


CMD ["./main"]