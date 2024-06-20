/******************************************************************************/
/* css_property.go                                                            */
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
/* IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY       */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT  */
/* OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE      */
/* OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                              */
/******************************************************************************/

package properties

import (
	"kaiju/engine"
	"kaiju/markup/css/rules"
	"kaiju/markup/document"
	"kaiju/ui"
)

type Property interface {
	Key() string
	Process(panel *ui.Panel, elm *document.Element, values []rules.PropertyValue, host *engine.Host) error
}

var PropertyMap = map[string]Property{
	"accent-color":                AccentColor{},
	"align-content":               AlignContent{},
	"align-items":                 AlignItems{},
	"align-self":                  AlignSelf{},
	"all":                         All{},
	"animation":                   Animation{},
	"animation-delay":             AnimationDelay{},
	"animation-direction":         AnimationDirection{},
	"animation-duration":          AnimationDuration{},
	"animation-fill-mode":         AnimationFillMode{},
	"animation-iteration-count":   AnimationIterationCount{},
	"animation-name":              AnimationName{},
	"animation-play-state":        AnimationPlayState{},
	"animation-timing-function":   AnimationTimingFunction{},
	"aspect-ratio":                AspectRatio{},
	"backdrop-filter":             BackdropFilter{},
	"backface-visibility":         BackfaceVisibility{},
	"background":                  Background{},
	"background-attachment":       BackgroundAttachment{},
	"background-blend-mode":       BackgroundBlendMode{},
	"background-clip":             BackgroundClip{},
	"background-color":            BackgroundColor{},
	"background-image":            BackgroundImage{},
	"background-origin":           BackgroundOrigin{},
	"background-position":         BackgroundPosition{},
	"background-position-x":       BackgroundPositionX{},
	"background-position-y":       BackgroundPositionY{},
	"background-repeat":           BackgroundRepeat{},
	"background-size":             BackgroundSize{},
	"block-size":                  BlockSize{},
	"border":                      Border{},
	"border-block":                BorderBlock{},
	"border-block-color":          BorderBlockColor{},
	"border-block-end-color":      BorderBlockEndColor{},
	"border-block-end-style":      BorderBlockEndStyle{},
	"border-block-end-width":      BorderBlockEndWidth{},
	"border-block-start-color":    BorderBlockStartColor{},
	"border-block-start-style":    BorderBlockStartStyle{},
	"border-block-start-width":    BorderBlockStartWidth{},
	"border-block-style":          BorderBlockStyle{},
	"border-block-width":          BorderBlockWidth{},
	"border-bottom":               BorderBottom{},
	"border-bottom-color":         BorderBottomColor{},
	"border-bottom-left-radius":   BorderBottomLeftRadius{},
	"border-bottom-right-radius":  BorderBottomRightRadius{},
	"border-bottom-style":         BorderBottomStyle{},
	"border-bottom-width":         BorderBottomWidth{},
	"border-collapse":             BorderCollapse{},
	"border-color":                BorderColor{},
	"border-image":                BorderImage{},
	"border-image-outset":         BorderImageOutset{},
	"border-image-repeat":         BorderImageRepeat{},
	"border-image-slice":          BorderImageSlice{},
	"border-image-source":         BorderImageSource{},
	"border-image-width":          BorderImageWidth{},
	"border-inline":               BorderInline{},
	"border-inline-color":         BorderInlineColor{},
	"border-inline-end-color":     BorderInlineEndColor{},
	"border-inline-end-style":     BorderInlineEndStyle{},
	"border-inline-end-width":     BorderInlineEndWidth{},
	"border-inline-start-color":   BorderInlineStartColor{},
	"border-inline-start-style":   BorderInlineStartStyle{},
	"border-inline-start-width":   BorderInlineStartWidth{},
	"border-inline-style":         BorderInlineStyle{},
	"border-inline-width":         BorderInlineWidth{},
	"border-left":                 BorderLeft{},
	"border-left-color":           BorderLeftColor{},
	"border-left-style":           BorderLeftStyle{},
	"border-left-width":           BorderLeftWidth{},
	"border-radius":               BorderRadius{},
	"border-right":                BorderRight{},
	"border-right-color":          BorderRightColor{},
	"border-right-style":          BorderRightStyle{},
	"border-right-width":          BorderRightWidth{},
	"border-spacing":              BorderSpacing{},
	"border-style":                BorderStyle{},
	"border-top":                  BorderTop{},
	"border-top-color":            BorderTopColor{},
	"border-top-left-radius":      BorderTopLeftRadius{},
	"border-top-right-radius":     BorderTopRightRadius{},
	"border-top-style":            BorderTopStyle{},
	"border-top-width":            BorderTopWidth{},
	"border-width":                BorderWidth{},
	"bottom":                      Bottom{},
	"box-decoration-break":        BoxDecorationBreak{},
	"box-reflect":                 BoxReflect{},
	"box-shadow":                  BoxShadow{},
	"box-sizing":                  BoxSizing{},
	"break-after":                 BreakAfter{},
	"break-before":                BreakBefore{},
	"break-inside":                BreakInside{},
	"caption-side":                CaptionSide{},
	"caret-color":                 CaretColor{},
	"charset":                     Charset{},
	"clear":                       Clear{},
	"clip":                        Clip{},
	"color":                       Color{},
	"column-count":                ColumnCount{},
	"column-fill":                 ColumnFill{},
	"column-gap":                  ColumnGap{},
	"column-rule":                 ColumnRule{},
	"column-rule-color":           ColumnRuleColor{},
	"column-rule-style":           ColumnRuleStyle{},
	"column-rule-width":           ColumnRuleWidth{},
	"column-span":                 ColumnSpan{},
	"column-width":                ColumnWidth{},
	"columns":                     Columns{},
	"content":                     Content{},
	"counter-increment":           CounterIncrement{},
	"counter-reset":               CounterReset{},
	"cursor":                      Cursor{},
	"direction":                   Direction{},
	"display":                     Display{},
	"empty-cells":                 EmptyCells{},
	"filter":                      Filter{},
	"flex":                        Flex{},
	"flex-basis":                  FlexBasis{},
	"flex-direction":              FlexDirection{},
	"flex-flow":                   FlexFlow{},
	"flex-grow":                   FlexGrow{},
	"flex-shrink":                 FlexShrink{},
	"flex-wrap":                   FlexWrap{},
	"float":                       Float{},
	"font":                        Font{},
	"font-face":                   FontFace{},
	"font-family":                 FontFamily{},
	"font-feature-settings":       FontFeatureSettings{},
	"font-feature-values":         FontFeatureValues{},
	"font-kerning":                FontKerning{},
	"font-language-override":      FontLanguageOverride{},
	"font-size":                   FontSize{},
	"font-size-adjust":            FontSizeAdjust{},
	"font-stretch":                FontStretch{},
	"font-style":                  FontStyle{},
	"font-synthesis":              FontSynthesis{},
	"font-variant":                FontVariant{},
	"font-variant-alternates":     FontVariantAlternates{},
	"font-variant-caps":           FontVariantCaps{},
	"font-variant-east-asian":     FontVariantEastAsian{},
	"font-variant-ligatures":      FontVariantLigatures{},
	"font-variant-numeric":        FontVariantNumeric{},
	"font-variant-position":       FontVariantPosition{},
	"font-weight":                 FontWeight{},
	"gap":                         Gap{},
	"grid":                        Grid{},
	"grid-area":                   GridArea{},
	"grid-auto-columns":           GridAutoColumns{},
	"grid-auto-flow":              GridAutoFlow{},
	"grid-auto-rows":              GridAutoRows{},
	"grid-column":                 GridColumn{},
	"grid-column-end":             GridColumnEnd{},
	"grid-column-gap":             GridColumnGap{},
	"grid-column-start":           GridColumnStart{},
	"grid-gap":                    GridGap{},
	"grid-row":                    GridRow{},
	"grid-row-end":                GridRowEnd{},
	"grid-row-gap":                GridRowGap{},
	"grid-row-start":              GridRowStart{},
	"grid-template":               GridTemplate{},
	"grid-template-areas":         GridTemplateAreas{},
	"grid-template-columns":       GridTemplateColumns{},
	"grid-template-rows":          GridTemplateRows{},
	"hanging-punctuation":         HangingPunctuation{},
	"height":                      Height{},
	"hyphens":                     Hyphens{},
	"image-rendering":             ImageRendering{},
	"import":                      Import{},
	"inline-size":                 InlineSize{},
	"inset":                       Inset{},
	"inset-block":                 InsetBlock{},
	"inset-block-end":             InsetBlockEnd{},
	"inset-block-start":           InsetBlockStart{},
	"inset-inline":                InsetInline{},
	"inset-inline-end":            InsetInlineEnd{},
	"inset-inline-start":          InsetInlineStart{},
	"isolation":                   Isolation{},
	"justify-content":             JustifyContent{},
	"justify-items":               JustifyItems{},
	"justify-self":                JustifySelf{},
	"keyframes":                   Keyframes{},
	"left":                        Left{},
	"letter-spacing":              LetterSpacing{},
	"line-break":                  LineBreak{},
	"line-height":                 LineHeight{},
	"list-style":                  ListStyle{},
	"list-style-image":            ListStyleImage{},
	"list-style-position":         ListStylePosition{},
	"list-style-type":             ListStyleType{},
	"margin":                      Margin{},
	"margin-block":                MarginBlock{},
	"margin-block-end":            MarginBlockEnd{},
	"margin-block-start":          MarginBlockStart{},
	"margin-bottom":               MarginBottom{},
	"margin-inline":               MarginInline{},
	"margin-inline-end":           MarginInlineEnd{},
	"margin-inline-start":         MarginInlineStart{},
	"margin-left":                 MarginLeft{},
	"margin-right":                MarginRight{},
	"margin-top":                  MarginTop{},
	"mask":                        Mask{},
	"mask-clip":                   MaskClip{},
	"mask-composite":              MaskComposite{},
	"mask-image":                  MaskImage{},
	"mask-mode":                   MaskMode{},
	"mask-origin":                 MaskOrigin{},
	"mask-position":               MaskPosition{},
	"mask-repeat":                 MaskRepeat{},
	"mask-size":                   MaskSize{},
	"mask-type":                   MaskType{},
	"max-height":                  MaxHeight{},
	"max-width":                   MaxWidth{},
	"media":                       Media{},
	"max-block-size":              MaxBlockSize{},
	"max-inline-size":             MaxInlineSize{},
	"min-block-size":              MinBlockSize{},
	"min-inline-size":             MinInlineSize{},
	"min-height":                  MinHeight{},
	"min-width":                   MinWidth{},
	"mix-blend-mode":              MixBlendMode{},
	"object-fit":                  ObjectFit{},
	"object-position":             ObjectPosition{},
	"offset":                      Offset{},
	"offset-anchor":               OffsetAnchor{},
	"offset-distance":             OffsetDistance{},
	"offset-path":                 OffsetPath{},
	"offset-rotate":               OffsetRotate{},
	"opacity":                     Opacity{},
	"order":                       Order{},
	"orphans":                     Orphans{},
	"outline":                     Outline{},
	"outline-color":               OutlineColor{},
	"outline-offset":              OutlineOffset{},
	"outline-style":               OutlineStyle{},
	"outline-width":               OutlineWidth{},
	"overflow":                    Overflow{},
	"overflow-anchor":             OverflowAnchor{},
	"overflow-wrap":               OverflowWrap{},
	"overflow-x":                  OverflowX{},
	"overflow-y":                  OverflowY{},
	"overscroll-behavior":         OverscrollBehavior{},
	"overscroll-behavior-block":   OverscrollBehaviorBlock{},
	"overscroll-behavior-inline":  OverscrollBehaviorInline{},
	"overscroll-behavior-x":       OverscrollBehaviorX{},
	"overscroll-behavior-y":       OverscrollBehaviorY{},
	"padding":                     Padding{},
	"padding-block":               PaddingBlock{},
	"padding-block-end":           PaddingBlockEnd{},
	"padding-block-start":         PaddingBlockStart{},
	"padding-bottom":              PaddingBottom{},
	"padding-inline":              PaddingInline{},
	"padding-inline-end":          PaddingInlineEnd{},
	"padding-inline-start":        PaddingInlineStart{},
	"padding-left":                PaddingLeft{},
	"padding-right":               PaddingRight{},
	"padding-top":                 PaddingTop{},
	"page-break-after":            PageBreakAfter{},
	"page-break-before":           PageBreakBefore{},
	"page-break-inside":           PageBreakInside{},
	"paint-order":                 PaintOrder{},
	"perspective":                 Perspective{},
	"perspective-origin":          PerspectiveOrigin{},
	"place-content":               PlaceContent{},
	"place-items":                 PlaceItems{},
	"place-self":                  PlaceSelf{},
	"pointer-events":              PointerEvents{},
	"position":                    Position{},
	"quotes":                      Quotes{},
	"resize":                      Resize{},
	"right":                       Right{},
	"rotate":                      Rotate{},
	"row-gap":                     RowGap{},
	"scale":                       Scale{},
	"scroll-behavior":             ScrollBehavior{},
	"scroll-margin":               ScrollMargin{},
	"scroll-margin-block":         ScrollMarginBlock{},
	"scroll-margin-block-end":     ScrollMarginBlockEnd{},
	"scroll-margin-block-start":   ScrollMarginBlockStart{},
	"scroll-margin-bottom":        ScrollMarginBottom{},
	"scroll-margin-inline":        ScrollMarginInline{},
	"scroll-margin-inline-end":    ScrollMarginInlineEnd{},
	"scroll-margin-inline-start":  ScrollMarginInlineStart{},
	"scroll-margin-left":          ScrollMarginLeft{},
	"scroll-margin-right":         ScrollMarginRight{},
	"scroll-margin-top":           ScrollMarginTop{},
	"scroll-padding":              ScrollPadding{},
	"scroll-padding-block":        ScrollPaddingBlock{},
	"scroll-padding-block-end":    ScrollPaddingBlockEnd{},
	"scroll-padding-block-start":  ScrollPaddingBlockStart{},
	"scroll-padding-bottom":       ScrollPaddingBottom{},
	"scroll-padding-inline":       ScrollPaddingInline{},
	"scroll-padding-inline-end":   ScrollPaddingInlineEnd{},
	"scroll-padding-inline-start": ScrollPaddingInlineStart{},
	"scroll-padding-left":         ScrollPaddingLeft{},
	"scroll-padding-right":        ScrollPaddingRight{},
	"scroll-padding-top":          ScrollPaddingTop{},
	"scroll-snap-align":           ScrollSnapAlign{},
	"scroll-snap-stop":            ScrollSnapStop{},
	"scroll-snap-type":            ScrollSnapType{},
	"scrollbar-color":             ScrollbarColor{},
	"tab-size":                    TabSize{},
	"table-layout":                TableLayout{},
	"text-align":                  TextAlign{},
	"text-align-last":             TextAlignLast{},
	"text-combine-upright":        TextCombineUpright{},
	"text-decoration":             TextDecoration{},
	"text-decoration-color":       TextDecorationColor{},
	"text-decoration-line":        TextDecorationLine{},
	"text-decoration-style":       TextDecorationStyle{},
	"text-decoration-thickness":   TextDecorationThickness{},
	"text-emphasis":               TextEmphasis{},
	"text-indent":                 TextIndent{},
	"text-justify":                TextJustify{},
	"text-orientation":            TextOrientation{},
	"text-overflow":               TextOverflow{},
	"text-shadow":                 TextShadow{},
	"text-transform":              TextTransform{},
	"text-underline-position":     TextUnderlinePosition{},
	"top":                         Top{},
	"transform":                   Transform{},
	"transform-origin":            TransformOrigin{},
	"transform-style":             TransformStyle{},
	"transition":                  Transition{},
	"transition-delay":            TransitionDelay{},
	"transition-duration":         TransitionDuration{},
	"transition-property":         TransitionProperty{},
	"transition-timing-function":  TransitionTimingFunction{},
	"translate":                   Translate{},
	"unicode-bidi":                UnicodeBidi{},
	"user-select":                 UserSelect{},
	"vertical-align":              VerticalAlign{},
	"visibility":                  Visibility{},
	"white-space":                 WhiteSpace{},
	"widows":                      Widows{},
	"width":                       Width{},
	"word-break":                  WordBreak{},
	"word-spacing":                WordSpacing{},
	"word-wrap":                   WordWrap{},
	"writing-mode":                WritingMode{},
	"z-index":                     ZIndex{},
}
