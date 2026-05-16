FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /boobles_backend

EXPOSE 8080
CMD [ "/boobles_backend" ]
