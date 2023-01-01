FROM golang:1.17 AS build
WORKDIR /build
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o api -v main.go

FROM alpine:3.9
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=build /build/api .


# ENTRYPOINT [ "/app/api" ]
# CMD [ "./app/api" ]
CMD ["sh", "-c", "/app/api"]