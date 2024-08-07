package performance

import (
	"fmt"
	"math/rand"
	"strconv"
	"syscall/js"
	"time"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
)

type TableRow struct {
	ID    int
	Name  string
	Value int
}

type PerformanceComponent struct {
	Rows     []TableRow
	SortBy   string
	TableRef dom.HTMLNode
}

func NewPerformanceComponent() *PerformanceComponent {
	return &PerformanceComponent{}
}

func (pc *PerformanceComponent) RenderRows() dom.HTMLNode {
	fragment := dom.Document().Call("createDocumentFragment")

	for _, row := range pc.Rows {
		tr := dom.TR()

		tdID := dom.TD().SetInnerHTML(fmt.Sprintf("%d", row.ID))
		tr.Child(tdID)

		tdName := dom.TD().SetInnerHTML(row.Name)
		tr.Child(tdName)

		tdValue := dom.TD().SetInnerHTML(fmt.Sprintf("%d", row.Value))
		tr.Child(tdValue)

		fragment.Child(tr)
	}

	return fragment
}

func (pc *PerformanceComponent) SortRows() {
	switch pc.SortBy {
	case "ID":
		pc.sortByID()
	case "Name":
		pc.sortByName()
	case "Value":
		pc.sortByValue()
	}
}

func (pc *PerformanceComponent) sortByID() {
	for i := 1; i < len(pc.Rows); i++ {
		for j := i; j > 0 && pc.Rows[j-1].ID > pc.Rows[j].ID; j-- {
			pc.Rows[j], pc.Rows[j-1] = pc.Rows[j-1], pc.Rows[j]
		}
	}
}

func (pc *PerformanceComponent) sortByName() {
	for i := 1; i < len(pc.Rows); i++ {
		for j := i; j > 0 && pc.Rows[j-1].Name > pc.Rows[j].Name; j-- {
			pc.Rows[j], pc.Rows[j-1] = pc.Rows[j-1], pc.Rows[j]
		}
	}
}

func (pc *PerformanceComponent) sortByValue() {
	for i := 1; i < len(pc.Rows); i++ {
		for j := i; j > 0 && pc.Rows[j-1].Value > pc.Rows[j].Value; j-- {
			pc.Rows[j], pc.Rows[j-1] = pc.Rows[j-1], pc.Rows[j]
		}
	}
}

func (pc *PerformanceComponent) RenderTable(rowCount int) {
	start := time.Now()

	pc.Rows = make([]TableRow, rowCount)
	for i := 0; i < rowCount; i++ {
		pc.Rows[i] = TableRow{
			ID:    i,
			Name:  fmt.Sprintf("Row %d", i),
			Value: rand.Intn(1000),
		}
	}

	pc.TableRef = dom.Table()
	pc.TableRef.AddClass("table-auto", "w-full", "border-collapse")

	header := dom.TR()
	headers := []string{"ID", "Name", "Value"}

	for _, h := range headers {
		th := dom.TH()
		th.SetInnerHTML(h)
		th.OnClick(pc.onHeaderClick(h, header))
		header.Child(th)
	}

	pc.TableRef.Child(header)
	pc.TableRef.Child(pc.RenderRows())

	content := dom.ElementByID("content-table")
	content.SetInnerHTML("")
	content.Child(pc.TableRef)

	elapsed := time.Since(start)
	dom.ConsoleLog(fmt.Sprintf("Render time: %s", elapsed))
	dom.Alert(fmt.Sprintf("Render time: %s", elapsed))
}

func (pc *PerformanceComponent) onHeaderClick(h string, header dom.HTMLNode) dom.Func {
	return func(e dom.Event) {
		start := time.Now()
		pc.SortBy = h
		pc.SortRows()
		pc.TableRef.SetInnerHTML("")
		pc.TableRef.Child(header)
		pc.TableRef.Child(pc.RenderRows())
		elapsed := time.Since(start)
		dom.ConsoleLog(fmt.Sprintf("Render time: %s", elapsed))
		dom.Alert(fmt.Sprintf("Render time: %s", elapsed))
	}
}

func (pc *PerformanceComponent) Render() dom.HTMLNode {
	rowCountInput := dom.Input().SetType("number").SetPlaceholder("Enter number of rows").SetValue("1000").
		Tailwind(tlw.P2, tlw.Border, tlw.BorderGray300, tlw.RoundedMd, tlw.Mr2)
	renderButton := dom.Button("Render Table").
		Tailwind(tlw.P2, tlw.BgBlue500, tlw.TextWhite, tlw.RoundedMd, tlw.HoverBgBlue700)

	container := dom.Div().Tailwind(tlw.MxAuto, tlw.P4, tlw.ShadowMd, tlw.RoundedLg).
		Child(
			dom.H2("Performance Test").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2).
				Child(rowCountInput, renderButton),
			dom.Div().SetID("content-table").Tailwind(tlw.Mt4),
		)

	renderButton.OnClick(func(_ dom.Event) {
		rowCount, err := strconv.Atoi(rowCountInput.GetValue().String())
		if err != nil {
			js.Global().Call("alert", "Invalid row count")
			return
		}

		pc.RenderTable(rowCount)
	})

	return container
}
