# LogTracker Backend & UI

FROM golang:latest

LABEL com.lazarusnetwork.version="0.1-beta"

LABEL developer = "Shachindra <shachindra@lazarus.network>"

WORKDIR /app

RUN git clone https://github.com/TheLazarusNetwork/LogTracker.git 

WORKDIR /app/LogTracker

RUN go build -o LogTracker .

COPY .env .

EXPOSE $PORT

#ENTRYPOINT ["/app/LogTracker/LogTracker"]

CMD ["/app/LogTracker/LogTracker"]