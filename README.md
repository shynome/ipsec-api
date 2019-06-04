### 简介
`ldap` 和 `ipsec-vpn` 用户同步的 http 接口

### 环境变量必须

```sh
LDAP_Host='ldaps://your.company.com:636' 
LDAP_BaseDN='ou=users,dc=company,dc=com' 
LDAP_Filter='(objectclass=inetOrgPerson)' 
LDAP_Attr='cn' 
LDAP_BindDN='cn=xxx-read,ou=apps,dc=company,dc=com' 
LDAP_Password='LDAP_BindDN password' 
# 用以接口认证
token='32345y6hy02uhh9049t43g'
```
