FROM golang:1.12.5@sha256:cf0b9f69ad1edd652a7f74a1586080b15bf6f688c545044407e28805066ef2cb as Build
RUN CGO_ENABLED=0 go get -u github.com/shynome/ipsec-api/cmd/ipsec-api

FROM alpine:3.9.4@sha256:769fddc7cc2f0a1c35abb2f91432e8beecf83916c421420e6a6da9f8975464b6
RUN apk add --no-cache ca-certificates
COPY --from=Build /go/bin/ipsec-api /ipsec-api

CMD [ "/ipsec-api" ]
