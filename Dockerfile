FROM alpine:latest
WORKDIR /app
COPY main main
ENTRYPOINT ["./main"]
EXPOSE 8080