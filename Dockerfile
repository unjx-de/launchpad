FROM golang:alpine AS go
RUN apk add build-base
WORKDIR /backend

COPY /backend/swagger.sh .
RUN ./swagger.sh install

COPY /backend/go.mod .
RUN go mod download

COPY ./backend .
RUN ./swagger.sh init
RUN go build -o app

FROM alpine AS logo
RUN apk add figlet
WORKDIR /logo

RUN figlet Launchpad > logo.txt

FROM node:alpine AS vue
WORKDIR /frontend
ENV BASE_URL=https://home.unjx.de/

COPY ./frontend/package.json .
COPY ./frontend/package-lock.json .
RUN npm install

COPY --from=go /backend/docs/ /backend/docs/
COPY ./frontend .
RUN npm run types:openapi -i /backend/docs/swagger.json -o src/services/openapi
RUN npm run generate

FROM alpine AS final
WORKDIR /app

COPY --from=logo /logo/logo.txt .

# copy all the configuration files and default bookmark json
COPY --from=go /backend/bookmark/bookmarks.json ./bookmark/bookmarks.json
COPY --from=go /backend/logging/logging.json ./logging/logging.json
COPY --from=go /backend/server/server.json ./server/server.json
COPY --from=go /backend/weather/weather.json ./weather/weather.json
COPY --from=go /backend/system/system.json ./system/system.json
COPY --from=go /backend/auth/auth.json ./auth/auth.json

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

COPY --from=go /backend/app .
COPY --from=vue /frontend/.output/public/ ./templates/

ENTRYPOINT ["/app/entrypoint.sh"]
