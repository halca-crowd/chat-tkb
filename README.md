# ChatTKBのリポジトリ

このリポジトリは筑波大学SOHO祭の特別講演で使用するChatTKBのソースコードを格納する場所です。

## ディレクトリ構成

本リポジトリはいわゆるモノレポ構成になっています。
ローカル環境ではDockerを用いることでフロント/バックエンドの開発環境を構築することが出来ます。


## 参考文献

- https://github.com/Carriage-Horse-Technologies/masakari-frontend/blob/main/src/utils/ChatService.ts
- https://github.com/Carriage-Horse-Technologies/masakari-backend/blob/main/api/src/redis/redis.go

## TEST


管理用メッセージ送信オブジェクト

```json
{"action":"chat_message","message":"hogehoge"}

```

certbot

```shell
docker compose -f compose.prod.yml exec /bin/sh
```
コンテナにログイン後、以下のコマンドを実行する
```shell 
certbot certonly --webroot -w / -d api-chat-tkb.crowd4u.org
```