# Etapa 1: Builder
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Etapa 2: Final (Distroless - Secure & Minimal)
# hadolint ignore=DL3006
FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/main /main
EXPOSE 8585
USER nonroot:nonroot
CMD ["/main"]