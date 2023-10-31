# ChatTKBのリポジトリ

このリポジトリは筑波大学SOHO祭の特別講演で使用するChatTKBのソースコードを格納する場所です。

## ディレクトリ構成

本リポジトリはいわゆるモノレポ構成になっています。
ローカル環境ではDockerを用いることでフロント/バックエンドの開発環境を構築することが出来ます。
バックエンドはDockerで起動し、フロントエンドはviteの開発サーバーを立ち上げてください。


```shell
docker compose build
docker compose up -d
```

## 本番環境での実行

本番環境では以下のコマンドで起動してください。なお、初回起動時には`nginx.conf`の443関連の設定を落としてください。
参考に初回起動時（certbot実行前）のnginx.confをexampleとして格納しています。

```shell
docker compose -f compose.prod.yml build
docekr compose -f compose.prod.yml up -d
```

## 参考文献

- https://github.com/Carriage-Horse-Technologies/masakari-frontend/blob/main/src/utils/ChatService.ts
- https://github.com/Carriage-Horse-Technologies/masakari-backend/blob/main/api/src/redis/redis.go
- https://qiita.com/mttt/items/aa2ba3a0677a803d0436


## TEST（バックエンド）

管理用メッセージ送信オブジェクト

```json
{"action":"chat_message","message":"hogehoge"}

```

## certbotの作成手順

```shell
docker compose -f compose.prod.yml exec /bin/sh
```

コンテナにログイン後、以下のコマンドを実行する


```shell 
certbot certonly --webroot -w /usr/share/nginx/html -d api-chat-tkb.crowd4u.org
```
