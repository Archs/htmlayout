// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package htmlayout

import (
	"errors"
	"github.com/Archs/go-htmlayout"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

const (
	htmlayoutClassName = "walkHtmLayoutCLassName"
)

type HtmLayout struct {
	walk.WidgetBase
	pageUrlChangedPublisher     walk.EventPublisher
	pageContentChangedPublisher walk.EventPublisher
	format                      string
	// htmlPath to load
	pageUrl string
	// html as string
	pageContent string
}

func newHtmLayout(parent walk.Container) (*HtmLayout, error) {
	de := new(HtmLayout)

	if err := walk.InitWidget(
		de,
		parent,
		htmlayoutClassName,
		win.CS_HREDRAW|win.CS_VREDRAW,
		0); err != nil {
		return nil, err
	}

	de.MustRegisterProperty("PageUrl", walk.NewProperty(
		func() interface{} {
			return de.pageUrl
		},
		func(v interface{}) error {
			de.pageUrl = v.(string)
			return nil
		},
		de.pageUrlChangedPublisher.Event()))

	de.MustRegisterProperty("PageContent", walk.NewProperty(
		func() interface{} {
			return de.pageContent
		},
		func(v interface{}) error {
			de.pageContent = v.(string)
			return nil
		},
		de.pageContentChangedPublisher.Event()))
	return de, nil
}

func NewHtmLayout(parent walk.Container, url string) (*HtmLayout, error) {
	de, err := newHtmLayout(parent)
	if err != nil {
		return nil, err
	}
	de.pageUrl = url
	return de, nil
}

func NewHtmLayoutWithContent(parent walk.Container, html string) (*HtmLayout, error) {
	w, err := newHtmLayout(parent)
	if err != nil {
		return nil, err
	}
	w.pageContent = html
	return w, nil
}

func (*HtmLayout) LayoutFlags() walk.LayoutFlags {
	return walk.GrowableHorz
}

func (de *HtmLayout) MinSizeHint() walk.Size {
	return walk.Size{400, 300}
}

func (de *HtmLayout) SizeHint() walk.Size {
	return de.MinSizeHint()
}

func newError(msg string) error {
	return errors.New(msg)
}

func (de *HtmLayout) PageUrlChanged() *walk.Event {
	return de.pageUrlChangedPublisher.Event()
}

func (de *HtmLayout) PageContentChanged() *walk.Event {
	return de.pageContentChangedPublisher.Event()
}

func (de *HtmLayout) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	println(msg, win.WM_CREATE)
	// htmlayout message loop
	ret, handled := gohl.ProcNoDefault(hwnd, msg, wParam, lParam)
	println("procNoDefault:", handled, ret)
	if handled { //先把消息给HTMLayout
		println("")
		return uintptr(ret)
	}
	// begin default message loop
	// switch msg {
	// case win.WM_CREATE:
	// case 34:
	// println("loading", "a.html", hwnd, de.Handle())
	// if err := gohl.LoadFile(hwnd, "a.html"); err != nil {
	// 	println("gohl.LoadFile failed:", err.Error())
	// }
	// }
	return de.WindowBase.WndProc(hwnd, msg, wParam, lParam)
}

func init() {
	// var wc win.WNDCLASSEX
	// wc.CbSize = uint32(unsafe.Sizeof(wc))
	// wc.Style = win.CS_HREDRAW | win.CS_VREDRAW
	// wc.LpfnWndProc = syscall.NewCallback(wndProc)
	// wc.CbClsExtra = 0
	// wc.CbWndExtra = 0
	// wc.HInstance = win.GetModuleHandle(nil)
	// wc.HbrBackground = win.GetSysColorBrush(win.COLOR_WINDOWFRAME)
	// wc.LpszMenuName = syscall.StringToUTF16Ptr("")
	// wc.LpszClassName = syscall.StringToUTF16Ptr(htmlayoutClassName)
	// wc.HIconSm = win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_APPLICATION))
	// wc.HIcon = win.LoadIcon(0, win.MAKEINTRESOURCE(win.IDI_APPLICATION))
	// wc.HCursor = win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))

	// atom := win.RegisterClassEx(&wc)
	// if atom == 0 {
	// 	panic("Registering Class Failed:")
	// }
	walk.MustRegisterWindowClass(htmlayoutClassName)
}
