FROM golang:1.8
WORKDIR /go/src/personalRepo/cryptogopherticker
COPY /go/src/personalRepo/cryptogopherticker
RUN go get -d -v /go/src/personalRepo/cryptogopherticker
RUN go install -v /go/src/personalRepo/cryptogopherticker
CMD ["cryptogopherticker"]
