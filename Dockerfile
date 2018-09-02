FROM golang:1.10.4

RUN mkdir -p $GOPATH/src/github.com/zkrhm/imd-socialnetwork
WORKDIR $GOPATH/src/github.com/zkrhm/imd-socialnetwork/

COPY . .
COPY Gopkg.toml Gopkg.lock ./

RUN apt-get install curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure -vendor-only
RUN go install

ENTRYPOINT [ "imd-socialnetwork" ]

EXPOSE 8000

