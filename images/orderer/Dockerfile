# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER
ARG ALPINE_VER
ARG GO_TAGS
FROM alpine:${ALPINE_VER} as base
RUN apk add tzdata
# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

RUN apk add \
	gcc \
	musl-dev \
	git \
	bash \
	make;
ADD . $GOPATH/src/github.com/ke-chain/fabric
WORKDIR $GOPATH/src/github.com/ke-chain/fabric


RUN make orderer GO_TAGS=${GO_TAGS}

ENV FABRIC_CFG_PATH /etc/ke-chain/fabric
VOLUME /etc/ke-chain/fabric
VOLUME /var/ke-chain
COPY  ./build/bin /usr/local/bin
COPY  ./sampleconfig/msp ${FABRIC_CFG_PATH}/msp
COPY  ./sampleconfig/orderer.yaml ${FABRIC_CFG_PATH}
COPY ./sampleconfig/configtx.yaml ${FABRIC_CFG_PATH}
EXPOSE 7050
CMD ["orderer"]


