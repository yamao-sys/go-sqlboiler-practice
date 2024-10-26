package main

import (
	database "app/database/generated"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type BindUser struct {
	database.User `boil:",bind"`
}

type BindTodoFields struct {
	Title   string
	Content string
}

type BindTodo struct {
	BindTodoFields `boil:",bind"`
}

func main() {
	fmt.Println(os.Getenv("MYSQL_DBNAME"))

	dsn := os.Getenv("MYSQL_USER") +
		":" + os.Getenv("MYSQL_PASS") +
		"@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" +
		os.Getenv("MYSQL_DBNAME") +
		"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	ctx := context.Background()

	// NOTE: SQL文を標準出力に出す
	boil.DebugMode = true

	/*
		// CREATE
		hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}
		user := &database.User{
			Name:     "test_name_2",
			Email:    "test_2@example.com",
			Password: string(hash),
		}
		createErr := user.Insert(ctx, db, boil.Infer())
		if createErr != nil {
			log.Fatalln(createErr)
		}
	*/

	/*
		// Select
		// 条件に合った単一レコード
		u, err := database.Users(
			qm.Where("id = ?", 1),
			qm.And("email = ?", "test@example.com"),
		).One(ctx, db)
		fmt.Println(u)

		us, err := database.Users(
			database.UserWhere.ID.EQ(1),
		).One(ctx, db)
		fmt.Println(us)
	*/

	/*
		// Primary IDによる抽出
		u, err = database.FindUser(ctx, db, 1)
		fmt.Println(u)
	*/

	/*
		// IN句
		// like rails find_each
		limit := 1
		offset := 0

		for {
			users, err := database.Users(
				qm.WhereIn("id IN ?", ints64ToInterfaces([]int64{1, 2})...),
				qm.Limit(limit),
				qm.Offset(offset),
			).All(ctx, db)
			if err != nil {
				log.Fatalf("Error fetching records: %v", err)
			}

			if len(users) == 0 {
				break // no more records to fetch
			}

			for _, user := range users {
				fmt.Println(user) // process each user
			}

			offset += limit
		}
	*/

	/*
		// 存在チェック
		isUserExists, err := database.Users(
			database.UserWhere.ID.EQ(1),
		).Exists(ctx, db)
		fmt.Println(isUserExists)
	*/

	/*
		// 指定カラムのみの取得
		users, err := database.Users(
			qm.Select("id", "name"),
		).All(ctx, db)
		for _, user := range users {
			fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
		}
	*/

	/*
		// Todo CREATE
		todo := &database.Todo{
			Title:   "test_title_1",
			Content: null.String{String: "test_content_1", Valid: true},
			UserID:  1,
		}
		createTodoErr := todo.Insert(ctx, db, boil.Infer())
		if createTodoErr != nil {
			log.Fatalln(createTodoErr)
		}
	*/

	/*
		// JOIN
		joinedUsers, err := database.Users(
			qm.InnerJoin(database.TableNames.Todos+" on "+database.TableNames.Todos+"."+database.TodoColumns.UserID+" = "+database.TableNames.Users+"."+database.UserColumns.ID),
			qm.Where(database.TodoColumns.Title+" = ?", "test_title_1"),
		).All(ctx, db)
		for _, user := range joinedUsers {
			fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
		}
	*/

	/*
		// EagerLoad
		eagerLoadedUsers, err := database.Users(
			qm.Load(qm.Rels(database.UserRels.Todos)),
		).All(ctx, db)
		for _, user := range eagerLoadedUsers {
			fmt.Println(user)
		}
	*/

	/*
		// RawSQL
		var bindUser BindUser
		bindSelectErr := queries.Raw(`
			select * from users where id = ?
		`, 1).Bind(ctx, db, &bindUser)
		if bindSelectErr != nil {
			log.Fatalln(bindSelectErr)
		}
		fmt.Println(bindUser)

		var bindSelectedTodo BindTodo
		bindSelectTodoErr := queries.Raw(`
			select todos.title, todos.content from users inner join todos on users.id = todos.user_id where todos.user_id = ?
		`, 1).Bind(ctx, db, &bindSelectedTodo)
		if bindSelectTodoErr != nil {
			log.Fatalln(bindSelectTodoErr)
		}
		fmt.Println(bindSelectedTodo)
	*/
}

func ints64ToInterfaces(nums []int64) []interface{} {
	results := make([]interface{}, len(nums))

	for i, n := range nums {
		results[i] = n
	}

	return results
}
