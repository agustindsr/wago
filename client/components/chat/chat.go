package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"

	"nhooyr.io/websocket"
)

type Conn struct {
	wsConn *websocket.Conn
}

type Chat struct {
	MessageList     dom.HTMLNode
	MessageInput    dom.HTMLNode
	SendButton      dom.HTMLNode
	LoginInput      dom.HTMLNode
	LoginButton     dom.HTMLNode
	LoginForm       dom.HTMLNode
	ChatContainer   dom.HTMLNode
	UserList        dom.HTMLNode
	ConnectionState dom.HTMLNode
	Conn            *Conn
	Username        string
	loggedIn        bool
	connected       bool
}

func New() *Chat {
	chat := &Chat{
		LoginInput:   dom.Input().SetPlaceholder("Enter your username").Tailwind(tlw.Mb4, tlw.P2, tlw.Border, tlw.BorderGray300, tlw.Rounded),
		LoginButton:  dom.Button("Login").Tailwind(tlw.Px4, tlw.Py2, tlw.BgBlue500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgBlue700, tlw.TextSm),
		MessageList:  dom.UL().Tailwind(tlw.ListNone, tlw.P0, tlw.Mt4, tlw.HFull, tlw.OverflowYScroll, tlw.Flex1),
		MessageInput: dom.Input().SetPlaceholder("Type a message...").Tailwind(tlw.Mb4, tlw.P2, tlw.Border, tlw.BorderGray300, tlw.Rounded),
		SendButton:   dom.Button("Send").Tailwind(tlw.Px4, tlw.Py2, tlw.BgBlue500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgBlue700, tlw.TextSm),
		UserList:     dom.UL().Tailwind(tlw.ListNone, tlw.P0, tlw.Mt4, tlw.FlexNone, tlw.W48, tlw.BgGray100, tlw.Py2, tlw.Px4, tlw.OverflowYScroll),
		ConnectionState: dom.Div().Tailwind(tlw.TextXl, tlw.FontBold, tlw.Mb4).
			SetInnerHTML("Estableciendo conexión con el servidor..."),
		loggedIn:  false,
		connected: false,
	}

	chat.LoginInput.OnKeyUp(chat.OnLoginKeyUp())
	chat.MessageInput.OnKeyUp(chat.OnMessageKeyUp())

	chat.LoginForm = dom.Div().Tailwind(tlw.MaxWXl, tlw.P4, tlw.BgWhite, tlw.RoundedLg, tlw.Hidden).
		Child(
			dom.H2("Login").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			chat.LoginInput,
			chat.LoginButton,
		)
	chat.LoginButton.OnClick(chat.Login)
	chat.SendButton.OnClick(chat.SendMessage)

	chat.ChatContainer = dom.Div().Tailwind(tlw.P4, tlw.BgWhite, tlw.RoundedLg, tlw.Hidden, tlw.W1_2).
		Child(
			dom.H2("Chat").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			dom.Div().Tailwind(tlw.Flex, tlw.Minh1_2).Child(
				chat.UserList.Tailwind(tlw.W1_4, tlw.HFull, tlw.OverflowYScroll, tlw.MinW1_2),
				dom.Div().Tailwind(tlw.Flex1).Child(
					chat.MessageList,
					dom.Div().Child(chat.MessageInput, chat.SendButton).Tailwind(tlw.Flex, tlw.SpaceX2),
				),
			),
		)

	chat.connectWebSocket()

	return chat
}

func (c *Chat) Render() dom.HTMLNode {
	container := dom.Div().
		Tailwind(tlw.Flex, tlw.FlexCol, tlw.ItemsStart, tlw.JustifyCenter, tlw.WFull, tlw.Minh1_2).
		Child(
			c.ConnectionState,
			c.LoginForm,
			c.ChatContainer,
		)

	c.updateVisibility()

	return container
}

func (c *Chat) updateVisibility() {
	if c.connected {
		c.ConnectionState.Tailwind(tlw.Hidden)
		if c.loggedIn {
			c.LoginForm.Tailwind(tlw.Hidden)
			c.ChatContainer.RemoveClass("hidden")
		} else {
			c.LoginForm.RemoveClass("hidden")
			c.ChatContainer.Tailwind(tlw.Hidden)
		}
	} else {
		c.ConnectionState.RemoveClass("hidden")
		c.LoginForm.Tailwind(tlw.Hidden)
		c.ChatContainer.Tailwind(tlw.Hidden)
	}
}

func (c *Chat) connectWebSocket() {
	conn := c.NewConn()
	c.Conn = conn
	go c.readMessage()
}

func (c *Chat) NewConn() *Conn {
	conn, _, err := websocket.Dial(context.Background(), "ws://localhost:8080/ws", nil)
	if err != nil {
		fmt.Println(err, "ERROR")
		return nil
	}
	c.connected = true
	c.updateVisibility()

	return &Conn{
		wsConn: conn,
	}
}

func (c *Chat) readMessage() {
	defer func() {
		c.Conn.wsConn.Close(websocket.StatusGoingAway, "BYE")
	}()

	for {
		_, payload, err := c.Conn.wsConn.Read(context.Background())
		if err != nil {
			log.Panicf(err.Error())
		}

		var msg map[string]interface{}
		err = json.Unmarshal(payload, &msg)
		if err != nil {
			fmt.Println("Error parsing message:", err)
			continue
		}

		switch msg["type"] {
		case "message":
			message := fmt.Sprintf("%s: %s", msg["username"], msg["message"])
			listItem := dom.LI().Tailwind(tlw.P4, tlw.Mb2, tlw.Border, tlw.BorderGray300, tlw.RoundedMd)
			listItem.SetInnerHTML(message)
			c.MessageList.Child(listItem)
		case "users":
			c.UserList.SetInnerHTML("")
			usernames := msg["users"].([]interface{})
			for _, username := range usernames {
				listItem := dom.LI().Tailwind(tlw.Flex, tlw.ItemsCenter, tlw.Py2).Child(
					dom.Div().Tailwind(tlw.W3, tlw.H3, tlw.BgGreen500, tlw.RoundedFull, tlw.Mr2),
					dom.Span(username.(string)),
				)
				c.UserList.Child(listItem)
			}
		}
	}
}

func (c *Chat) Login(_ dom.Event) {
	username := c.LoginInput.GetValue().String()
	if username == "" {
		return
	}
	c.Username = username
	c.loggedIn = true

	loginMessage := map[string]interface{}{
		"type":     "login",
		"username": c.Username,
	}
	loginMessageJSON, _ := json.Marshal(loginMessage)
	err := c.Conn.wsConn.Write(context.Background(), websocket.MessageText, loginMessageJSON)
	if err != nil {
		fmt.Println("Error sending login message:", err)
	}

	c.updateVisibility()
}

func (c *Chat) SendMessage(_ dom.Event) {
	if c.Conn == nil || c.Conn.wsConn == nil {
		fmt.Println("WebSocket is not connected")
		return
	}

	message := c.MessageInput.Get("value").String()
	if message == "" {
		return
	}

	msg := map[string]interface{}{
		"type":     "message",
		"username": c.Username,
		"message":  message,
	}
	msgJSON, _ := json.Marshal(msg)
	err := c.Conn.wsConn.Write(context.Background(), websocket.MessageText, msgJSON)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
	c.MessageInput.Set("value", "")
}

func (c *Chat) OnLoginKeyUp() dom.Func {
	return func(e dom.Event) {
		if e.IsKeyCodeEnter() {
			c.Login(e)
		}
	}
}

func (c *Chat) OnMessageKeyUp() dom.Func {
	return func(e dom.Event) {
		if e.IsKeyCodeEnter() {
			c.SendMessage(e)
		}
	}
}
