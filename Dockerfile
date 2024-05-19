FROM golang:1.22 as build-stage
WORKDIR /usr/src/app

COPY . .

RUN go get ./... && go mod download
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o rate ./internal/cmd/main.go

FROM alpine:latest
WORKDIR /usr/src/app

COPY --from=build-stage /usr/src/app/rate .
COPY ./configs/.env ./configs/.env

EXPOSE 8080

CMD [ "./rate" ]