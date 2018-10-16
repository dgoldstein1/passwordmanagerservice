FROM golang:1.9

ENV PRJ_DIR $GOPATH/src/github.com/dgoldstein1/passwordservice
# create project dir
RUN mkdir -p $PRJ_DIR
# add src, service communication ,and docs
COPY ./cmd $PRJ_DIR/cmd
RUN mkdir -p mkdir -p /opt/services/passwordservice
COPY ./docs /opt/services/passwordservice/docs/
COPY ./protobuf $PRJ_DIR/protobuf
COPY ./Gopkg.toml $PRJ_DIR
COPY ./Gopkg.lock $PRJ_DIR

# add additional utils
# RUN apk add git

# setup go
ENV GOBIN $GOPATH/bin 
ENV PATH $GOBIN:/usr/local/go/bin:$PATH

# install go utils, realize for hot reloading and dep for installing deps
RUN go get github.com/pilu/fresh 
RUN go get github.com/golang/dep/cmd/dep

# copy over source code
WORKDIR $PRJ_DIR

# install dependencies
RUN dep ensure

# configure entrypoint
COPY ./settings.toml /etc/passwordservice/settings.toml
COPY ./docker/entrypoint.sh /entrypoint.sh
COPY ./runner.conf $PRJ_DIR

# build executable
RUN go build -v -o passwordservice ./cmd/server

ENTRYPOINT ["fresh"]

# expose service ports
EXPOSE 10000
EXPOSE 10001
EXPOSE 8080