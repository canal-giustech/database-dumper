FROM alpine

RUN apk update
RUN apk add ca-certificates tzdata
RUN cp /usr/share/zoneinfo/Brazil/East /etc/localtime
RUN echo "Brazil/East" > /etc/timezone
RUN update-ca-certificates

RUN apk add postgresql-client

RUN apk add --no-cache bash

COPY config ~/.aws/config

RUN apk add groff
RUN apk add --no-cache python3 py3-pip
RUN pip3 install --upgrade pip
RUN pip3 install awscli

ENV PGPASSFILE=.pgpass

RUN rm -rf /var/cache/apk/*

ENV CGO_ENABLED=0
COPY main /

ENTRYPOINT ["/main"]