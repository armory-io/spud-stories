FROM golang:1.14.2 AS builder

COPY . .
RUN GOPATH= CGO_ENABLED=0 go build -o /bin/spud-stories


FROM alpine:3.12

RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/spud-stories /spud-stories
ENTRYPOINT ["./spud-stories"]