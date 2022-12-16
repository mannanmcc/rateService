# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as build

#Project path
ENV PACKAGENAME 'rateservice'
ENV PACKAGE github.com/mannanmcc/rateservice

#ENV variable
ENV GOPATH /go
ENV GOPROJECT ${GOPATH}/src/${PACKAGE}

WORKDIR ${GOPROJECT}


ADD . .


#install ginkgo
RUN go install -mod=mod github.com/onsi/ginkgo/ginkgo@v1.16.5

# #build executable
# RUN go install -mod vendor
