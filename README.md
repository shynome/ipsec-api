### 简介

`ldap` 和 `ipsec-vpn` 用户同步的 http 接口

### 环境变量必须

```sh
docker run --rm -ti \
  -p 7070:7070 \
  -e LDAP_Host='ldaps://your.company.com:636' \
  -e LDAP_BaseDN='ou=users,dc=company,dc=com' \
  -e LDAP_Filter='(objectclass=inetOrgPerson)' \
  -e LDAP_Attr='cn' \
  -e LDAP_BindDN='cn=xxx-read,ou=apps,dc=company,dc=com' \
  -e LDAP_Password='LDAP_BindDN password' \
  -e token='32345y6hy02uhh9049t43g' \
  # 挂载 ipsec vpn 配置文件, 以供修改
  -v /root/docker-ipsec-vpn-server-master/etc:/ipsec-etc/ \
  shynome/ipsec-api:0.0.2
```
