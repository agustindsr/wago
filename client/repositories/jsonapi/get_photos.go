package jsonapi

import (
	"wasm/client/wa/dom"
	"wasm/client/wa/http"
)

type Photos []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

func GetPhotosTable() chan dom.TableElement {
	photosTableCh := make(chan dom.TableElement)

	// Ejecutar la solicitud HTTP en una goroutine
	go func() {
		photos, err := http.FetchData[Photos]("https://jsonplaceholder.typicode.com/photos", nil, nil)
		if err != nil {
			dom.Alert("Error fetching data: " + err.Error())
			photosTableCh <- dom.Table() // Enviar una tabla vacía en caso de error
			return
		}

		photosTableCh <- renderTable(photos)
	}()

	return photosTableCh
}

func renderTable(photos Photos) dom.TableElement {
	table := dom.Table()

	// Clear existing table rows
	table.ClearChildren()

	for _, photo := range photos {
		row := dom.TR()
		row.Child(dom.TD().SetInnerHTML(photo.ID))
		row.Child(dom.TD().SetInnerHTML(photo.Title))
		row.Child(dom.TD().SetInnerHTML(photo.URL))
		row.Child(dom.TD().SetInnerHTML(photo.ThumbnailURL))

		table.Child(row.HTMLNode)
	}

	return table
}
