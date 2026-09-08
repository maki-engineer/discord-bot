# 235bot Discord App

235プロダクションの Discord サーバーで動作する Bot 本体です。Discord.js を使ったイベント処理、スラッシュコマンド、定期的なお祝いメッセージ、音声機能を TypeScript で実装しています。

## 主な機能

- 235bot のヘルプと各種スラッシュコマンド
- 誕生日、記念日、ライブ関連のメッセージ送信
- 235士官学校などの日程文章の作成
- ボイスチャンネルの人数に応じた部屋分割
- VOICEVOX Engine を使った音声読み上げ
- Gemini API と Matsurihime API の利用
- PostgreSQL に保存したメンバー・辞書・メッセージ情報の読み書き

### コマンド

- `235help`: 現在利用できるコマンド一覧を表示
- `235birthday`: オンライン飲み会の企画文章を作成。一部のメンバーのみ利用可能
- `235men`: 235士官学校の日程文章を作成。一部のメンバーのみ利用可能
- `235roomdivision`: 特定のボイスチャンネルの参加人数が10人以上のときに部屋を分割

`235birthday` は開催月・日程・時間の3つの半角数字を入力します。例: `235birthday 12 14 21`

`235men` は開催候補日を2～10個の半角数字で入力します。例: `235men 12 14 16 17`

詳細なコマンド一覧は Discord 上で `235help` を実行してください。

## 技術構成

- Node.js 22
- TypeScript
- discord.js 14
- Sequelize 6 / sequelize-cli
- PostgreSQL
- Jest / ESLint
- Docker

## 稼働時間

GitHub Actions の定期実行時間の上限を考慮し、235bot は5時間55分ごとに停止・再起動しています。目安の稼働時間は次のとおりです。

- 05:00～10:55
- 11:00～16:55
- 17:00～22:55
- 23:00～04:55

各時間帯の開始・終了には若干の遅延が発生する場合があります。

## 開発方法

```bash
npm install
npm run build
npm run start
```

開発中にビルド済み JavaScript を監視して起動する場合は `npm run start:dev` を使います。必要なトークン、サーバー ID、データベース接続情報などは環境変数で設定してください。値の一覧は `config` と `.env.example` を確認し、秘密情報はコミットしないでください。

## データベース

マイグレーションとシーダーはこのフォルダで管理します。

```bash
npm run migrate
npm run seed
```

テスト用 PostgreSQL はリポジトリルートの Docker Compose で起動できます。テストでは `NODE_ENV=unittest` とテスト用 DB の設定が必要です。

## コマンド

```bash
npm test          # マイグレーション後に Jest を実行
npm run build     # TypeScript をビルド
npm run compile   # TypeScript をコンパイル
npm run start     # Bot を起動
npm run start:dev # 開発用起動
```

Discord へのコマンド登録は Bot の Ready 処理から行われます。コマンドやイベントを追加する場合は、既存の `src/discord_bot` 配下の責務ごとの構成に合わせて実装してください。

## CI/CD

- ESLint による Lint と Jest によるテストを実行
- Husky のコミットフックで、コミット時にテスト・Lint を確認
- GitHub Actions の `bot-run.yml` で定期的に Bot を起動
