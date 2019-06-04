FROM golang:1.12.5@sha256:cf0b9f69ad1edd652a7f74a1586080b15bf6f688c545044407e28805066ef2cb as Build
RUN go get -u github.com/shynome/ipsec-api/cmd/ipsec-api

FROM debian:stable-slim@sha256:a02ca7e73e03d13a57e47ad6bd5cf77a15384f46637cdf1d3a6fda48619a46a0
COPY --from=Build /go/bin/ipsec-api /ipsec-api

ENV \
  l2tpdCoonfigFilepath='/ipsec-etc/ppp/chap-secrets' \
  ipsecConfigFilepath='/ipsec-etc/ipsec.d/passwd' \
  ipsecSecretsFilepath='/ipsec-etc/ipsec.secrets'

CMD [ "/ipsec-api" ]
