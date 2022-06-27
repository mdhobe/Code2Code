FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src ./

RUN go build -o /middleware-go .
EXPOSE 8089

CMD [ "/middleware-go" ]
