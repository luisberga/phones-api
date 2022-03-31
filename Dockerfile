FROM golang:1.16.14

WORKDIR $GOPATH/src/github.com/lbergamim-daitan/golang-rump-up

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 5000

CMD ["golang-rump-up"]