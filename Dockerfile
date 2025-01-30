FROM golang:1.23.4-alpine3.21 as base
WORKDIR /go/src/app
COPY . .
RUN go build -o main cmd/main.go

FROM alpine:3.21 as binary
COPY --from=base /go/src/app/main .
EXPOSE 3000
CMD [ "./main" ]
