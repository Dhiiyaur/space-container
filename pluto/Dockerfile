# multistage  + distroless

FROM golang:1.22.1-alpine3.19 as build

WORKDIR /go/src/pluto

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/pluto

# DOCS https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/pluto /

# COPY .env .

EXPOSE 1001

CMD ["/pluto"]