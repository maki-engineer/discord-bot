# 235bot Frontend

235bot のデータをブラウザで確認するための Next.js アプリケーションです。現在は、誕生月を選択して該当するメンバー一覧を表示できます。

## 技術構成

- Next.js 16（App Router）
- React 19
- TypeScript
- Tailwind CSS 4

## 開発方法

```bash
npm install
npm run dev
```

http://localhost:3000 を開いてください。ルートの Docker Compose を使う場合は、リポジトリのルートで `docker compose up frontend` を実行します。

## 環境変数

`.env.local` に Backend API の URL を設定します。

```env
NEXT_PUBLIC_API_BASE_URL=http://localhost:8080/api
```

未設定の場合も、開発時は `http://localhost:8080/api` が既定値として使われます。

## 画面とデータ取得

- `app/page.tsx`: トップページ
- `app/birthday/members/page.tsx`: 誕生日メンバー画面
- `components/BirthdayMonthSelect.tsx`: 誕生月選択
- `components/BirthdayMemberTable.tsx`: メンバー一覧
- `lib/api/get-birthday-members.ts`: Backend の `GET /api/members?birthday_month={month}` 呼び出し

API がエラーを返した場合は、画面に取得失敗を表示し、一覧を空に戻します。

## コマンド

```bash
npm run dev      # 開発サーバー
npm run build    # 本番ビルド
npm run start    # 本番ビルドの起動

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.
