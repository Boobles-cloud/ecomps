FROM alpine/git:latest AS cloner
WORKDIR /src/app

RUN git clone https://github.com/Mikel-217/boobles.git .

FROM golang:latest AS builder

WORKDIR /app

COPY --from=clonner /src/app .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /boobles_backend


FROM alpine:latest
WORKDIR /

COPY --from=builder /boobles_backend /boobles_backend

EXPOSE 8080
CMD [ "/boobles_backend" ]
