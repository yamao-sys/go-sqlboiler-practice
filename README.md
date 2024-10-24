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

- マイグレーション
	- sql-migrate
		- sqlboilerでreferenceとして書かれていた
		- https://github.com/rubenv/sql-migrate
		- https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
		- https://qiita.com/hideki_okawa/items/25d2f4e751bb202a06e8
		- sql-migrate new -env="mysql" コマンド
