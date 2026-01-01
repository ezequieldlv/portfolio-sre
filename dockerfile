FROM klakegg/hugo:ext-alpine AS builder

COPY . /src
WORKDIR /src

RUN hugo --minify

FROM nginx:alpine

COPY --from=builder /src/public /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]