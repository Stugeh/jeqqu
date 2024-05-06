package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createTextView(text string) *tview.TextView {
	view := tview.NewTextView().
		SetText(text).
		SetTextAlign(tview.AlignCenter).
		SetTextColor(tview.Styles.PrimaryTextColor)

	view.SetBackgroundColor(tcell.ColorBlack.TrueColor())

	return view
}

func main() {
	app := tview.NewApplication()

	mainLayout := tview.NewGrid().SetColumns(0, 0).SetRows(0, 0).SetBorders(true)
	mainLayout.SetBackgroundColor(tcell.ColorBlack.TrueColor())

	queryPicker := createTextView("Picker")
	mainLayout.AddItem(queryPicker, 0, 0, 1, 1, 0, 0, false)

	preview := createTextView("Preview")
	mainLayout.AddItem(preview, 0, 2, 1, 1, 0, 0, false)

	historyPanel := createTextView("History")
	mainLayout.AddItem(historyPanel, 2, 0, 1, 1, 0, 0, false)

	if err := app.SetRoot(mainLayout, true).Run(); err != nil {
		panic(err)
	}
}
