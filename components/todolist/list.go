package todolist

import (
	"syscall/js"
	"wasm/wa/dom"
)

type Task struct {
	Text string
}

var tasks []Task

func Render() dom.DivElement {
	todoContainer := dom.Div()
	todoContainer.AddClass("todo-container")

	todoTitle := dom.H2().SetInnerText("Todo List")
	todoContainer.Child(todoTitle)

	todoInput := dom.Input()
	todoInput.SetAttribute("placeholder", "Add a new task").AddClass("todo-input")
	todoButton := dom.Button().SetInnerText("Add Task").AddClass("todo-button")
	todoList := dom.UL()
	todoList.AddClass("todo-list").SetAttribute("id", "todo-list")

	todoContainer.Child(todoInput.HTMLNode)
	todoContainer.Child(todoButton)
	todoContainer.Child(todoList.HTMLNode)

	todoButton.OnClick(addTask(todoInput, todoList))

	return todoContainer
}

func addTask(input dom.InputElement, list dom.ULElement) dom.Func {
	return func(this js.Value, p []js.Value) any {
		taskText := input.Get("value").String()
		if taskText == "" {
			return nil
		}

		// Añadir tarea al slice de tareas
		tasks = append(tasks, Task{Text: taskText})

		// Renderizar lista de tareas
		renderTaskList(list)

		// Limpiar input
		input.Set("value", "")

		return nil
	}
}

func deleteTask(index int, list dom.ULElement) dom.Func {
	return func(this js.Value, p []js.Value) any {
		// Eliminar tarea del slice de tareas
		tasks = append(tasks[:index], tasks[index+1:]...)

		// Renderizar lista de tareas
		renderTaskList(list)

		return nil
	}
}

func renderTaskList(list dom.ULElement) {
	// Limpiar lista existente
	list.SetInnerHTML("")

	// Renderizar tareas desde el slice
	for i, task := range tasks {
		listItem := dom.LI()
		listItem.SetInnerText(task.Text).AddClass("todo-item")

		deleteButton := dom.Button().SetInnerText("Delete").AddClass("delete-button")
		listItem.Child(deleteButton)

		deleteButton.OnClick(deleteTask(i, list))

		list.Child(listItem.HTMLNode)
	}
}
