#FROM golang:1.22.5
FROM golang:1.22.5 AS builder
WORKDIR /app
COPY . .

#RUN go build -o binapp .
#RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download && go mod verify
RUN go mod tidy
#RUN go get
RUN templ generate && CGO_ENABLED=0 GOOS=linux go build -o /app/main .

RUN chmod +x /app/main

FROM scratch
#FROM alpine:latest

WORKDIR /app


COPY --from=builder /app/main /app/main

EXPOSE 8080
#RUN chmod +x ./bin/main


CMD ["./main"]
#ENTRYPOINT ["./main"]

#CMD ["air"]



