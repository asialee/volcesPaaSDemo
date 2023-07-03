FROM golang:1.14-alpine
RUN mkdir /app
COPY ./web.go /app
# COPY ./output /app/
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]