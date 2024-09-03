FROM golang:1.22.5
#FROM golang:1.22.5 AS build
WORKDIR /app
COPY . .

#RUN go build -o binapp .
#RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download && go mod verify

RUN go get
RUN templ generate && CGO_ENABLED=0 GOOS=linux go build -o /app/main .

RUN chmod +x /app/main

FROM scratch
#FROM alpine:latest
COPY --from=0 /app/main /app/main

EXPOSE 8080
#RUN chmod +x ./bin/main


#CMD ["/app/main"]
ENTRYPOINT ["/app/main"]

#CMD ["air"]



