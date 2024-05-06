package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func createFrame(text string) *tview.Frame {
	frame := tview.NewFrame()
	frame.SetTitleAlign(tview.AlignCenter)

	return frame
}

func CreateGui() {

	app := tview.NewApplication()

	mainLayout := tview.NewGrid().SetColumns(0, 0).SetRows(0, 0).SetBorders(true)
	mainLayout.SetBackgroundColor(tcell.ColorBlack.TrueColor())

	queryPickerFrame := createFrame("Picker")
	mainLayout.AddItem(queryPickerFrame, 0, 0, 2, 1, 0, 0, false)

	previewFrame := createFrame("Preview")
	mainLayout.AddItem(previewFrame, 0, 1, 4, 3, 0, 0, false)

	historyFrame := createFrame("History")
	mainLayout.AddItem(historyFrame, 2, 0, 2, 1, 0, 0, false)

	if err := app.SetRoot(mainLayout, true).Run(); err != nil {
		panic(err)
	}
}
