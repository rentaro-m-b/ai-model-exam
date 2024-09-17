# AI Model コーディング試験用レポジトリ

## 概要
書籍管理APIサーバの実装
- GET /books -> 全ての書籍情報の一覧を返す
- POST /books -> 書籍情報を登録する

## 環境構築
1. レポジトリのクローン
```bash
git clone `本レポジトリ`
```

2. dockerコンテナの起動
```bash
docker compose up -d
```

3. golang-migrateのインストールとマイグレーションの実行
```bash
brew install golang-migrate
make migrate-up
```

4. レコードの挿入
```bash
docker compose exec -it postgres psql -U <username> -d <dbname>

INSERT INTO books (id, title, author, publisher, price) VALUES (nextval('BOOK_ID_SEQ'), 'テスト駆動開発', 'Kent Beck', 'オーム社', 3080);
INSERT INTO books (id, title, author, publisher, price) VALUES (nextval('BOOK_ID_SEQ'), 'アジャイルサムライ', 'Jonathan Rasmusson', 'オーム社', 2860);
INSERT INTO books (id, title, author, publisher, price) VALUES (nextval('BOOK_ID_SEQ'), 'エクストリームプログラミング', 'Kent Beck', 'オーム社', 2420);
INSERT INTO books (id, title, author, publisher, price) VALUES (nextval('BOOK_ID_SEQ'), 'Clean Agile', 'Robert C. Martin', 'ドワンゴ', 2640);
```

5. （付録）sqlcのインストール
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

6. （付録）mockgenのインストール
```bash
go install github.com/golang/mock/mockgen@latest
```
