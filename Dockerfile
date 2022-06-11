FROM golang:1.17-alpine

RUN apk update; apk add gcc musl-dev gcompat

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./build/app

CMD [ "./build/app" ]