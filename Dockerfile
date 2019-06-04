FROM golang:1.12.5@sha256:cf0b9f69ad1edd652a7f74a1586080b15bf6f688c545044407e28805066ef2cb as Build
RUN CGO_ENABLED=0 go get -u github.com/shynome/ipsec-api/cmd/ipsec-api

FROM alpine:3.9.4@sha256:769fddc7cc2f0a1c35abb2f91432e8beecf83916c421420e6a6da9f8975464b6
COPY --from=Build /go/bin/ipsec-api /ipsec-api

# ENV \
# LDAP_Host='ldaps://your.company.com:636' \
# LDAP_BaseDN='ou=users,dc=company,dc=com' \
# LDAP_Filter='(objectclass=inetOrgPerson)' \
# LDAP_Attr='cn' \
# LDAP_BindDN='cn=xxx-read,ou=apps,dc=company,dc=com' \
# LDAP_Password='LDAP_BindDN password' \
# token='5555555555555'

ENV \
  l2tpdCoonfigFilepath='/ipsec-etc/ppp/chap-secrets' \
  ipsecConfigFilepath='/ipsec-etc/ipsec.d/passwd' \
  ipsecSecretsFilepath='/ipsec-etc/ipsec.secrets'

CMD [ "/ipsec-api" ]
