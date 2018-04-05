FROM golang:onbuild

RUN mkdir /app

ADD . /app/

COPY ./examples/ /app/examples/

WORKDIR /app

RUN go build -o avalanche .

CMD ["/app/avalanche"]