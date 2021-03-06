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
  # token 鉴权必需, 需要自行生成
  -e token='generate uuid by yourself' \
  # 挂载 ipsec vpn 配置文件, PWD 是 vpn 所在的路径
  -v "$PWD/etc/ipsec.secrets:/xxx/ipsec.secrets" \
  -v "$PWD/etc/ppp/chap-secrets:/xxx/ppp/chap-secrets" \
  -v "$PWD/etc/ipsec.d/passwd:/xxx/ipsec.d/passwd" \
  shynome/ipsec-api
```

查看 ldap 服务器是否正常

```
curl -s -H 'token: generate uuid by yourself' http://127.0.0.1:7070/ldap/list
```

### 接口

| 接口名              | 参数               | 返回格式                         | 说明                                                                                                                                      |
| ------------------- | ------------------ | -------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `/user/list`        | `{}`               | `{users:string[]}`               | 获取服务器上的用户                                                                                                                        |
| `/user/getpassword` | `{user:string}`    | `{password:string}`              | 获取服务器上的用户密码                                                                                                                    |
| `/user/add`         | `{users:string[]}` | `{users:string[]}`               | 添加用户                                                                                                                                  |
| `/user/sync`        | `{confirm?:true}`  | `{add:string[],delete:stirng[]}` | 以 ldap 为基准同步用户, 不存在于 ldap 的用户会被删除, 不存在于服务器的用户会被添加, 传 `confirm` 才会执行操作, 不传只列出将要被操作的用户 |
| `/ldap/list`        | `{}`               | `{user:string[]}`                | 列出 ldap 上的用户, 同时可用来检测 ldap 服务是否可用                                                                                      |
