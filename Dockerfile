# development stage (lts-alpine)
FROM node:12.16.0-alpine3.11 as ui-devel-stage
WORKDIR /app
COPY ui/package*.json /app/
COPY ui /app/
RUN mkdir /tmp/node_modules && ln -s /tmp/node_modules /app/node_modules

# build angular application (lts-alpine)
FROM node:12.16.0-alpine3.11 as ui-build-stage
WORKDIR /app
COPY ui/package*.json /app/
RUN npm install
COPY ui /app/
RUN npm run-script build

# build Go application
FROM golang:1.15.2-buster as backend-build-stage
COPY .  /go/src/github.com/rclsilver/nventory
WORKDIR /go/src/github.com/rclsilver/nventory
RUN make nventory

# build final image
FROM golang:1.15.2-buster
RUN mkdir /app
WORKDIR /app

COPY --from=ui-build-stage /app/dist/nventory /app/ui
COPY --from=backend-build-stage /go/src/github.com/rclsilver/nventory/nventory /app/nventory

EXPOSE 8080

CMD ["/app/nventory"]
