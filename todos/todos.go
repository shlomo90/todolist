/*
*/


package todos

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"
	"time"
)

const (
	UNKNOWN int = iota
	OPENED int = iota
	CLOSED int = iota

)

var Unknown int = UNKNOWN
var Opened int = OPENED
var Closed int = CLOSED

const (
	TIMESTAMP = iota
	CONTENT = iota
	STATE = iota
	TOTAL = iota
)

type Todo struct {
	timestamp int
	content string
	state int
}

type Todos []Todo

type TodosInterface interface {
	LoadTodos()
	SaveTodos()
	ListTodos()
	AddTodo(content string, state int)
	DelTodo(index int)
	GetTodo(index int) Todo
}

func StateConvertor(state int) string {
	if state == OPENED {
		return "OPEN"
	} else if state == CLOSED {
		return "CLOSE"
	} else {
		return "UNKNOWN"
	}
}

func (t *Todos) LoadTodos() {
	rf, err := os.Open("/root/.todos")
	if err != nil {
		t.SaveTodos()
		rf, err = os.Open("/root/.todos")
		if err != nil {
			fmt.Printf("Err: open Error\n")
			return
		}
	}

	defer rf.Close()

	gen := t.todo_file_reader(rf)
	for s := gen(); s != ""; s = gen() {
		split := strings.Split(s, "\t")
		if len(split) != TOTAL {
			fmt.Printf("Err: Parsing Error. token is not %d.\n", TOTAL)
			return
		}

		todo := Todo{}
		if timestamp, err := strconv.Atoi(split[TIMESTAMP]); err == nil {
			todo.timestamp = timestamp
		} else {
			fmt.Println("Err timestamp is not int type.")
			return
		}

		if state, err := strconv.Atoi(split[STATE]); err == nil {
			todo.state = int(state)
		} else {
			fmt.Println("Err state is not int type.")
			return
		}

		todo.content = split[CONTENT]
		*t = append(*t, todo)
	}
}

func (t *Todos) SaveTodos() {
	wf, err := os.Create("/root/.todos")
	if err != nil {
		fmt.Printf("Err: open Error\n")
		return
	}

	defer wf.Close()

	t.todo_file_writer(wf)
	return
}

func (t *Todos) AddTodo(content string, state int) {
	now := time.Now()
	timestamp := int(now.Unix())
	todo := Todo{
		//timestamp: time.Now().Unix(),
		timestamp: timestamp,
		content: content,
		state: state,
	}
	*t = append(*t, todo)
}

func (t *Todos) DelTodo(index int) {

}

func (t *Todos) GetTodo(index int) Todo {
	var last Todo
	for _, todo := range(*t) {
		last = todo
	}
	return last
}

func (t *Todos) ListTodos() {
	for i, todo := range(*t) {
		fmt.Printf("%d. %s (state:%s, time:%d)\n", i, todo.content, StateConvertor(todo.state), todo.timestamp)
	}
}

func (t *Todos) todo_file_reader(r io.Reader) func() string {
	scanner := bufio.NewScanner(r)
	return func () string {
		if scanner.Scan() {
			return scanner.Text()
		}
		return ""
	}
}

func (t *Todos) todo_file_writer(w io.Writer) {
	for _, todo := range(*t) {
		if _, err := fmt.Fprintf(w, "%d\t%s\t%d\n", todo.timestamp, todo.content, todo.state); err != nil {
			fmt.Println("Err: ", err)
		}
	}
}
