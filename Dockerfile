ARG CODE_VERSION=18.04

FROM ubuntu:${CODE_VERSION}

RUN apt-get update && apt-get install -y curl \ 
    && apt-get clean \
    && apt-get install nginx -y \
    && apt-get install systemd -y \
    && rm -rf /var/lib/apt/lists/*

COPY index.nginx-debian.html /var/www/html/

RUN service nginx restart

ENV USER=benura
ENV SHELL=/bin/bash
ENV LOGNAME=benura

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]