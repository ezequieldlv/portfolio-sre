# ==========================================
# Etapa 1: Builder (Compilación)
# ==========================================
FROM ghcr.io/gohugoio/hugo:v0.148.0 AS builder

WORKDIR /src

COPY . .

RUN hugo --minify --environment production

# ==========================================
# Etapa 2: Servidor Web (Producción)
# ==========================================
FROM nginx:1.25-alpine-slim

LABEL maintainer="MyssTic Warden Cloud Ops"
LABEL environment="production"

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /src/public /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]