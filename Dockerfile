FROM golang:1.18.1

RUN groupadd gorunner && \
    useradd -m -g gorunner gorunner

USER gorunner

WORKDIR /usr/src/app

COPY ./scanner.go .
RUN go install scanner.go

ENTRYPOINT ["scanner"]
