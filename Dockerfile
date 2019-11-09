FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . /app
RUN go build .

EXPOSE 9999
ENTRYPOINT ["./rest-api-sample"]