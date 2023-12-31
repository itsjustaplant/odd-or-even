FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /odd-or-even

EXPOSE 3007

ENV PORT 3007

ENV HOSTNAME "0.0.0.0"

CMD [ "/odd-or-even" ]

