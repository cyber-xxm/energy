//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ipc

// EventCallback IPC context callback
type EventCallback func(context IEvent)

type ICallback interface {
	Invoke(context IEvent)
}

// Callback IPC Listening callback function
type Callback struct {
	callback EventCallback
}

// Invoke event function
func (m *Callback) Invoke(context IEvent) {
	if m.callback != nil {
		m.callback(context)
	}
}

// IEvent
//
//	Inter process IPC communication callback context
type IEvent interface {
	Data() interface{}          //ArgumentList
	Result(data ...interface{}) //callback function return Result
}

// Event IPC Event
type Event struct {
	windowId uint
	frameId  uint
	data     interface{}
	result   []interface{}
}

// NewEvent create IPC event message
func NewEvent(windowId, frameId uint, data interface{}) IEvent {
	ctx := &Event{
		windowId: windowId,
		frameId:  frameId,
		data:     data,
		result:   nil,
	}
	return ctx
}

func (m *Event) Data() interface{} {
	return m.data
}

func (m *Event) Result(results ...interface{}) {
	m.result = results
}