FROM golang:alpine as builder

#ENV GOPROXY https://goproxy.cn

COPY ./ /source/

WORKDIR /source/
RUN go build -o NotificationService main.go

FROM alpine
# China mirrors
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Taipei
COPY --from=builder /source/NotificationService /app/NotificationService

WORKDIR /app
EXPOSE 80

ENTRYPOINT ["/app/NotificationService"]