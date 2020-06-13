FROM node AS node
WORKDIR /web
COPY front .
RUN yarn install
RUN yarn build

FROM golang:alpine AS golang
WORKDIR /server
COPY . .
RUN GOOS=linux go build -o server ./cmd/server

FROM alpine AS server
WORKDIR /web
COPY --from=node /web/dist dist
COPY --from=golang /server/server /server
ENTRYPOINT ["/server"]