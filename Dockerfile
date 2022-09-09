FROM golang:1.18.1

WORKDIR /usr/src/app

COPY ./getweather.go .
RUN go install getweather.go

CMD ["getweather"]