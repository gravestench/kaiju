/*****************************************************************************/
/* x11.c                                                                     */
/*****************************************************************************/
/*                           This file is part of:                           */
/*                                KAIJU ENGINE                               */
/*                          https://kaijuengine.org                          */
/*****************************************************************************/
/* MIT License                                                               */
/*                                                                           */
/* Copyright (c) 2023-present Kaiju Engine contributors (CONTRIBUTORS.md).   */
/* Copyright (c) 2015-2023 Brent Farris.                                     */
/*                                                                           */
/* May all those that this source may reach be blessed by the LORD and find  */
/* peace and joy in life.                                                    */
/* Everyone who drinks of this water will be thirsty again; but whoever      */
/* drinks of the water that I will give him shall never thirst; John 4:13-14 */
/*                                                                           */
/* Permission is hereby granted, free of charge, to any person obtaining a   */
/* copy of this software and associated documentation files (the "Software"),*/
/* to deal in the Software without restriction, including without limitation */
/* the rights to use, copy, modify, merge, publish, distribute, sublicense,  */
/* and/or sell copies of the Software, and to permit persons to whom the     */
/* Software is furnished to do so, subject to the following conditions:      */
/*                                                                           */
/* The above copyright, blessing, biblical verse, notice and                 */
/* this permission notice shall be included in all copies or                 */
/* substantial portions of the Software.                                     */
/*                                                                           */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS   */
/* OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF                */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.    */
/* IN NO EVENT SHALL THE /* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT */
/* OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE     */
/* OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                             */
/*****************************************************************************/

#if defined(__linux__) || defined(__unix__) || defined(__APPLE__)

#include "x11.h"
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <stdbool.h>
#include <pthread.h>
#include <X11/Xlib.h>

#define EVT_MASK	ExposureMask | KeyPressMask | KeyReleaseMask | ButtonPressMask | ButtonReleaseMask | PointerMotionMask

static bool isExtensionSupported(const char* extList, const char* extension) {
	const char* start;
	const char* where, *terminator;
	where = strchr(extension, ' ');
	if (where || *extension == '\0') {
		return false;
	}
	for (start = extList;;) {
		where = strstr(start, extension);
		if (!where) {
			break;
		}
		terminator = where + strlen(extension);
		if (where == start || *(where - 1) == ' ') {
			if (*terminator == ' ' || *terminator == '\0') {
				return true;
			}
		}
		start = terminator;
	}
	return false;
}

void window_main(const char* windowTitle, int width, int height, void* evtSharedMem, int size) {
	Display* d = XOpenDisplay(NULL);
	if (d == NULL) {
		write_fatal(evtSharedMem, size, "Failed to open display");
		return;
	}
	Window w = XCreateSimpleWindow(d, RootWindow(d, DefaultScreen(d)), 10, 10,
		width, height, 1, BlackPixel(d, DefaultScreen(d)), WhitePixel(d, DefaultScreen(d)));
	if (w == None) {
		write_fatal(evtSharedMem, size, "Failed to create window");
		return;
	}
	XStoreName(d, w, windowTitle);
	XSetIconName(d, w, windowTitle);
	XSelectInput(d, w, EVT_MASK);
	XMapWindow(d, w);
	X11State* x11State = malloc(sizeof(X11State));
	x11State->sm.sharedMem = evtSharedMem;
	x11State->sm.size = size;
	x11State->w = w;
	x11State->d = d;
	x11State->WM_DELETE_WINDOW = XInternAtom(d, "WM_DELETE_WINDOW", False);
	XSetWMProtocols(d, w, &x11State->WM_DELETE_WINDOW, 1);
	memcpy(evtSharedMem+SHARED_MEM_DATA_START, &x11State, sizeof(x11State));
}

void window_show(void* x11State) {
	X11State* s = x11State;
	// Flush initial events
	while (window_poll(x11State) != Expose) {}
	while (window_poll(x11State) != 0) {}
}

int window_poll_controller(void* x11State) {
	// TODO:  Implement for controllers
	return 0;
}

int window_poll(void* x11State) {
	X11State* s = x11State;
	XEvent e = {};

	if (!XCheckMaskEvent(s->d, EVT_MASK, &e)) {
		return 0;
	}
	bool filtered = XFilterEvent(&e, None);
	s->sm.evt->evtType = e.type;
	switch (e.type) {
		case Expose:
			//return 0;
			break;
		case KeyPress:
		case KeyRelease:
			s->sm.evt->keyboard.keyId = XLookupKeysym(&e.xkey, 0);
			break;
		case ButtonPress:
		case ButtonRelease:
			switch (e.xbutton.button) {
				case Button1:
					s->sm.evt->mouse.mouseButtonId = MOUSE_BUTTON_LEFT;
					break;
				case Button2:
					s->sm.evt->mouse.mouseButtonId = MOUSE_BUTTON_MIDDLE;
					break;
				case Button3:
					s->sm.evt->mouse.mouseButtonId = MOUSE_BUTTON_RIGHT;
					break;
				case Button4:
					s->sm.evt->mouse.mouseButtonId = MOUSE_BUTTON_X1;
					break;
				case Button5:
					s->sm.evt->mouse.mouseButtonId = MOUSE_BUTTON_X2;
					break;
			}
			s->sm.evt->mouse.mouseX = e.xbutton.x;
			s->sm.evt->mouse.mouseY = e.xbutton.y;
			break;
		case MotionNotify:
			s->sm.evt->mouse.mouseButtonId = -1;
			s->sm.evt->mouse.mouseX = e.xmotion.x;
			s->sm.evt->mouse.mouseY = e.xmotion.y;
			break;
		case ClientMessage:
			if (!filtered) {
				const Atom protocol = e.xclient.data.l[0];
				if (protocol == s->WM_DELETE_WINDOW) {
					shared_memory_set_write_state(&s->sm, SHARED_MEM_QUIT);
				}
			}
			break;
	}
	return e.type;
}

void window_destroy(void* x11State) {
	X11State* s = x11State;
	XDestroyWindow(s->d, s->w);
	XCloseDisplay(s->d);
}

void* display(void* x11State) { return ((X11State*)x11State)->d; }
void* window(void* x11State) { return &((X11State*)x11State)->w; }

#endif
