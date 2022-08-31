FROM golang:1.18 AS builder

COPY ./ .

RUN go mod tidy
RUN go build -o stonks .

FROM ubuntu:kinetic

COPY --from=builder /stonks ./
CMD [ "./stonks" ]