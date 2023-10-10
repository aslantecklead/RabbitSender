FROM golang:1.20.5

WORKDIR /build

COPY ./main/main.go ./build

RUN go build -o main main.go

CMD ["./main"]