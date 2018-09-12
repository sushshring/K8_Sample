# STEP 1 build executable binary
FROM golang:onbuild as builder
COPY . $GOPATH/src/sushshring/K8_Sample/
WORKDIR $GOPATH/src/sushshring/K8_Sample
#get dependancies
#you can also use dep
RUN go get -d -v
#build the binary
RUN go build -o /go/bin/sample
FROM scratch
COPY --from=builder /go/bin/sample /go/bin/sample
USER appuser
ENTRYPOINT ["/go/bin/sample"]
EXPOSE 8080