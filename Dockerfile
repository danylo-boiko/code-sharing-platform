FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/main.go

CMD [ "./main" ]
