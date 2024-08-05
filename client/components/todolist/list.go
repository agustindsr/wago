package todolist

import (
	"wasm/client/wa/dom"
	"wasm/client/wa/dom/tailwind"
)

type Task struct {
	Text string
}

var tasks []Task

type Component struct {
	Input  dom.HTMLNode
	Button dom.HTMLNode
	List   dom.HTMLNode
}

func Render() dom.HTMLNode {
	component := Component{
		Input: dom.Input().SetPlaceholder("Add a new task").
			Tailwind(tlw.P2, tlw.Border, tlw.BorderGray300, tlw.RoundedMd, tlw.Mb2),
		Button: dom.Button("Add Task").
			Tailwind(tlw.P2, tlw.BgBlue500, tlw.TextWhite, tlw.RoundedMd, tlw.HoverBgBlue700),
		List: dom.UL().
			Tailwind(tlw.ListNone, tlw.P0, tlw.Mt2),
	}

	component.Button.OnClick(component.AddTask)
	component.Input.OnKeyUp(component.AddTaskOnEnter)

	container := dom.Div().Tailwind(tlw.MxAuto, tlw.P4, tlw.BgWhite, tlw.ShadowMd, tlw.RoundedLg).
		Child(
			dom.H2("Todo List").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			dom.Div().Child(component.Input, component.Button).Tailwind(tlw.Flex, tlw.SpaceX2),
			component.List,
		)

	return container
}

func (c *Component) AddTask(_ dom.Event) {
	text := c.Input.GetValue().String()
	if text == "" {
		return
	}

	tasks = append(tasks, Task{Text: text})

	listItem := dom.LI().SetInnerHTML(text).
		Tailwind(tlw.P2, tlw.Mb2, tlw.Border, tlw.BorderGray300, tlw.RoundedMd, tlw.Flex, tlw.ItemsCenter, tlw.JustifyBetween)

	deleteButton := dom.Button("Delete").
		OnClick(c.DeleteTask(listItem, text)).
		Tailwind(tlw.P2, tlw.BgRed500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgRed700)

	listItem.Child(deleteButton)
	c.List.Child(listItem)

	c.Input.SetValue("")
}

func (c *Component) AddTaskOnEnter(e dom.Event) {
	if e.IsKeyCodeEnter() {
		c.AddTask(e)
	}
}

func (c *Component) DeleteTask(listItem dom.HTMLNode, text string) dom.Func {
	return func(_ dom.Event) {
		for i, item := range tasks {
			if item.Text == text {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}

		c.List.RemoveChild(listItem)
	}
}
