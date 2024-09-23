package main

import "strconv"

type Todo struct {
	Id          int
	Task        string
	IsCompleted bool
}

type TodoList struct {
	TodoList []Todo
}

func (tl *TodoList) List(cnt int) {
	todoList := make([]Todo, cnt)
	for i, _ := range todoList {
		id := i + 1
		task := "タスク No." + strconv.Itoa(id)
		newTodo := Todo{
			Id:          id,
			Task:        task,
			IsCompleted: false,
		}
		todoList[i] = newTodo
	}
	tl.TodoList = todoList
}
