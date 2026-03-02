package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"todo-list/internal/store"
	"todo-list/internal/todo"
)

var (
	s    = store.New("")
	list todo.List
)

func main() {
	list, _ = s.Load()

	a := app.New()
	w := a.NewWindow("Todo List")
	w.Resize(fyne.NewSize(600, 500))

	items := container.NewVBox()
	entry := widget.NewEntry()
	entry.SetPlaceHolder("New task...")

	var refresh func()

	refresh = func() {
		items.Objects = nil
		for _, t := range list {
			t := t
			check := widget.NewCheck(fmt.Sprintf("#%d  %s", t.ID, t.Description), nil)
			check.Checked = t.Status == todo.StatusCompleted
			check.OnChanged = func(checked bool) {
				if checked {
					list.Complete(t.ID)
					_ = s.Save(list)
					refresh()
				}
			}
			deleteBtn := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
				list.Delete(t.ID)
				_ = s.Save(list)
				refresh()
			})
			row := container.NewBorder(nil, nil, nil, deleteBtn, check)
			items.Add(row)
		}
		items.Refresh()
	}

	addBtn := widget.NewButton("Add", func() {
		desc := strings.TrimSpace(entry.Text)
		if desc == "" {
			return
		}
		list.Add(desc)
		_ = s.Save(list)
		entry.SetText("")
		refresh()
	})
	entry.OnSubmitted = func(_ string) { addBtn.OnTapped() }

	refresh()

	w.SetContent(container.NewBorder(
		nil,
		container.NewBorder(nil, nil, nil, addBtn, entry),
		nil, nil,
		container.NewVScroll(items),
	))
	w.ShowAndRun()
}
