FROM alpine:latest
ADD *.go /coco-ebs-vol-manager/
RUN apk add --update git go                                                 \
  && export GOPATH=/go                                                      \
  && export PROJECT_HOME="github.com/Financial-Times/coco-ebs-vol-manager"  \
  && export PROJECT_PATH=${GOPATH}/src/${PROJECT_HOME}                      \
  && mkdir -p ${PROJECT_PATH}                                               \
  && cp -r /coco-ebs-vol-manager/* $PROJECT_PATH                            \
  && cd $PROJECT_PATH                                                       \
  && go get ./...                                                           \
  && go build                                                               \
  && mv coco-ebs-vol-manager /app                                           \
  && apk del go git                                                         \
  && rm -rf $GOPATH /var/cache/apk/*
CMD exec /app
