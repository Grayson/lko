FROM golang:1.23

WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify

COPY main.go ./
EXPOSE 8000
RUN go build
CMD [ "./server" ]