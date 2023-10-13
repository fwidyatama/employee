FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

# Install migrate
RUN wget -O - -q https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin


COPY . .

RUN go mod tidy

RUN go mod vendor

RUN go build -o /build ./cmd/main.go

EXPOSE 3000

CMD ["/build"]