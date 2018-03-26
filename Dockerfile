FROM golang:1.9.4
LABEL maintainer  "c-yokoyama <c.yokoyama.ttr@gmail.com>"

RUN mkdir -p /go/src/github.com/c-yokoyama/mash-slack-mybot
ADD ./credentials.sh  /go/src/github.com/c-yokoyama/mash-slack-mybot
ADD ./main.go /go/src/github.com/c-yokoyama/mash-slack-mybot
ADD ./mynokiahealth  /go/src/github.com/c-yokoyama/mash-slack-mybot/mynokiahealth

RUN apt-get -y update && \
apt-get -y upgrade && \
cd /go/src/github.com/c-yokoyama/mash-slack-mybot && \
go get -v 

WORKDIR  /go/src/github.com/c-yokoyama/mash-slack-mybot/

CMD bash -c "source ./credentials.sh && go run main.go"