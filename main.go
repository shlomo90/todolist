package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/shlomolim90/todolist/todos"
)

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Empty argument.")
		return
	}

	if args[0] == "add" {
		if len(args) != 2 {
			fmt.Println("Invalid argument:", args)
			return
		}

		todos := &todos.Todos{}
		todos.LoadTodos()
		fmt.Printf("str:%s\n", args[1])
		todos.AddTodo(args[1], 1)
		todos.SaveTodos()
	} else if args[0] == "del" {
		if len(args) != 2 {
			fmt.Println("Invalid argument:", args)
			return
		}

		index, _ := strconv.Atoi(args[1])

		todos := &todos.Todos{}
		todos.LoadTodos()
		todos.DelTodo(index)
		todos.SaveTodos()
	} else if args[0] == "list" {
		todos := &todos.Todos{}
		todos.LoadTodos()
		todos.ListTodos()
	}
}
