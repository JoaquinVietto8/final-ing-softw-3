FROM golang:1.20

EXPOSE 8080

ADD . /backend
WORKDIR /backend

RUN go mod init mvc-go
RUN go mod tidy
RUN go build -o mvc-go .
RUN chmod +x /backend

EXPOSE 8080

ENTRYPOINT ["/backend/mvc-go"]