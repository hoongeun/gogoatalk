package ui

import "github.com/rivo/tview"

func NewLayout(sb *Sidebar, mc *messageContainer, mi *messageInput) *tview.Grid {
	main := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mc.view, 0, 1, false).
		AddItem(mi.view, 1, 1, true)

	grid := tview.NewGrid().
		SetRows(1, 0).
		SetBorders(true).
		AddItem(tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText("Gogoatalk v0.0.1"),
			0, 0, 1, 4, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(sb.view, 0, 0, 0, 1, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(sb.view, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 3, 0, 100, true)

	return grid
}
