FROM golang:1.14.2 AS builder
ARG APP_VERSION="dev"

COPY . .
RUN GOPATH= CGO_ENABLED=0 go build -ldflags="-X 'main.version=${APP_VERSION}'" -o /bin/spud-stories


FROM alpine:3.12

RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/spud-stories /spud-stories
ENTRYPOINT ["./spud-stories"]