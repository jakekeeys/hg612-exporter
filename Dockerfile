FROM golang:alpine as build
WORKDIR /build
COPY . .
RUN go get -d -v .
RUN go build -v -o app .

FROM alpine
WORKDIR /service
COPY --from=build /build/app .
ENTRYPOINT ["./app"]