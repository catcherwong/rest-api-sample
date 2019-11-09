#FROM golang:latest AS development
#ENV GOPROXY https://goproxy.cn,direct
#WORKDIR /app
#COPY . .
##RUN GOOS=linux GOARCH=amd64 go build -o api
#RUN go build -o api

#FROM alpine:latest AS production
#WORKDIR /root/
#COPY --from=development /app/api .
#EXPOSE 9999
#ENTRYPOINT ["./api"]

FROM golang:latest
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . .
#RUN GOOS=linux GOARCH=amd64 go build -o api
RUN go build -o api
EXPOSE 9999
ENTRYPOINT ["./api"]