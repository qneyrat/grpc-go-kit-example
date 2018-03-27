FROM golang:alpine

RUN apk add --update git
RUN rm -rf /var/cache/apk/*; rm -rf /tmp/*
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN go get github.com/pilu/fresh

WORKDIR /go/src/grpc-go-kit-example/graphql-gateway
ADD . .
#RUN dep ensure

CMD ["fresh"]
EXPOSE 4000
