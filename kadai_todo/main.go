package main

import "fmt"

// Step3. ToDoアイテムの構造体を定義
type Todo struct {
	ID        int
	Title     string
	Completed bool
}

// Step4. ToDoを完了状態にするメソッド
func (t *Todo) Complete() {
	fmt.Printf("ID：%dのToDoを完了に更新します\n", t.ID)
	t.Completed = true
}

// Step5. ToDoリストを一覧表示する関数
func printTodos(todos []Todo) {
	for _, todo := range todos {
		if todo.Completed {
			fmt.Printf("[完　了]（ID：%d）%s\n", todo.ID, todo.Title)
		} else {
			fmt.Printf("[未完了]（ID：%d）%s\n", todo.ID, todo.Title)
		}
	}
}

// Step6. メイン処理
func main() {
	// ToDo型のスライスを作成し、3件のデータを追加
	todos := []Todo{
		{ID: 1, Title: "学習計画", Completed: false},
		{ID: 2, Title: "環境構築", Completed: false},
		{ID: 3, Title: "基礎文法", Completed: false},
	}

	// 初期状態のToDoリストを表示
	fmt.Println("--- ToDoリスト（初期状態） ---")
	printTodos(todos)

	fmt.Println()

	// IDが1のToDoを完了状態に更新
	todos[0].Complete()

	// IDが2のToDoを完了状態に更新
	todos[1].Complete()

	fmt.Println()

	// 最終状態のToDoリストを表示
	fmt.Println("--- ToDoリスト（最終状態） ---")
	printTodos(todos)
}
