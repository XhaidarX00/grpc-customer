FROM golang:alpine3.19 as build
WORKDIR /app
COPY . .
RUN go build -o myapp main.go


FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/ .

EXPOSE 8082
CMD [ "/app/" ]