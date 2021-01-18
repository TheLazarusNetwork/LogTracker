# LogTracker Backend & UI

FROM golang:latest

LABEL com.lazarusnetwork.version="0.1-beta"

LABEL developer = "Shachindra <shachindra@lazarus.network>"

WORKDIR /app

RUN go get github.com/TheLazarusNetwork/LogTracker

RUN go build -o LogTracker github.com/TheLazarusNetwork/LogTracker

COPY .env .

EXPOSE $PORT

#ENTRYPOINT ["/app/LogTracker"]

CMD ["/app/LogTracker"]