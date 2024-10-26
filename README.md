# go-sqlboiler-practice
GoのSQL Boilerの検証

- sqlboiler
	- https://github.com/volatiletech/sqlboiler
	- https://zenn.dev/sagae/articles/c6b2e460201d31
	- https://zenn.dev/jy8752/books/73769005e6afa9/viewer/chapter5
	- https://zenn.dev/ryomak/articles/sqlboiler-go
	- マイグレーションは別で行う必要がある
	- 拡張ライブラリがある
		- https://github.com/tiendc/sqlboiler-extensions
		- bulk insertやupsertを扱えるらしい
	- sqlboiler mysql
        - https://zenn.dev/gami/articles/0fb2cf8b36aa09
        - https://qiita.com/ishihaya/items/38520c5c11f31e382ecf
        - https://qiita.com/piggydev/items/831b41462bcadfc34cf5

- マイグレーション
	- sql-migrate
		- sqlboilerでreferenceとして書かれていた
		- https://github.com/rubenv/sql-migrate
		- https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
		- https://qiita.com/hideki_okawa/items/25d2f4e751bb202a06e8
		- sql-migrate new -env="mysql" コマンド

## 使ってみた所管
- 大体railsのActiveRecordのような感覚で使用できる
  - ORMでやりたいCRUD操作はほとんどできる
    - bulk insertやupsert, delete_allを使いたければsqlboiler-extensionsと合わせて使用することになる
  - hooksも使える
  - dependent: :destoryはなしのため、自分でトランザクションを込みでメソッドを作る必要がある
- 自動生成されるコードによる型が使えて、型安全にやりやすそう
- GORMのようなinterface型で書くことが少なくなり、実行時エラーを防ぎやすそう
- ライブラリも継続的に更新されていて、ドキュメントも詳しく、参考URLも豊富
- マイグレーションツールは別途用意が必要(sql-migrationが公式で紹介されていたし、それで良さそう)
