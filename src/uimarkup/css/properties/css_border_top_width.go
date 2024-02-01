package properties

import (
	"errors"
	"kaiju/engine"
	"kaiju/ui"
	"kaiju/uimarkup/css/rules"
	"kaiju/uimarkup/markup"
)

// medium|thin|thick|length|initial|inherit
func (p BorderTopWidth) Process(panel *ui.Panel, elm markup.DocElement, values []rules.PropertyValue, host *engine.Host) error {
	if len(values) != 1 {
		return errors.New("BorderTopWidth requires exactly 1 value")
	} else {
		current := panel.Layout().Border()
		size := borderSizeFromStr(values[0].Str, host.Window, current.Y())
		panel.SetBorderSize(current.X(), size, current.Z(), current.W())
		return nil
	}
}
