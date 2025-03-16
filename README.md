# go-mmuc

It is MMUC(Multimedia Multi User Chat) application in golang.

I want to build a chat room based on WebRTC and JMPP(Json Messaging and Presence Protocol) protocol based on XMPP protocol.

refer to based on https://xmpp.org/extensions/xep-0045.html

It is not finished, still building...



```shell
go mod init github.com/walterfan/go-mmuc
```

## DB initialization

```sql

-- 创建数据库 mmuc
CREATE DATABASE IF NOT EXISTS mmuc;

-- 创建用户 walter 并设置密码
CREATE USER IF NOT EXISTS 'your_username'@'localhost' IDENTIFIED BY 'your_password';

-- 授予用户 walter 对数据库 mmuc 的所有权限
GRANT ALL PRIVILEGES ON mmuc.* TO 'your_username'@'localhost';

-- 刷新权限以使更改生效
FLUSH PRIVILEGES;
```

* create a file - .env

```
DB_USER=your_db_username
DB_PWD=your_db_password

```