FROM alpine:latest
WORKDIR /app
COPY hello hello
EXPOSE 8080

ENTRYPOINT ["./hello"]