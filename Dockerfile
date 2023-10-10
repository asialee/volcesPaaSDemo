FROM golang:1.20.5-alpine
RUN mkdir /app
COPY ./web.go /app
# COPY ./output /app/
# COPY ./main  /app/
WORKDIR /app
RUN go build -o main web.go
CMD ["/app/main"]