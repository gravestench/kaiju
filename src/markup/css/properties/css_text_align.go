/*****************************************************************************/
/* css_text_align.go                                                         */
/*****************************************************************************/
/*                           This file is part of:                           */
/*                                KAIJU ENGINE                               */
/*                          https://kaijuengine.org                          */
/*****************************************************************************/
/* MIT License                                                               */
/*                                                                           */
/* Copyright (c) 2022-present Kaiju Engine contributors (CONTRIBUTORS.md).   */
/* Copyright (c) 2015-2022 Brent Farris.                                     */
/*                                                                           */
/* May all those that this source may reach be blessed by the LORD and find  */
/* peace and joy in life.                                                    */
/* "Everyone who drinks of this water will be thirsty again; but whoever     */
/* drinks of the water that I will give him shall never thirst;" -Jesus      */
/*                                                                           */
/* Permission is hereby granted, free of charge, to any person obtaining a   */
/* copy of this software and associated documentation files (the "Software"),*/
/* to deal in the Software without restriction, including without limitation */
/* the rights to use, copy, modify, merge, publish, distribute, sublicense,  */
/* and/or sell copies of the Software, and to permit persons to whom the     */
/* Software is furnished to do so, subject to the following conditions:      */
/*                                                                           */
/* The above copyright, blessing, notice and this permission notice shall    */
/* be included in all copies or substantial portions of the Software.        */
/*                                                                           */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS   */
/* OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF                */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.    */
/* IN NO EVENT SHALL THE /* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT */
/* OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE     */
/* OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                             */
/*****************************************************************************/

package properties

import (
	"fmt"
	"kaiju/engine"
	"kaiju/markup/css/rules"
	"kaiju/markup/document"
	"kaiju/rendering"
	"kaiju/ui"
)

func childLabels(elm document.DocElement) []*ui.Label {
	labels := make([]*ui.Label, 0)
	for _, c := range elm.HTML.Children {
		if c.DocumentElement.HTML.IsText() {
			labels = append(labels, c.DocumentElement.UI.(*ui.Label))
		} else {
			labels = append(labels, childLabels(*c.DocumentElement)...)
		}
	}
	return labels
}

// left|right|center|justify|initial|inherit
func (p TextAlign) Process(panel *ui.Panel, elm document.DocElement, values []rules.PropertyValue, host *engine.Host) error {
	if len(values) != 1 {
		return fmt.Errorf("expected exactly 1 value but got %d", len(values))
	}
	labels := childLabels(elm)
	switch values[0].Str {
	case "left":
		for _, l := range labels {
			l.Layout().AnchorTo(l.Layout().Anchor().ConvertToLeft())
			l.SetJustify(rendering.FontJustifyLeft)
		}
	case "right":
		for _, l := range labels {
			l.Layout().AnchorTo(l.Layout().Anchor().ConvertToRight())
			l.SetJustify(rendering.FontJustifyRight)
		}
	case "center":
		for _, l := range labels {
			l.Layout().AnchorTo(l.Layout().Anchor().ConvertToCenter())
			l.SetJustify(rendering.FontJustifyCenter)
		}
	case "justify":
		// TODO:  Support text justification
	case "initial":
	case "inherit":
	}
	return nil
}
