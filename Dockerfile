FROM alpine:3.9

WORKDIR /app
EXPOSE 80

COPY . /app

CMD ["./main"]
