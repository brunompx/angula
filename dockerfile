FROM golang:1.22.5 AS build
WORKDIR /app
COPY . .

RUN go build -v -o app .
RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download && go mod verify

#ENTRYPOINT ["/app"]
RUN chmod +x /app
RUN chmod +x app

CMD ["air"]