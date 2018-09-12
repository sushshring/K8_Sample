# STEP 1 build executable binary
FROM golang:onbuild as builder
COPY . $GOPATH/src/sushshring/K8_Sample/
WORKDIR $GOPATH/src/sushshring/K8_Sample
#get dependencies
#you can also use dep
RUN go get -d -v
#build the binary
RUN go build -o /go/bin/sample
CMD [ "/bin/sh", "/go/bin/sample" ]
EXPOSE 8080