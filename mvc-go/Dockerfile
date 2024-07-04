FROM golang:1.20

EXPOSE 8080

ADD . /mvc-go
WORKDIR /mvc-go

RUN go mod init mvc-go
RUN go mod tidy
RUN go build -o mvc-go .
RUN chmod +x mvc-go

ENTRYPOINT ["./mvc-go"]
