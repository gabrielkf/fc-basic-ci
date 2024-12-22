FROM golang:1.15

WORKDIR /app

COPY . .
RUN go mod init mathApp
RUN go build -o math
RUN chmod +x math

CMD [ "./math" ]