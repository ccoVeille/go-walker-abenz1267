// This file is autogenerated from: wlr-foreign-toplevel-management-unstable-v1.xml
// Do not edit

// Package wlr implements the wlr_foreign_toplevel_management_unstable_v1 protocol
package wlr

import (
	"sync"

	"github.com/neurlang/wayland/wl"
)

// ZwlrForeignToplevelHandleV1StateMaximized means the toplevel is maximized
const ZwlrForeignToplevelHandleV1StateMaximized = 0

// ZwlrForeignToplevelHandleV1StateMinimized means the toplevel is minimized
const ZwlrForeignToplevelHandleV1StateMinimized = 1

// ZwlrForeignToplevelHandleV1StateActivated means the toplevel is active
const ZwlrForeignToplevelHandleV1StateActivated = 2

// ZwlrForeignToplevelHandleV1StateFullscreen means the toplevel is fullscreen
const ZwlrForeignToplevelHandleV1StateFullscreen = 3

// ZwlrForeignToplevelHandleV1ErrorInvalidRectangle means the provided rectangle is invalid
const ZwlrForeignToplevelHandleV1ErrorInvalidRectangle = 0

// ZwlrForeignToplevelManagerV1 list and control opened apps
type ZwlrForeignToplevelManagerV1 struct {
	wl.BaseProxy
	mu                                           sync.RWMutex
	privateZwlrForeignToplevelManagerV1Toplevels []ZwlrForeignToplevelManagerV1ToplevelHandler
	privateZwlrForeignToplevelManagerV1Finisheds []ZwlrForeignToplevelManagerV1FinishedHandler
}

// NewZwlrForeignToplevelManagerV1 is a constructor for the ZwlrForeignToplevelManagerV1 object
func NewZwlrForeignToplevelManagerV1(ctx *wl.Context) *ZwlrForeignToplevelManagerV1 {
	ret := new(ZwlrForeignToplevelManagerV1)
	ctx.Register(ret)
	return ret
}

// Stop stop sending events
func (p *ZwlrForeignToplevelManagerV1) Stop() error {
	return p.Context().SendRequest(p, 0)
}

// Dispatch dispatches event for object ZwlrForeignToplevelManagerV1
func (p *ZwlrForeignToplevelManagerV1) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.privateZwlrForeignToplevelManagerV1Toplevels) > 0 {
			ev := ZwlrForeignToplevelManagerV1ToplevelEvent{}
			ev.Toplevel = event.NewId(new(ZwlrForeignToplevelHandleV1), p.Context()).(*ZwlrForeignToplevelHandleV1)
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelManagerV1Toplevels {
				h.HandleZwlrForeignToplevelManagerV1Toplevel(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.privateZwlrForeignToplevelManagerV1Finisheds) > 0 {
			ev := ZwlrForeignToplevelManagerV1FinishedEvent{}
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelManagerV1Finisheds {
				h.HandleZwlrForeignToplevelManagerV1Finished(ev)
			}
			p.mu.RUnlock()
		}

	}
}

// ZwlrForeignToplevelManagerV1ToplevelEvent is the a toplevel has been created
type ZwlrForeignToplevelManagerV1ToplevelEvent struct {
	// Toplevel is the
	Toplevel *ZwlrForeignToplevelHandleV1
}

// ZwlrForeignToplevelManagerV1FinishedEvent is the the compositor has finished with the toplevel manager
type ZwlrForeignToplevelManagerV1FinishedEvent struct{}

// ZwlrForeignToplevelManagerV1ToplevelHandler is the handler interface for ZwlrForeignToplevelManagerV1ToplevelEvent
type ZwlrForeignToplevelManagerV1ToplevelHandler interface {
	HandleZwlrForeignToplevelManagerV1Toplevel(ZwlrForeignToplevelManagerV1ToplevelEvent)
}

// AddToplevelHandler removes the Toplevel handler
func (p *ZwlrForeignToplevelManagerV1) AddToplevelHandler(h ZwlrForeignToplevelManagerV1ToplevelHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelManagerV1Toplevels = append(p.privateZwlrForeignToplevelManagerV1Toplevels, h)
		p.mu.Unlock()
	}
}

// RemoveToplevelHandler adds the Toplevel handler
func (p *ZwlrForeignToplevelManagerV1) RemoveToplevelHandler(h ZwlrForeignToplevelManagerV1ToplevelHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelManagerV1Toplevels {
		if e == h {
			p.privateZwlrForeignToplevelManagerV1Toplevels = append(p.privateZwlrForeignToplevelManagerV1Toplevels[:i], p.privateZwlrForeignToplevelManagerV1Toplevels[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelManagerV1FinishedHandler is the handler interface for ZwlrForeignToplevelManagerV1FinishedEvent
type ZwlrForeignToplevelManagerV1FinishedHandler interface {
	HandleZwlrForeignToplevelManagerV1Finished(ZwlrForeignToplevelManagerV1FinishedEvent)
}

// AddFinishedHandler removes the Finished handler
func (p *ZwlrForeignToplevelManagerV1) AddFinishedHandler(h ZwlrForeignToplevelManagerV1FinishedHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelManagerV1Finisheds = append(p.privateZwlrForeignToplevelManagerV1Finisheds, h)
		p.mu.Unlock()
	}
}

// RemoveFinishedHandler adds the Finished handler
func (p *ZwlrForeignToplevelManagerV1) RemoveFinishedHandler(h ZwlrForeignToplevelManagerV1FinishedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelManagerV1Finisheds {
		if e == h {
			p.privateZwlrForeignToplevelManagerV1Finisheds = append(p.privateZwlrForeignToplevelManagerV1Finisheds[:i], p.privateZwlrForeignToplevelManagerV1Finisheds[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1 an opened toplevel
type ZwlrForeignToplevelHandleV1 struct {
	wl.BaseProxy
	mu                                             sync.RWMutex
	privateZwlrForeignToplevelHandleV1Titles       []ZwlrForeignToplevelHandleV1TitleHandler
	privateZwlrForeignToplevelHandleV1AppIds       []ZwlrForeignToplevelHandleV1AppIdHandler
	privateZwlrForeignToplevelHandleV1OutputEnters []ZwlrForeignToplevelHandleV1OutputEnterHandler
	privateZwlrForeignToplevelHandleV1OutputLeaves []ZwlrForeignToplevelHandleV1OutputLeaveHandler
	privateZwlrForeignToplevelHandleV1States       []ZwlrForeignToplevelHandleV1StateHandler
	privateZwlrForeignToplevelHandleV1Dones        []ZwlrForeignToplevelHandleV1DoneHandler
	privateZwlrForeignToplevelHandleV1Closeds      []ZwlrForeignToplevelHandleV1ClosedHandler
	privateZwlrForeignToplevelHandleV1Parents      []ZwlrForeignToplevelHandleV1ParentHandler
}

// NewZwlrForeignToplevelHandleV1 is a constructor for the ZwlrForeignToplevelHandleV1 object
func NewZwlrForeignToplevelHandleV1(ctx *wl.Context) *ZwlrForeignToplevelHandleV1 {
	ret := new(ZwlrForeignToplevelHandleV1)
	ctx.Register(ret)
	return ret
}

// SetMaximized requests that the toplevel be maximized
func (p *ZwlrForeignToplevelHandleV1) SetMaximized() error {
	return p.Context().SendRequest(p, 0)
}

// UnsetMaximized requests that the toplevel be unmaximized
func (p *ZwlrForeignToplevelHandleV1) UnsetMaximized() error {
	return p.Context().SendRequest(p, 1)
}

// SetMinimized requests that the toplevel be minimized
func (p *ZwlrForeignToplevelHandleV1) SetMinimized() error {
	return p.Context().SendRequest(p, 2)
}

// UnsetMinimized requests that the toplevel be unminimized
func (p *ZwlrForeignToplevelHandleV1) UnsetMinimized() error {
	return p.Context().SendRequest(p, 3)
}

// Activate activate the toplevel
func (p *ZwlrForeignToplevelHandleV1) Activate(Seat *wl.Seat) error {
	return p.Context().SendRequest(p, 4, Seat)
}

// Close request that the toplevel be closed
func (p *ZwlrForeignToplevelHandleV1) Close() error {
	return p.Context().SendRequest(p, 5)
}

// SetRectangle the rectangle which represents the toplevel
func (p *ZwlrForeignToplevelHandleV1) SetRectangle(Surface *wl.Surface, X int32, Y int32, Width int32, Height int32) error {
	return p.Context().SendRequest(p, 6, Surface, X, Y, Width, Height)
}

// Destroy destroy the zwlr_foreign_toplevel_handle_v1 object
func (p *ZwlrForeignToplevelHandleV1) Destroy() error {
	return p.Context().SendRequest(p, 7)
}

// SetFullscreen request that the toplevel be fullscreened
func (p *ZwlrForeignToplevelHandleV1) SetFullscreen(Output *wl.Output) error {
	return p.Context().SendRequest(p, 8, Output)
}

// UnsetFullscreen request that the toplevel be unfullscreened
func (p *ZwlrForeignToplevelHandleV1) UnsetFullscreen() error {
	return p.Context().SendRequest(p, 9)
}

// Dispatch dispatches event for object ZwlrForeignToplevelHandleV1
func (p *ZwlrForeignToplevelHandleV1) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.privateZwlrForeignToplevelHandleV1Titles) > 0 {
			ev := ZwlrForeignToplevelHandleV1TitleEvent{}
			ev.Title = event.String()
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1Titles {
				h.HandleZwlrForeignToplevelHandleV1Title(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.privateZwlrForeignToplevelHandleV1AppIds) > 0 {
			ev := ZwlrForeignToplevelHandleV1AppIdEvent{}
			ev.AppId = event.String()
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1AppIds {
				h.HandleZwlrForeignToplevelHandleV1AppId(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.privateZwlrForeignToplevelHandleV1OutputEnters) > 0 {
			ev := ZwlrForeignToplevelHandleV1OutputEnterEvent{}
			ev.Output = wl.SafeCast[*wl.Output](event.Proxy(p.Context()))
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1OutputEnters {
				h.HandleZwlrForeignToplevelHandleV1OutputEnter(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.privateZwlrForeignToplevelHandleV1OutputLeaves) > 0 {
			ev := ZwlrForeignToplevelHandleV1OutputLeaveEvent{}
			ev.Output = wl.SafeCast[*wl.Output](event.Proxy(p.Context()))
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1OutputLeaves {
				h.HandleZwlrForeignToplevelHandleV1OutputLeave(ev)
			}
			p.mu.RUnlock()
		}
	case 4:
		if len(p.privateZwlrForeignToplevelHandleV1States) > 0 {
			ev := ZwlrForeignToplevelHandleV1StateEvent{}
			ev.State = event.Array()
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1States {
				h.HandleZwlrForeignToplevelHandleV1State(ev)
			}
			p.mu.RUnlock()
		}
	case 5:
		if len(p.privateZwlrForeignToplevelHandleV1Dones) > 0 {
			ev := ZwlrForeignToplevelHandleV1DoneEvent{}
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1Dones {
				h.HandleZwlrForeignToplevelHandleV1Done(ev)
			}
			p.mu.RUnlock()
		}
	case 6:
		if len(p.privateZwlrForeignToplevelHandleV1Closeds) > 0 {
			ev := ZwlrForeignToplevelHandleV1ClosedEvent{}
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1Closeds {
				h.HandleZwlrForeignToplevelHandleV1Closed(ev)
			}
			p.mu.RUnlock()
		}
	case 7:
		if len(p.privateZwlrForeignToplevelHandleV1Parents) > 0 {
			ev := ZwlrForeignToplevelHandleV1ParentEvent{}
			ev.Parent = wl.SafeCast[*ZwlrForeignToplevelHandleV1](event.Proxy(p.Context()))
			p.mu.RLock()
			for _, h := range p.privateZwlrForeignToplevelHandleV1Parents {
				h.HandleZwlrForeignToplevelHandleV1Parent(ev)
			}
			p.mu.RUnlock()
		}

	}
}

// ZwlrForeignToplevelHandleV1TitleEvent is the title change
type ZwlrForeignToplevelHandleV1TitleEvent struct {
	// Title is the
	Title string
}

// ZwlrForeignToplevelHandleV1AppIdEvent is the app-id change
type ZwlrForeignToplevelHandleV1AppIdEvent struct {
	// AppId is the
	AppId string
}

// ZwlrForeignToplevelHandleV1OutputEnterEvent is the toplevel entered an output
type ZwlrForeignToplevelHandleV1OutputEnterEvent struct {
	// Output is the
	Output *wl.Output
}

// ZwlrForeignToplevelHandleV1OutputLeaveEvent is the toplevel left an output
type ZwlrForeignToplevelHandleV1OutputLeaveEvent struct {
	// Output is the
	Output *wl.Output
}

// ZwlrForeignToplevelHandleV1StateEvent is the the toplevel state changed
type ZwlrForeignToplevelHandleV1StateEvent struct {
	// State is the
	State []int32
}

// ZwlrForeignToplevelHandleV1DoneEvent is the all information about the toplevel has been sent
type ZwlrForeignToplevelHandleV1DoneEvent struct{}

// ZwlrForeignToplevelHandleV1ClosedEvent is the this toplevel has been destroyed
type ZwlrForeignToplevelHandleV1ClosedEvent struct{}

// ZwlrForeignToplevelHandleV1ParentEvent is the parent change
type ZwlrForeignToplevelHandleV1ParentEvent struct {
	// Parent is the
	Parent *ZwlrForeignToplevelHandleV1
}

// ZwlrForeignToplevelHandleV1TitleHandler is the handler interface for ZwlrForeignToplevelHandleV1TitleEvent
type ZwlrForeignToplevelHandleV1TitleHandler interface {
	HandleZwlrForeignToplevelHandleV1Title(ZwlrForeignToplevelHandleV1TitleEvent)
}

// AddTitleHandler removes the Title handler
func (p *ZwlrForeignToplevelHandleV1) AddTitleHandler(h ZwlrForeignToplevelHandleV1TitleHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1Titles = append(p.privateZwlrForeignToplevelHandleV1Titles, h)
		p.mu.Unlock()
	}
}

// RemoveTitleHandler adds the Title handler
func (p *ZwlrForeignToplevelHandleV1) RemoveTitleHandler(h ZwlrForeignToplevelHandleV1TitleHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1Titles {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1Titles = append(p.privateZwlrForeignToplevelHandleV1Titles[:i], p.privateZwlrForeignToplevelHandleV1Titles[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1AppIdHandler is the handler interface for ZwlrForeignToplevelHandleV1AppIdEvent
type ZwlrForeignToplevelHandleV1AppIdHandler interface {
	HandleZwlrForeignToplevelHandleV1AppId(ZwlrForeignToplevelHandleV1AppIdEvent)
}

// AddAppIdHandler removes the AppId handler
func (p *ZwlrForeignToplevelHandleV1) AddAppIdHandler(h ZwlrForeignToplevelHandleV1AppIdHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1AppIds = append(p.privateZwlrForeignToplevelHandleV1AppIds, h)
		p.mu.Unlock()
	}
}

// RemoveAppIdHandler adds the AppId handler
func (p *ZwlrForeignToplevelHandleV1) RemoveAppIdHandler(h ZwlrForeignToplevelHandleV1AppIdHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1AppIds {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1AppIds = append(p.privateZwlrForeignToplevelHandleV1AppIds[:i], p.privateZwlrForeignToplevelHandleV1AppIds[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1OutputEnterHandler is the handler interface for ZwlrForeignToplevelHandleV1OutputEnterEvent
type ZwlrForeignToplevelHandleV1OutputEnterHandler interface {
	HandleZwlrForeignToplevelHandleV1OutputEnter(ZwlrForeignToplevelHandleV1OutputEnterEvent)
}

// AddOutputEnterHandler removes the OutputEnter handler
func (p *ZwlrForeignToplevelHandleV1) AddOutputEnterHandler(h ZwlrForeignToplevelHandleV1OutputEnterHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1OutputEnters = append(p.privateZwlrForeignToplevelHandleV1OutputEnters, h)
		p.mu.Unlock()
	}
}

// RemoveOutputEnterHandler adds the OutputEnter handler
func (p *ZwlrForeignToplevelHandleV1) RemoveOutputEnterHandler(h ZwlrForeignToplevelHandleV1OutputEnterHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1OutputEnters {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1OutputEnters = append(p.privateZwlrForeignToplevelHandleV1OutputEnters[:i], p.privateZwlrForeignToplevelHandleV1OutputEnters[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1OutputLeaveHandler is the handler interface for ZwlrForeignToplevelHandleV1OutputLeaveEvent
type ZwlrForeignToplevelHandleV1OutputLeaveHandler interface {
	HandleZwlrForeignToplevelHandleV1OutputLeave(ZwlrForeignToplevelHandleV1OutputLeaveEvent)
}

// AddOutputLeaveHandler removes the OutputLeave handler
func (p *ZwlrForeignToplevelHandleV1) AddOutputLeaveHandler(h ZwlrForeignToplevelHandleV1OutputLeaveHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1OutputLeaves = append(p.privateZwlrForeignToplevelHandleV1OutputLeaves, h)
		p.mu.Unlock()
	}
}

// RemoveOutputLeaveHandler adds the OutputLeave handler
func (p *ZwlrForeignToplevelHandleV1) RemoveOutputLeaveHandler(h ZwlrForeignToplevelHandleV1OutputLeaveHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1OutputLeaves {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1OutputLeaves = append(p.privateZwlrForeignToplevelHandleV1OutputLeaves[:i], p.privateZwlrForeignToplevelHandleV1OutputLeaves[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1StateHandler is the handler interface for ZwlrForeignToplevelHandleV1StateEvent
type ZwlrForeignToplevelHandleV1StateHandler interface {
	HandleZwlrForeignToplevelHandleV1State(ZwlrForeignToplevelHandleV1StateEvent)
}

// AddStateHandler removes the State handler
func (p *ZwlrForeignToplevelHandleV1) AddStateHandler(h ZwlrForeignToplevelHandleV1StateHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1States = append(p.privateZwlrForeignToplevelHandleV1States, h)
		p.mu.Unlock()
	}
}

// RemoveStateHandler adds the State handler
func (p *ZwlrForeignToplevelHandleV1) RemoveStateHandler(h ZwlrForeignToplevelHandleV1StateHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1States {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1States = append(p.privateZwlrForeignToplevelHandleV1States[:i], p.privateZwlrForeignToplevelHandleV1States[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1DoneHandler is the handler interface for ZwlrForeignToplevelHandleV1DoneEvent
type ZwlrForeignToplevelHandleV1DoneHandler interface {
	HandleZwlrForeignToplevelHandleV1Done(ZwlrForeignToplevelHandleV1DoneEvent)
}

// AddDoneHandler removes the Done handler
func (p *ZwlrForeignToplevelHandleV1) AddDoneHandler(h ZwlrForeignToplevelHandleV1DoneHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1Dones = append(p.privateZwlrForeignToplevelHandleV1Dones, h)
		p.mu.Unlock()
	}
}

// RemoveDoneHandler adds the Done handler
func (p *ZwlrForeignToplevelHandleV1) RemoveDoneHandler(h ZwlrForeignToplevelHandleV1DoneHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1Dones {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1Dones = append(p.privateZwlrForeignToplevelHandleV1Dones[:i], p.privateZwlrForeignToplevelHandleV1Dones[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1ClosedHandler is the handler interface for ZwlrForeignToplevelHandleV1ClosedEvent
type ZwlrForeignToplevelHandleV1ClosedHandler interface {
	HandleZwlrForeignToplevelHandleV1Closed(ZwlrForeignToplevelHandleV1ClosedEvent)
}

// AddClosedHandler removes the Closed handler
func (p *ZwlrForeignToplevelHandleV1) AddClosedHandler(h ZwlrForeignToplevelHandleV1ClosedHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1Closeds = append(p.privateZwlrForeignToplevelHandleV1Closeds, h)
		p.mu.Unlock()
	}
}

// RemoveClosedHandler adds the Closed handler
func (p *ZwlrForeignToplevelHandleV1) RemoveClosedHandler(h ZwlrForeignToplevelHandleV1ClosedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1Closeds {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1Closeds = append(p.privateZwlrForeignToplevelHandleV1Closeds[:i], p.privateZwlrForeignToplevelHandleV1Closeds[i+1:]...)
			break
		}
	}
}

// ZwlrForeignToplevelHandleV1ParentHandler is the handler interface for ZwlrForeignToplevelHandleV1ParentEvent
type ZwlrForeignToplevelHandleV1ParentHandler interface {
	HandleZwlrForeignToplevelHandleV1Parent(ZwlrForeignToplevelHandleV1ParentEvent)
}

// AddParentHandler removes the Parent handler
func (p *ZwlrForeignToplevelHandleV1) AddParentHandler(h ZwlrForeignToplevelHandleV1ParentHandler) {
	if h != nil {
		p.mu.Lock()
		p.privateZwlrForeignToplevelHandleV1Parents = append(p.privateZwlrForeignToplevelHandleV1Parents, h)
		p.mu.Unlock()
	}
}

// RemoveParentHandler adds the Parent handler
func (p *ZwlrForeignToplevelHandleV1) RemoveParentHandler(h ZwlrForeignToplevelHandleV1ParentHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.privateZwlrForeignToplevelHandleV1Parents {
		if e == h {
			p.privateZwlrForeignToplevelHandleV1Parents = append(p.privateZwlrForeignToplevelHandleV1Parents[:i], p.privateZwlrForeignToplevelHandleV1Parents[i+1:]...)
			break
		}
	}
}
