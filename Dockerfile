FROM golang:1.19
ENV TZ="Asia/Tokyo"
WORKDIR /go/src/app
COPY . .
RUN go install github.com/cosmtrek/air@latest
