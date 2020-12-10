FROM golang:1.15

RUN go get -u github.com/cmouli84/xgameodds
RUN go build github.com/cmouli84/xgameodds
RUN mv src/github.com/cmouli84/xgameodds/static ./

ENV TZ=America/Los_Angeles
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 5000

CMD ["xgameodds"]