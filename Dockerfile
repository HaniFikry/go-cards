FROM golang:1.16-alpine
RUN apk add build-base

WORKDIR /project/go-cards/

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /project/go-cards/build/myapp .

EXPOSE 8080
ENTRYPOINT [ "/project/go-cards/build/myapp" ]