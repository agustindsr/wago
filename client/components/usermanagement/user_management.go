package usermanagement

import (
	"wasm/client/components/shared/modal"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

type User struct {
	Name  string
	Email string
}

var users []User

type Component struct {
	NameInput      dom.HTMLNode
	EmailInput     dom.HTMLNode
	AddButton      dom.HTMLNode
	UserList       dom.HTMLNode
	NameEditInput  dom.HTMLNode
	EmailEditInput dom.HTMLNode
	SaveButton     dom.HTMLNode
	EditModal      *modal.Modal
	SelectedUser   *User
}

func Render() dom.HTMLNode {
	component := &Component{
		NameInput: dom.Input().SetPlaceholder("Name").
			Tailwind(tlw.Mb4, tlw.P2, tlw.Border, tlw.BorderGray300, tlw.Rounded),
		EmailInput: dom.Input().SetPlaceholder("Email").
			Tailwind(tlw.Mb4, tlw.P2, tlw.Border, tlw.BorderGray300, tlw.Rounded),
		AddButton: dom.Button("Add User").
			Tailwind(tlw.Px4, tlw.Py2, tlw.BgBlue500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgBlue700, tlw.TextSm),
		UserList: dom.UL().
			Tailwind(tlw.ListNone, tlw.P0, tlw.Mt4),
	}

	component.AddButton.OnClick(component.AddUser)

	component.NameEditInput = dom.Input().SetPlaceholder("Name").
		Tailwind(tlw.WFull, tlw.P2, tlw.Border, tlw.RoundedMd, tlw.Mb4)
	component.EmailEditInput = dom.Input().SetPlaceholder("Email").
		Tailwind(tlw.WFull, tlw.P2, tlw.Border, tlw.RoundedMd, tlw.Mb4)
	component.SaveButton = dom.Button("Save").
		Tailwind(tlw.P2, tlw.BgGreen500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgGreen700)

	modalContent := dom.Div().Tailwind(tlw.P4).Child(
		dom.H2("Edit User").Tailwind(tlw.TextXl, tlw.FontBold, tlw.Mb4, tlw.TextSm),
		component.NameEditInput,
		component.EmailEditInput,
		component.SaveButton,
	)

	component.EditModal = modal.NewModal(
		modal.WithTitle("Edit User"),
		modal.WithContent(modalContent),
		modal.WithSize(modal.Medium),
	)

	component.SaveButton.OnClick(component.SaveUser)

	container := dom.Div().Tailwind(tlw.MaxWXl, tlw.P4, tlw.RoundedLg).
		Child(
			dom.H2("User Management").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			dom.Div().Child(component.NameInput, component.EmailInput, component.AddButton).Tailwind(tlw.Flex, tlw.SpaceX2),
			component.UserList,
		)

	for _, user := range users {
		component.addUserToDOM(user.Name, user.Email)
	}

	return container
}

func (c *Component) AddUser(_ dom.Event) {
	name := c.NameInput.GetValue().String()
	email := c.EmailInput.GetValue().String()
	if name == "" {
		c.NameInput.ToggleClass("border-red-500")
	}

	if email == "" {
		c.EmailInput.ToggleClass("border-red-500")
	}

	if name == "" || email == "" {
		return
	}

	user := User{Name: name, Email: email}
	users = append(users, user)

	c.addUserToDOM(name, email)

	c.NameInput.SetValue("")
	c.EmailInput.SetValue("")
}

func (c *Component) addUserToDOM(name, email string) {
	listItem := dom.LI().Tailwind(tlw.P4, tlw.Mb2, tlw.Border, tlw.BorderGray300, tlw.RoundedMd, tlw.Flex, tlw.ItemsCenter, tlw.JustifyBetween)
	nameElem := dom.Span(name)
	emailElem := dom.Span(email)
	editButton := dom.Button("Edit").Tailwind(tlw.P2, tlw.BgYellow500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgYellow700)
	deleteButton := dom.Button("Delete").Tailwind(tlw.P2, tlw.BgRed500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgRed700)

	editButton.OnClick(c.EditUser(nameElem, emailElem))
	deleteButton.OnClick(c.DeleteUser(listItem, name))

	listItem.Child(nameElem, emailElem, editButton, deleteButton)
	c.UserList.Child(listItem)
}

func (c *Component) EditUser(nameElem, emailElem dom.HTMLNode) func(e dom.Event) {
	return func(e dom.Event) {
		name := nameElem.GetInnerHTML()
		email := emailElem.GetInnerHTML()

		for _, user := range users {
			if user.Name == name && user.Email == email {
				c.SelectedUser = &user
				break
			}
		}

		c.NameEditInput.SetValue(name)
		c.EmailEditInput.SetValue(email)

		c.EditModal.Open()
	}
}

func (c *Component) SaveUser(_ dom.Event) {
	if c.SelectedUser != nil {
		newName := c.NameEditInput.GetValue().String()
		newEmail := c.EmailEditInput.GetValue().String()
		if newName != "" && newEmail != "" {
			c.SelectedUser.Name = newName
			c.SelectedUser.Email = newEmail

			c.UserList.SetInnerHTML("")
			for _, user := range users {
				c.addUserToDOM(user.Name, user.Email)
			}
		}
		c.EditModal.Close()
	}
}

func (c *Component) DeleteUser(listItem dom.HTMLNode, name string) func(e dom.Event) {
	return func(e dom.Event) {
		for i, user := range users {
			if user.Name == name {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
		c.UserList.RemoveChild(listItem)
	}
}
