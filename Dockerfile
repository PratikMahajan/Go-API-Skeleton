FROM golang:1.12-alpine

RUN apk add git
RUN adduser -D app
WORKDIR /home/app

COPY . .

RUN go build -o go-app-twitter

USER app

CMD /home/app/go-app-twitter