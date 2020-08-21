FROM golang:1.13.1-alpine as builder

ARG tmp_app_name=default_value
ENV APP_NAME=$tmp_app_name

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o $APP_NAME main.go

FROM alpine:3.7

ARG tmp_app_name=default_value
ENV APP_NAME=$tmp_app_name

COPY --from=builder /build/$APP_NAME /usr/local/bin/$APP_NAME
RUN chmod +x /usr/local/bin/$APP_NAME
RUN $APP_NAME version