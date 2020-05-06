FROM golang as build
ADD . /usr/lcoal/go/src/goVideo
WORKDIR /usr/lcoal/go/src/goVideo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_server


FROM alpine:3.7

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /goVideo/go_server /usr/bin/go_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/go_server

ENTRYPOINT ["go_server"]