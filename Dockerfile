FROM golang:1.18.1

RUN groupadd -g 1001 gorunner && \
    useradd -u 1001 -m -g gorunner gorunner

USER gorunner

WORKDIR /usr/src/app

COPY ./scanner.go .
RUN mkdir /home/gorunner/scanner && \
    go install scanner.go

ENTRYPOINT ["scanner"]
