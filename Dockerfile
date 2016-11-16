# Docker image for the Drone Azure Storage plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-azure-storage
#     make deps build docker

FROM alpine:3.3

RUN apk update && \
  apk add \
    ca-certificates \
    python \
    py-pip \
    build-base \
    python-dev \
    libffi-dev \
    openssl-dev && \
  pip install --upgrade \
    pip && \
  pip install \
    azure-common==1.1.4 \
    azure-storage==0.33.0 \
    azure-servicemanagement-legacy==0.20.5 \
    cryptography>=1.5.2 \
    requests==2.11.1 \
    blobxfer==0.12.0 && \
  rm -rf /var/cache/apk/*

ADD drone-azure-storage /bin/
ENTRYPOINT ["/bin/drone-azure-storage"]
