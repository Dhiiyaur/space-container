# multistage  + distroless

FROM golang:1.22.1-alpine3.19 as build

WORKDIR /go/src/neptune

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/neptune

# DOCS https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/neptune /

# COPY .env .

EXPOSE 1000

CMD ["/neptune"]