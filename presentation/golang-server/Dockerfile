FROM golang:1.22.0-alpine3.19 as builder
WORKDIR /presentation/golang-server
COPY . .
RUN echo $(ls /presentation/golang-server)
RUN go build -o /bin/main .

FROM alpine:3.19
EXPOSE 8080
COPY --from=builder /bin/main /bin/main
RUN ["chmod", "+x", "/bin/main"]
ENTRYPOINT ["/bin/main"]