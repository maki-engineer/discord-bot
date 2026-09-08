# 235bot Backend

235bot の PostgreSQL データを取得する Go API です。Frontend に対してメンバー情報を返し、将来的に Bot のデータを Web から管理するための拡張ポイントを提供します。

## 技術構成

- Go
- Gin
- GORM / PostgreSQL Driver
- Swagger
- Docker / Google Cloud Run

コードは、`domain`、`application`、`infrastructure`、`presentation`、`server` に分けた構成です。

## API

↓↓ でAPI定義書を見ることができます。

https://discord-bot-backend-758736590208.asia-northeast1.run.app/swagger/index.html

Basic認証を設定しているため、見るにはIDとパスワードが必要です。見たい場合は管理者に相談してください。

## 開発方法

```bash
go mod download
go run ./src/cmd
```

API は `:8080` で起動します。ルートの Docker Compose では `backend` とテスト用 PostgreSQL が起動し、開発用コンテナは Air によるホットリロードを利用します。

接続先や CORS の許可 URL は環境変数で設定します。ルートの `.env.example` を参考にしてください。

## テスト・Lint

```bash
task test-local
task lint
```

`task swagger` で Swagger ドキュメントを生成できます。API のコメントやルートを変更した場合は、生成された `docs` の差分も確認してください。

## デプロイ

`main` ブランチの Backend 関連変更は GitHub Actions から Google Cloud Build でイメージを作成し、Google Cloud Run の `discord-bot-backend` サービスへデプロイされます。
