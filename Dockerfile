FROM "golang:alpine"

WORKDIR /go/src/mhdbs/mqtt-golang-vernemq

COPY . /go/src/mhdbs/mqtt-golang-vernemq

RUN cd /go/src/mhdbs/mqtt-golang-vernemq && go build -o main

ENTRYPOINT ["./main"]

