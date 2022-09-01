FROM golang:1.18 AS builder

COPY ./ /src

WORKDIR /src
RUN go mod tidy
RUN go build -o stonks .

FROM ubuntu:kinetic

# Certs
RUN apt-get update
RUN apt-get install ca-certificates -y

COPY --from=builder /src/stonks ./
COPY ./res /res

CMD [ "./stonks" ]