FROM debian:stretch-slim

LABEL maintainer="wushaobo.china@gmail.com"

COPY docker/apt_sources/cn.list /etc/apt/sources.list
RUN apt-get update && \
    apt-get install -y ca-certificates

EXPOSE 80
RUN mkdir -p /var/log/zipper

COPY docker/prod/run.sh /opt/zipper/run.sh
COPY bin/zipper /opt/zipper/zipper
