FROM debian:jessie

COPY docker/apt_sources/cn.list /etc/apt/sources.list
RUN apt-get update && \
    apt-get install -y curl telnet vim locales-all

ENV LC_ALL en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

EXPOSE 80
RUN mkdir -p /var/log/zipper

COPY docker/prod/run.sh /opt/zipper/run.sh
COPY bin/zipper /opt/zipper/zipper
