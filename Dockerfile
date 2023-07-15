FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build -o mcrailway .

FROM scratch

COPY --from=builder /app/mcrailway /app/mcrailway

CMD ["/app/mcrailway"]