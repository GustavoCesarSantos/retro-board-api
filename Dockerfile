FROM cgr.dev/chainguard/go:latest as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM alpine:3.16
WORKDIR /
COPY --from=build /app/api ./
CMD ["./api"]
