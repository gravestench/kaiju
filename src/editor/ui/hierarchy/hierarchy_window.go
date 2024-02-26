/******************************************************************************/
/* hierarchy.go                                                               */
/******************************************************************************/
/*                           This file is part of:                            */
/*                                KAIJU ENGINE                                */
/*                          https://kaijuengine.org                           */
/******************************************************************************/
/* MIT License                                                                */
/*                                                                            */
/* Copyright (c) 2023-present Kaiju Engine authors (AUTHORS.md).              */
/* Copyright (c) 2015-present Brent Farris.                                   */
/*                                                                            */
/* May all those that this source may reach be blessed by the LORD and find   */
/* peace and joy in life.                                                     */
/* Everyone who drinks of this water will be thirsty again; but whoever       */
/* drinks of the water that I will give him shall never thirst; John 4:13-14  */
/*                                                                            */
/* Permission is hereby granted, free of charge, to any person obtaining a    */
/* copy of this software and associated documentation files (the "Software"), */
/* to deal in the Software without restriction, including without limitation  */
/* the rights to use, copy, modify, merge, publish, distribute, sublicense,   */
/* and/or sell copies of the Software, and to permit persons to whom the      */
/* Software is furnished to do so, subject to the following conditions:       */
/*                                                                            */
/* The above copyright, blessing, biblical verse, notice and                  */
/* this permission notice shall be included in all copies or                  */
/* substantial portions of the Software.                                      */
/*                                                                            */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS    */
/* OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF                 */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.     */
/* IN NO EVENT SHALL THE /* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY    */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT  */
/* OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE      */
/* OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                              */
/******************************************************************************/

package hierarchy

import (
	"kaiju/editor/interfaces"
	"kaiju/engine"
	"kaiju/klib"
	"kaiju/markup"
	"kaiju/markup/document"
	"kaiju/matrix"
	"kaiju/systems/events"
	"kaiju/ui"
	"log/slog"
	"strings"
)

type Hierarchy struct {
	editor         interfaces.Editor
	doc            *document.Document
	input          *ui.Input
	onChangeId     events.Id
	selectChangeId events.Id
	query          string
	uiGroup        *ui.Group
}

type entityEntry struct {
	Entity          *engine.Entity
	ShowingChildren bool
}

type hierarchyData struct {
	Entries []entityEntry
	Query   string
}

func (e entityEntry) Depth() []struct{} {
	depth := make([]struct{}, 0, 10)
	p := e.Entity
	for p.Parent != nil {
		depth = append(depth, struct{}{})
		p = p.Parent
	}
	return depth
}

func New(editor interfaces.Editor, uiGroup *ui.Group) *Hierarchy {
	h := &Hierarchy{
		editor:  editor,
		uiGroup: uiGroup,
	}
	h.editor.Host().OnClose.Add(func() {
		if h.doc != nil {
			h.doc.Destroy()
		}
	})
	return h
}

func (h *Hierarchy) Toggle() {
	if h.doc == nil {
		h.Show()
	} else {
		if h.doc.Elements[0].UI.Entity().IsActive() {
			h.Hide()
		} else {
			h.Show()
		}
	}
}

func (h *Hierarchy) Show() {
	if h.doc == nil {
		h.Reload()
	} else {
		h.doc.Activate()
	}
}

func (h *Hierarchy) Hide() {
	if h.doc != nil {
		h.doc.Deactivate()
	}
}

func (h *Hierarchy) orderEntitiesVisually() []entityEntry {
	allEntities := h.editor.Host().Entities()
	entries := make([]entityEntry, 0, len(allEntities))
	roots := make([]*engine.Entity, 0, len(allEntities))
	for _, entity := range allEntities {
		if entity.IsRoot() {
			roots = append(roots, entity)
		}
	}
	var addChildren func(*engine.Entity)
	addChildren = func(entity *engine.Entity) {
		entries = append(entries, entityEntry{entity, false})
		for _, c := range entity.Children {
			addChildren(c)
		}
	}
	for _, r := range roots {
		addChildren(r)
	}
	return entries
}

func (h *Hierarchy) Init() {
	h.Reload()
}

func (h *Hierarchy) filter(entries []entityEntry) []entityEntry {
	if h.query == "" {
		return entries
	}
	filtered := make([]entityEntry, 0, len(entries))
	// TODO:  Append the entire path to the entity kf not already appended
	for _, e := range entries {
		if strings.Contains(strings.ToLower(e.Entity.Name()), h.query) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func (h *Hierarchy) Reload() {
	if h.doc != nil {
		h.doc.Destroy()
	}
	data := hierarchyData{
		Entries: h.filter(h.orderEntitiesVisually()),
		Query:   h.query,
	}
	host := h.editor.Host()
	host.CreatingEditorEntities()
	h.doc = klib.MustReturn(markup.DocumentFromHTMLAsset(
		host, "editor/ui/hierarchy_window.html", data,
		map[string]func(*document.DocElement){
			"selectedEntity": h.selectedEntity,
		}))
	h.doc.SetGroup(h.uiGroup)
	host.DoneCreatingEditorEntities()
	h.selectChangeId = h.editor.Selection().Changed.Add(h.onSelectionChanged)
	if elm, ok := h.doc.GetElementById("searchInput"); !ok {
		slog.Error(`Failed to locate the "searchInput" for the hierarchy`)
	} else {
		h.input = elm.UI.(*ui.Input)
		h.input.Data().OnSubmit.Add(h.submit)
	}
}

func (h *Hierarchy) submit() {
	h.query = strings.ToLower(strings.TrimSpace(h.input.Text()))
	h.Reload()
}

func (h *Hierarchy) onSelectionChanged() {
	elm, ok := h.doc.GetElementById("list")
	if !ok {
		slog.Error("Could not find hierarchy list, reopen the hierarchy window")
		return
	}
	for i := range elm.HTML.Children {
		elm.HTML.Children[i].DocumentElement.UIPanel.UnEnforceColor()
	}
	for i := range elm.HTML.Children {
		c := &elm.HTML.Children[i]
		id := c.Attribute("id")
		for _, se := range h.editor.Selection().Entities() {
			if se.Id() == id {
				c.DocumentElement.UIPanel.EnforceColor(matrix.ColorDarkBlue())
				break
			}
		}
	}
}

func (h *Hierarchy) selectedEntity(elm *document.DocElement) {
	id := elm.HTML.Attribute("id")
	if e, ok := h.editor.Host().FindEntity(id); !ok {
		slog.Error("Could not find entity", slog.String("id", id))
	} else {
		kb := &h.editor.Host().Window.Keyboard
		if kb.HasCtrl() {
			h.editor.Selection().Toggle(e)
		} else if kb.HasShift() {
			h.editor.Selection().Add(e)
		} else {
			h.editor.Selection().Set(e)
		}
		h.onSelectionChanged()
	}
}
