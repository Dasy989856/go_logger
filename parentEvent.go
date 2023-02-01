package logger

import (
	"net/http"
	"strings"
	"time"
)

func (p *ParentEventStruct) Critical(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:    p.Logger.UserId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:   p.Logger.NameService,
		Package:   p.Package,
		Function:  p.Function,
		Level:     "critical",
		Code:      codeMessage,
		Message:   MapCodes[codeMessage],
		StatusHTTP: http.StatusInternalServerError,
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	if p.Logger.LogLevel <= CriticalLevel {
		p.Logger.Events = append(p.Logger.Events, &event)
	}

	return &event
}

func (p *ParentEventStruct) Error(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:    p.Logger.UserId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:   p.Logger.NameService,
		Package:   p.Package,
		Function:  p.Function,
		Level:     "error",
		Code:      codeMessage,
		Message:   MapCodes[codeMessage],
		StatusHTTP: http.StatusInternalServerError,
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	if p.Logger.LogLevel <= ErrorLevel {
		p.Logger.Events = append(p.Logger.Events, &event)
	}
	return &event
}

func (p *ParentEventStruct) Warning(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:    p.Logger.UserId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:   p.Logger.NameService,
		Package:   p.Package,
		Function:  p.Function,
		Level:     "warning",
		Code:      codeMessage,
		Message:   MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	if p.Logger.LogLevel <= WarningLevel {
		p.Logger.Events = append(p.Logger.Events, &event)
	}
	return &event
}

func (p *ParentEventStruct) Info(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:    p.Logger.UserId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:   p.Logger.NameService,
		Package:   p.Package,
		Function:  p.Function,
		Level:     "info",
		Code:      codeMessage,
		Message:   MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	if p.Logger.LogLevel <= InfoLevel {
		p.Logger.Events = append(p.Logger.Events, &event)
	}
	return &event
}

func (p *ParentEventStruct) Debug(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:    p.Logger.UserId,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:   p.Logger.NameService,
		Package:   p.Package,
		Function:  p.Function,
		Level:     "debug",
		Code:      codeMessage,
		Message:   MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	if p.Logger.LogLevel <= DebugLevel {
		p.Logger.Events = append(p.Logger.Events, &event)
	}
	return &event
}
