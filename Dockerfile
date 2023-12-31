# build backend
FROM docker.io/golang:alpine AS back
RUN mkdir /app
COPY backend/ /app/
WORKDIR /app
RUN go build -o main .

# build frontend
FROM docker.io/node:alpine AS front
RUN mkdir /app
COPY frontend/ /app/
WORKDIR /app
RUN npm install
RUN npm run build

# deploy
FROM docker.io/alpine:latest
WORKDIR /
COPY --from=back /app/main main
COPY --from=front /app/build/ static

ENTRYPOINT [ "/main" ]
