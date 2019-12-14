FROM golang:alpine
LABEL maintainer="tonymtz <hello@tonymtz.com>"

WORKDIR $GOPATH/src/github.com/tonymtz/dolarenbancos

COPY . .

RUN apk update && apk add git npm
RUN go get -u github.com/kardianos/govendor
RUN govendor sync
RUN go build
#RUN cd static && npm install && npm run build
#RUN apk del git npm

EXPOSE 8008

CMD ["./dolarenbancos"]