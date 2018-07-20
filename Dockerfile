# STEP 1 build executable binary
FROM golang:alpine as builder
COPY . $GOPATH/src/personalRepo/cryptogopherticker/
WORKDIR $GOPATH/src/personalRepo/cryptogopherticker/
#get dependancies
#you can also use dep
RUN go get -d -v
#build the binary
RUN go build -o /go/bin/cryptogopherticker
# STEP 2 build a small image
# start from scratch
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/cryptogopherticker /go/bin/cryptogopherticker
ENTRYPOINT ["/go/bin/cryptogopherticker"]
