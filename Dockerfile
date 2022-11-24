# Этап, на котором выполняется сборка приложения
FROM golang:alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main  cmd/main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3
#FROM scratch
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]