FROM golang:1.18.1

WORKDIR /usr/src/app

COPY ./scanner.go .
RUN go install scanner.go

ENTRYPOINT ["scanner"]
