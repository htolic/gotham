FROM golang:1.18.1

WORKDIR /usr/src/app

COPY ./getweather.go .
RUN groupadd gorunner && \
    useradd -g gorunner gorunner && \
    go install getweather.go

USER gorunner
CMD ["getweather"]
