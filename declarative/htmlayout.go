// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/Archs/go-htmlayout"
	"github.com/Archs/htmlayout"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type HtmLayout struct {
	AssignTo             **htmlayout.HtmLayout
	Name                 string
	Enabled              Property
	Visible              Property
	Font                 Font
	ToolTipText          Property
	MinSize              Size
	MaxSize              Size
	StretchFactor        int
	Row                  int
	RowSpan              int
	Column               int
	ColumnSpan           int
	AlwaysConsumeSpace   bool
	ContextMenuItems     []MenuItem
	OnKeyDown            walk.KeyEventHandler
	OnKeyPress           walk.KeyEventHandler
	OnKeyUp              walk.KeyEventHandler
	OnMouseDown          walk.MouseEventHandler
	OnMouseMove          walk.MouseEventHandler
	OnMouseUp            walk.MouseEventHandler
	OnSizeChanged        walk.EventHandler
	PageUrl              Property
	PageContent          Property
	OnPageUrlChanged     walk.EventHandler
	OnPageContentChanged walk.EventHandler
}

func (de HtmLayout) Create(builder *Builder) error {
	var w *htmlayout.HtmLayout
	var err error

	if de.PageContent != nil {
		w, err = htmlayout.NewHtmLayoutWithContent(builder.Parent(), de.PageContent.(string))
	} else {
		w, err = htmlayout.NewHtmLayout(builder.Parent(), de.PageUrl.(string))
	}
	if err != nil {
		return err
	}

	return builder.InitWidget(de, w, func() error {

		if de.OnPageUrlChanged != nil {
			w.PageUrlChanged().Attach(de.OnPageUrlChanged)
		}

		if de.OnPageContentChanged != nil {
			w.PageContentChanged().Attach(de.OnPageContentChanged)
		}

		if de.AssignTo != nil {
			*de.AssignTo = w
		}

		println("loading", "a.html", w.Handle(), w.Handle())
		if err := gohl.LoadFile(w.Handle(), "a.html"); err != nil {
			println("gohl.LoadFile failed:", err.Error())
		}

		return nil
	})
}

func (w HtmLayout) WidgetInfo() (name string, disabled, hidden bool, font *Font, toolTipText string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, alwaysConsumeSpace bool, contextMenuItems []MenuItem, OnKeyDown walk.KeyEventHandler, OnKeyPress walk.KeyEventHandler, OnKeyUp walk.KeyEventHandler, OnMouseDown walk.MouseEventHandler, OnMouseMove walk.MouseEventHandler, OnMouseUp walk.MouseEventHandler, OnSizeChanged walk.EventHandler) {
	return w.Name, false, false, &w.Font, "", w.MinSize, w.MaxSize, w.StretchFactor, w.Row, w.RowSpan, w.Column, w.ColumnSpan, w.AlwaysConsumeSpace, w.ContextMenuItems, w.OnKeyDown, w.OnKeyPress, w.OnKeyUp, w.OnMouseDown, w.OnMouseMove, w.OnMouseUp, w.OnSizeChanged
}
