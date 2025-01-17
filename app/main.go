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

func myHook(ctx context.Context, exec boil.ContextExecutor, todo *database.Todo) error {
	fmt.Printf("todo title %v", todo.Title)
	return nil
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

	// ctx := context.Background()

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

	/*
		// Bulk Insert
		var todosSlice database.TodoSlice
		todosSlice = append(todosSlice, &database.Todo{
			Title:   "test_bulk_insert_title_1",
			Content: null.String{String: "test_bulk_insert_content_1", Valid: true},
			UserID:  1,
		})
		todosSlice = append(todosSlice, &database.Todo{
			Title:   "test_bulk_insert_title_1",
			Content: null.String{String: "test_bulk_insert_content_2", Valid: true},
			UserID:  1,
		})
		todosSlice.InsertAll(ctx, db, boil.Infer())
	*/

	/*
		// Update
		user, _ := database.FindUser(ctx, db, 1)
		user.Name = "updated_name_1"
		rowAff, err := user.Update(ctx, db, boil.Infer())
		fmt.Println(rowAff)
	*/

	/*
		// Update All
		// users, _ := database.Users().All(ctx, db)
		// rowsAff, _ := users.UpdateAll(ctx, db, database.M{"Name": "update_all_name"})
		// fmt.Println(rowsAff)

		// rowsAff, _ := database.Users().UpdateAll(ctx, db, database.M{"Name": "updated_all_name"})
		// fmt.Println(rowsAff)
	*/

	/*
		// Upsert
		var todosSlice database.TodoSlice
		limit := 1
		offset := 0

		for {
			todos, err := database.Todos(
				qm.Limit(limit),
				qm.Offset(offset),
			).All(ctx, db)
			if err != nil {
				log.Fatalf("Error fetching records: %v", err)
			}

			if len(todos) == 0 {
				break // no more records to fetch
			}

			for _, todo := range todos {
				todo.Title = todo.Title + "_upserted"
				todosSlice = append(todosSlice, todo)
			}

			offset += limit
		}
		todosSlice.UpsertAll(ctx, db, boil.Infer(), boil.Infer())
	*/

	/*
		// Delete
		todo := &database.Todo{
			Title:   "test_title_delete",
			Content: null.String{String: "test_content_delete", Valid: true},
			UserID:  1,
		}
		createTodoErr := todo.Insert(ctx, db, boil.Infer())
		if createTodoErr != nil {
			log.Fatalln(createTodoErr)
		}

		rowsAff, deleteErr := todo.Delete(ctx, db)
		if deleteErr != nil {
			log.Fatalln(deleteErr)
		}
		fmt.Println(rowsAff)
	*/

	/*
		// Delete All
		var todosSlice database.TodoSlice
		todosSlice = append(todosSlice, &database.Todo{
			Title:   "test_delete_title",
			Content: null.String{String: "test_delete_content_1", Valid: true},
			UserID:  1,
		})
		todosSlice = append(todosSlice, &database.Todo{
			Title:   "test_delete_title",
			Content: null.String{String: "test_delete_content_2", Valid: true},
			UserID:  1,
		})
		todosSlice.InsertAll(ctx, db, boil.Infer())

		rowsAff, deleteAllErr := database.Todos(qm.Where("title = ?", "test_delete_title")).DeleteAll(ctx, db)
		if deleteAllErr != nil {
			log.Fatalln(deleteAllErr)
		}
		fmt.Println(rowsAff)
	*/

	/*
		// Dependent delete
		tx, txErr := db.BeginTx(ctx, nil)
		if txErr != nil {
			log.Fatalln(txErr)
		}

		users, _ := database.Users(qm.Where("id = ?", 1)).All(ctx, db)
		for _, user := range users {
			// メタプロっぽくはできなさそう
			// tv := reflect.TypeOf(user.R)
			// for i := 0; i < tv.NumField(); i++ {
			// 	t := tv.Field(i)
			// 	t.DeleteAll(ctx, db)
			// }
			if _, err := user.Todos().DeleteAll(ctx, db); err != nil {
				tx.Rollback()
				break
			}

			if _, err := user.Delete(ctx, db); err != nil {
				tx.Rollback()
				break
			}
		}
		tx.Commit()
	*/

	/*
		// Hooks
		database.AddTodoHook(boil.AfterInsertHook, myHook)
		todo := &database.Todo{
			Title:   "test_title_after_insert_hook",
			Content: null.String{String: "test_content_after_insert_hook", Valid: true},
			UserID:  2,
		}
		createTodoErr := todo.Insert(ctx, db, boil.Infer())
		if createTodoErr != nil {
			log.Fatalln(createTodoErr)
		}
	*/
}

func ints64ToInterfaces(nums []int64) []interface{} {
	results := make([]interface{}, len(nums))

	for i, n := range nums {
		results[i] = n
	}

	return results
}
