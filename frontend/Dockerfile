# ==========================================
# Etapa 1: Builder (Compilación)
# ==========================================
FROM klakegg/hugo:ext-alpine AS builder

WORKDIR /src

COPY . .

RUN hugo --minify --environment production

# ==========================================
# Etapa 2: Servidor Web (Producción)
# ==========================================
FROM nginx:alpine-slim

LABEL maintainer="MyssTic Warden Cloud Ops"
LABEL environment="production"

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder /src/public /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]