FROM golang:1.22.0

ADD . /mvc-g
WORKDIR /mvc-g

RUN go mod init mvc-g
RUN go mod tidy
RUN go build -o mvc-g .
RUN chmod +x /mvc-g

ENTRYPOINT ["mvc-go"]