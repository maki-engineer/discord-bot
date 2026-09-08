# Infrastructure

Backend のデプロイ先である Google Cloud のインフラストラクチャを、Terraform でコード管理するためのフォルダです。

現時点では Terraform のリソースコードは未実装です。まず Cloud Run を中心に必要な Google Cloud リソースを定義し、手作業による設定との差分をコードで確認できる状態を目指します。

## 管理対象の予定

Backend の運用に必要な次のリソースを、依存関係が分かる形で管理します。

- Cloud Run のサービス、リージョン、コンテナ設定、環境変数
- Artifact Registry の Docker リポジトリ
- Cloud Build に必要な API とサービスアカウント
- Cloud Run の実行用サービスアカウントと IAM 権限
- GitHub Actions からのデプロイに使う Workload Identity Federation
- Terraform の状態を保存する Google Cloud Storage バケット

データベースのパスワード、Discord のトークン、Gemini API キーなどの秘密情報は Terraform のコードや `.tfvars` に直接書きません。Secret Manager への保存と Cloud Run からの参照を基本方針とします。

## 想定ディレクトリ構成

```text
infra/
├── README.md
├── versions.tf       # Terraform と Provider のバージョン
├── providers.tf      # Google Provider
├── variables.tf      # 入力値
├── main.tf           # リソース定義
├── outputs.tf        # 作成したリソースの出力
├── iam.tf            # サービスアカウントと IAM
├── cloud-run.tf      # Cloud Run
├── artifact-registry.tf
├── backend.tf        # Remote State（GCS）
└── envs/
	└── production.tfvars
```

実装時は、プロジェクト ID やリージョンなど環境ごとに変わる値を変数化します。Terraform の State には機密情報が含まれる可能性があるため、Remote State のバケットはアクセス権を限定し、バージョニングを有効にします。

## 現在のデプロイ構成

Terraform 移行前の Backend デプロイは、`.github/workflows/cloud-run-deploy.yml` で行っています。

1. `main` ブランチの Backend 変更を検知
2. Swagger ドキュメントを生成
3. Cloud Build で `backend/Dockerfile.prod` からイメージをビルド
4. Artifact Registry の `backend:latest` にイメージを登録
5. `asia-northeast1` の Cloud Run サービス `discord-bot-backend` を更新

Terraform で Cloud Run を管理し始めた後は、Terraform が管理する設定と GitHub Actions が更新するコンテナイメージの責務を分けます。Terraform の Apply と通常のアプリケーションデプロイが同じ設定を上書きしないよう、移行時にワークフローの役割を見直します。

## Terraform の基本コマンド

Terraform コードを追加した後は、`infra` フォルダで実行します。

```bash
terraform init
terraform fmt -check
terraform validate
terraform plan -var-file=envs/production.tfvars
terraform apply -var-file=envs/production.tfvars
```

初回導入時は、必ず `plan` の内容と対象プロジェクト・リージョンを確認してから `apply` します。`terraform destroy` は Cloud Run や Artifact Registry を削除するため、通常の開発手順では実行しません。

## 認証と権限

ローカルでは Google Cloud CLI の Application Default Credentials を利用します。

```bash
gcloud auth application-default login
gcloud config set project discord-bot-507515
```

CI では長期保存したサービスアカウントキーを使わず、既存の GitHub Actions と同じ Workload Identity Federation を利用します。Terraform 用の権限は、プロジェクト全体の管理者権限ではなく、管理対象リソースの作成・更新に必要な最小権限へ限定します。

## 実装時の確認事項

- 既存の Cloud Run サービスや Artifact Registry を Terraform に Import してから管理を開始する
- `terraform plan` で既存リソースの意図しない再作成がないことを確認する
- State バケットを Terraform 自身で作る場合は、最初の State だけ安全に手動作成するか、別の Bootstrap 構成に分ける
- 本番の `apply` は Pull Request で Plan を確認してから実行する
- State、サービスアカウント鍵、`.tfvars` の秘密情報を Git に追加しない
