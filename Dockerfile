ARG GOVERSION="1.16.15"

FROM golang:${GOVERSION}-alpine

WORKDIR $GOPATH/src/github.com/leb4r/gh-repo

COPY . .

RUN go install

WORKDIR $GOPATH

ENTRYPOINT ["gh-repo"]
