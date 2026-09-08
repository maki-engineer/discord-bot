# Repository Automation

このフォルダには、開発フローを自動化する GitHub Actions と Pull Request のテンプレートを置いています。

## Workflows

- `test-for-backend.yml`: Backend のテスト
- `linter-check-for-backend.yml`: Backend の Lint チェック
- `cloud-run-deploy.yml`: `main` ブランチの Backend 変更を Google Cloud Run へデプロイ
- `bot-run.yml`: Discord Bot の定期実行

`cloud-run-deploy.yml` では Swagger ドキュメントの生成、Docker イメージのビルド、Cloud Run へのデプロイを行います。Google Cloud の認証情報は GitHub Secrets から渡し、ワークフローや README に直接記載しません。

Issue と Pull Request のテンプレートもこのフォルダで管理しています。新しい自動化を追加する場合は、対象パスを限定し、必要な権限と Secrets を最小限にしてください。
