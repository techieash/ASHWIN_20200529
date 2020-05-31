FROM golang:1.14.3
RUN mkdir /crawler
ADD ./crawler
WORKDIR /crawler
RUN go build -o main
CMD ["/crawler/main"]