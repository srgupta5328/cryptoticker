FROM golang:onbuild
RUN mkdir /crpytogoticker
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]