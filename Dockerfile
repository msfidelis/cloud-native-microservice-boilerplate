FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/change-me

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest

COPY --from=builder /go/src/change-me/main ./
COPY --from=builder /go/src/change-me/configs ./configs

EXPOSE 8080

ENTRYPOINT ["./main"]
