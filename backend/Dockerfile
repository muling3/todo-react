FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . ./

RUN go build -o main main.go

FROM alpine
WORKDIR /app
COPY --from=builder ./app/main .
COPY app.env .

EXPOSE 9090

CMD [ "/app/main" ]