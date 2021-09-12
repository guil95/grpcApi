FROM golang:buster as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/checkout

EXPOSE 8000
CMD ["/usr/local/bin/checkout"]