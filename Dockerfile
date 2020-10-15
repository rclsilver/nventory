# development stage (lts-alpine)
FROM node:12.16.0-alpine3.11 as ui-devel-stage
WORKDIR /app
COPY ui/package*.json /app/
COPY ui /app/
RUN mkdir /tmp/node_modules && ln -s /tmp/node_modules /app/node_modules
