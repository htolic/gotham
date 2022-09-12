FROM golang:1.18.1

RUN groupadd gorunner && \
    useradd -m -g gorunner gorunner

USER gorunner

WORKDIR /usr/src/app

COPY ./getweather.go .
RUN go install getweather.go

CMD ["getweather"]
