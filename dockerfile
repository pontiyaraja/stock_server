FROM golang:1.12

ADD . $GOPATH/src/github.com/pontiyaraja/stock_server

WORKDIR $GOPATH/src/github.com/pontiyaraja/stock_server/cmd

#RUN go install -v

# Install the package
RUN go install $GOPATH/src/github.com/pontiyaraja/stock_server/cmd

ENTRYPOINT $GOPATH/bin/cmd

# This container exposes port 8080 to the outside world
EXPOSE 8060
