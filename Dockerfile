FROM golang:1.14-alpine
RUN mkdir /app
COPY ./main /app/
COPY ./output /app/
WORKDIR /app
# RUN go build -o main ./web.go 
# CMD ["/app/main"]