# -- BUILD --
FROM node:alpine@sha256:4faa7d0ec8b49051ddf8336fa4bca7eb42a16855e29ea4a6fe3204482a0bc220 as build

WORKDIR /usr/src/app

COPY package.json .
COPY yarn.lock .
RUN yarn
COPY . .
RUN yarn build

# -- RELEASE --
FROM nginx:stable-alpine as release

COPY --from=build /usr/src/app/build /usr/share/nginx/html
# copy .env.example as .env to the release build
COPY --from=build /usr/src/app/.env /usr/share/nginx/html/.env
RUN rm -rf /etc/nginx/conf.d
COPY conf /etc/nginx

RUN apk add --update nodejs
RUN apk add --update npm
RUN npm install -g runtime-env-cra@0.2.4

WORKDIR /usr/share/nginx/html

EXPOSE 80

CMD ["/bin/sh", "-c", "runtime-env-cra && nginx -g \"daemon off;\""]
