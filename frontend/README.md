# フロントエンド

事前のnodeとpnpmをインストールしてください。nodenvを使用している場合は.node-versionに指定されているバージョンでインストールしてください。
推奨は20.x以上です。なお、コマンドが見つからない場合はシェルを再起動してください。

```shell
nodenv install 21.1.0
npm install -g pnpm
```

## 開発方法

開発サーバーは以下のコマンドで実行されます。環境変数は`.env.dev`を読みに行くので別途バックエンドを起動しておいてください。

```shell
pnpm install
pnpm run dev
```

## デプロイ方法

netlify-cliをインストールしておいてください。

```shell
npm install -g netlify-cli
```

以下のコマンドでデプロイできます。`--prod`を抜くとプレビュー環境にデプロイされます。

```shell
netlify build
netlify deploy --prod
```
