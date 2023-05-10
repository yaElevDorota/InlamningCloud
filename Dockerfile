FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
#COPY data/*.go ./data/
#COPY templates/* ./templates/
COPY RockPaperScissors/* ./RockPaperScissors/

RUN go build -o /applikationen

EXPOSE 8080
CMD [ "/applikationen" ]