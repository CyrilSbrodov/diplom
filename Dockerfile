FROM golang:latest

WORKDIR /apps

COPY .. /apps

EXPOSE 8282

VOLUME [ "/apps" ]

CMD ["go", "run", "main.go"]

#RUN go mod download && go build -o app

#CMD ./app