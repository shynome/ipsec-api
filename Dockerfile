FROM golang:1.12.5@sha256:cf0b9f69ad1edd652a7f74a1586080b15bf6f688c545044407e28805066ef2cb as Build
WORKDIR /ipsec-api
ADD go.mod go.sum /ipsec-api/
RUN go mod download
COPY . /ipsec-api
RUN set -e \
  && cd /ipsec-api/cmd/ipsec-api \
  && CGO_ENABLED=0 go build -o main
FROM alpine:3.9.4@sha256:769fddc7cc2f0a1c35abb2f91432e8beecf83916c421420e6a6da9f8975464b6
RUN apk add --no-cache ca-certificates openssl
COPY --from=Build /ipsec-api/cmd/ipsec-api/main /ipsec-api

ENV \
  l2tpdCoonfigFilepath=/ipsec-etc/ppp/chap-secrets \
  ipsecConfigFilepath=/ipsec-etc/ipsec.d/passwd \
  ipsecSecretsFilepath=/ipsec-etc/ipsec.secrets 

CMD [ "/ipsec-api" ]
