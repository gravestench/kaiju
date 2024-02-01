package ui

import (
	"kaiju/engine"
	"kaiju/hid"
	"kaiju/matrix"
	"kaiju/rendering"
	"kaiju/windowing"
)

type DirtyType = int

const (
	DirtyTypeNone = iota
	DirtyTypeLayout
	DirtyTypeResize
	DirtyTypeGenerated
	DirtyTypeReGenerated
	DirtyTypeColorChange
	DirtyTypeScissor
	DirtyTypeParent
	DirtyTypeParentLayout
	DirtyTypeParentResize
	DirtyTypeParentGenerated
	DirtyTypeParentReGenerated
	DirtyTypeParentColorChange
	DirtyTypeParentScissor
)

type UI interface {
	Entity() *engine.Entity
	ExecuteEvent(evtType EventType) bool
	AddEvent(evtType EventType, evt func()) engine.EventId
	RemoveEvent(evtType EventType, evtId engine.EventId)
	Event(evtType EventType) *engine.Event
	Update(deltaTime float64)
	SetDirty(dirtyType DirtyType)
	Layout() *Layout
	ShaderData() *ShaderData
	Clean()
	SetGroup(group *Group)
	Host() *engine.Host
	generateScissor()
	hasScissor() bool
	selfScissor() matrix.Vec4
	dirty() DirtyType
	setScissor(scissor matrix.Vec4)
	layoutChanged(dirtyType DirtyType)
	clean()
}

type uiBase struct {
	host                *engine.Host
	entity              *engine.Entity
	events              [EventTypeEnd]engine.Event
	group               *Group
	dragStartPos        matrix.Vec3
	downPos             matrix.Vec2
	layout              Layout
	dirtyType           DirtyType
	shaderData          ShaderData
	textureSize         matrix.Vec2
	updateId            int
	hovering            bool
	cantMiss            bool
	isDown              bool
	drag                bool
	lastActive          bool
	disconnectedScissor bool
}

func (ui *uiBase) init(host *engine.Host, textureSize matrix.Vec2, anchor Anchor, self UI) {
	ui.host = host
	ui.entity = host.NewEntity()
	ui.shaderData.ShaderDataBase = rendering.NewShaderDataBase()
	ui.shaderData.Scissor = matrix.Vec4{-matrix.FloatMax, -matrix.FloatMax, matrix.FloatMax, matrix.FloatMax}
	ui.entity.AddNamedData(EntityDataName, self)
	ui.textureSize = textureSize
	ui.layout.initialize(ui, anchor)
	if ui.updateId == 0 {
		ui.updateId = host.Updater.AddUpdate(ui.Update)
	}
}

func (ui *uiBase) Entity() *engine.Entity   { return ui.entity }
func (ui *uiBase) Layout() *Layout          { return &ui.layout }
func (ui *uiBase) hasScissor() bool         { return ui.shaderData.Scissor.X() > -matrix.FloatMax }
func (ui *uiBase) selfScissor() matrix.Vec4 { return ui.shaderData.Scissor }
func (ui *uiBase) Host() *engine.Host       { return ui.host }
func (ui *uiBase) dirty() DirtyType         { return ui.dirtyType }
func (ui *uiBase) ShaderData() *ShaderData  { return &ui.shaderData }
func (ui *uiBase) SetGroup(group *Group)    { ui.group = group }

func (ui *uiBase) ExecuteEvent(evtType EventType) bool {
	ui.events[evtType].Execute()
	return !ui.events[evtType].IsEmpty()
}

func (ui *uiBase) AddEvent(evtType EventType, evt func()) engine.EventId {
	return ui.events[evtType].Add(evt)
}

func (ui *uiBase) RemoveEvent(evtType EventType, evtId engine.EventId) {
	ui.events[evtType].Remove(evtId)
}

func (ui *uiBase) Event(evtType EventType) *engine.Event {
	return &ui.events[evtType]
}

func (ui *uiBase) SetDirty(dirtyType DirtyType) {
	if ui.dirtyType == DirtyTypeNone || ui.dirtyType >= DirtyTypeParent || dirtyType == DirtyTypeGenerated || dirtyType == DirtyTypeReGenerated {
		ui.dirtyType = dirtyType
		for i := 0; i < len(ui.entity.Children); i++ {
			kid := ui.entity.Children[i]
			all := AllOnEntity(kid)
			for _, cui := range all {
				if cui.dirty() == DirtyTypeNone || cui.dirty() > DirtyTypeParent {
					// TODO:  Let it know it was from the parent and what type
					if ui.dirtyType < DirtyTypeParent {
						cui.SetDirty(DirtyTypeParent + ui.dirtyType)
					} else {
						cui.SetDirty(ui.dirtyType)
					}
				}
			}
		}
	}
}

func (ui *uiBase) Clean() {
	if ui.dirtyType != DirtyTypeNone {
		if ui.entity.Parent != nil {
			cleanParent(ui.host, ui.entity.Parent)
		}
		ui.clean()
		for i := 0; i < len(ui.entity.Children); i++ {
			c := ui.entity.Children[i]
			all := AllOnEntity(c)
			for _, childUI := range all {
				if childUI.dirty() != DirtyTypeNone {
					childUI.Clean()
				}
			}
		}
	}
}

func (ui *uiBase) clean() {
	ui.layout.update()
	if !ui.events[EventTypeRebuild].IsEmpty() {
		ui.events[EventTypeRebuild].Execute()
	}
	ui.entity.Transform.SetDirty()
	if ui.dirtyType == DirtyTypeReGenerated {
		ui.dirtyType = DirtyTypeGenerated
	} else {
		ui.dirtyType = DirtyTypeNone
	}
	ui.shaderData.setSize2d(ui, ui.textureSize.X(), ui.textureSize.Y())
}

func cleanParent(host *engine.Host, entity *engine.Entity) {
	if entity.Parent != nil {
		cleanParent(host, entity.Parent)
	}
	all := AllOnEntity(entity)
	for i := 0; i < len(all); i++ {
		if all[i].dirty() != DirtyTypeNone {
			all[i].clean()
		}
	}
}

func (ui *uiBase) generateScissor() {
	if !ui.hasScissor() {
		pos := ui.entity.Transform.WorldPosition()
		size := ui.entity.Transform.WorldScale()
		bounds := matrix.Vec4{
			pos.X() - size.X()*0.5,
			pos.Y() - size.Y()*0.5,
			pos.X() + size.X()*0.5,
			pos.Y() + size.Y()*0.5,
		}
		ui.setScissor(bounds)
	}
}

func (ui *uiBase) setScissor(scissor matrix.Vec4) {
	if ui.disconnectedScissor {
		return
	}
	ui.shaderData.Scissor = scissor
	for i := 0; i < len(ui.entity.Children); i++ {
		cUI := FirstOnEntity(ui.entity.Children[i])
		if cUI != nil {
			cUI.setScissor(scissor)
		}
	}
	ui.shaderData.Scissor = scissor
	if ui.dirtyType == DirtyTypeNone {
		ui.SetDirty(DirtyTypeScissor)
	}
}

func (ui *uiBase) requestEvent(evtType EventType) {
	if ui.group != nil {
		ui.group.requestEvent(ui, evtType)
	} else {
		ui.ExecuteEvent(evtType)
	}
}

func (ui *uiBase) Update(deltaTime float64) {
	cursor := &ui.host.Window.Cursor
	if cursor.Moved() {
		pos := ui.cursorPos(cursor)
		ui.containedCheck(cursor, ui.entity)
		if ui.isDown && !ui.drag {
			w := ui.Host().Window.Width()
			h := ui.Host().Window.Height()
			wmm, hmm, _ := ui.host.Window.GetDPI()
			threshold := max(windowing.DPI2PX(w, wmm, 4), windowing.DPI2PX(h, hmm, 4))
			if ui.downPos.Distance(pos) > float32(threshold) {
				ui.dragStartPos = ui.entity.Transform.WorldPosition()
				ui.drag = true
			}
		}
	}
	if cursor.Pressed() {
		ui.containedCheck(cursor, ui.entity)
		if ui.hovering && !ui.isDown {
			ui.isDown = true
			ui.downPos = ui.cursorPos(cursor)
			ui.requestEvent(EventTypeDown)
			ui.cantMiss = true
		}
	}
	if cursor.Released() {
		if ui.hovering {
			ui.requestEvent(EventTypeUp)
		}
		if ui.lastActive {
			if ui.isDown {
				ui.isDown = false
				dragged := false
				if ui.drag {
					p := ui.entity.Transform.WorldPosition()
					dragged = ui.dragStartPos.Distance(p) > 5
				}
				ui.drag = false
				if ui.hovering && !dragged {
					ui.requestEvent(EventTypeClick)
				}
			} else if !ui.hovering && !ui.cantMiss {
				ui.requestEvent(EventTypeMiss)
			}
			ui.cantMiss = false
		}
	}
	if ui.host.Window.Mouse.Scrolled() && ui.hovering {
		ui.requestEvent(EventTypeScroll)
	}
	if ui.dirtyType != DirtyTypeNone {
		if ui.entity.Parent != nil {
			cleanParent(ui.Host(), ui.entity.Parent)
		}
		ui.clean()
	}
	ui.lastActive = ui.entity.IsActive()
}

func (ui *uiBase) cursorPos(cursor *hid.Cursor) matrix.Vec2 {
	//camPos := ui.host.UICamera.Position()
	//return cursor.ScreenPosition().Add(matrix.Vec2{camPos.X(), camPos.Y()})
	return cursor.Position()
}

func (ui *uiBase) containedCheck(cursor *hid.Cursor, entity *engine.Entity) {
	cp := ui.cursorPos(cursor)
	contained := entity.Transform.ContainsPoint2D(cp)
	if contained && ui.hasScissor() {
		contained = ui.shaderData.Scissor.ScreenAreaContains(cp.X(), cp.Y())
	}
	if !ui.hovering && contained {
		ui.hovering = true
		ui.requestEvent(EventTypeEnter)
	} else if ui.hovering && !contained {
		ui.hovering = false
		ui.requestEvent(EventTypeExit)
	}
}

func (ui *uiBase) changed() {
	ui.ExecuteEvent(EventTypeChange)
}

func (ui *uiBase) SetScissorToParent() {
	ui.disconnectedScissor = false
	if elm := FirstOnEntity(ui.entity.Parent); elm != nil {
		elm.setScissor(elm.selfScissor())
	} else {
		ui.generateScissor()
	}
}

func (ui *uiBase) DisconnectParentScissor() {
	if ui.hasScissor() {
		ui.setScissor(matrix.Vec4{-matrix.FloatMax, -matrix.FloatMax, matrix.FloatMax, matrix.FloatMax})
		ui.generateScissor()
	}
	ui.disconnectedScissor = true
}

func (ui uiBase) layoutChanged(dirtyType DirtyType) {
	ui.SetDirty(dirtyType)
	if ui.Entity().Parent != nil {
		if pui := FirstOnEntity(ui.Entity().Parent); pui != nil {
			if pui.dirty() == DirtyTypeNone {
				pui.SetDirty(DirtyTypeParentLayout)
			} else {
				pui.SetDirty(DirtyTypeReGenerated)
			}
		}
	}
}
