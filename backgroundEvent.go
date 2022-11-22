package logger

import (
	"strings"
	"time"
)

func (b *BackgroundEventStruct) Critical(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		LogServiceAPI: b.Logger.LogServiceAPI,
		UserId:        b.Logger.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       b.Logger.Service,
		Package:       b.Package,
		Function:      b.Function,
		Level:         "CRITICAL",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	b.Logger.Events = append(b.Logger.Events, &event)
	return &event
}

func (b *BackgroundEventStruct) Error(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		LogServiceAPI: b.Logger.LogServiceAPI,
		UserId:        b.Logger.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       b.Logger.Service,
		Package:       b.Package,
		Function:      b.Function,
		Level:         "ERROR",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	b.Logger.Events = append(b.Logger.Events, &event)
	return &event
}

func (b *BackgroundEventStruct) Warning(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		LogServiceAPI: b.Logger.LogServiceAPI,
		UserId:        b.Logger.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       b.Logger.Service,
		Package:       b.Package,
		Function:      b.Function,
		Level:         "WARNING",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	b.Logger.Events = append(b.Logger.Events, &event)
	return &event
}

func (b *BackgroundEventStruct) Info(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		LogServiceAPI: b.Logger.LogServiceAPI,
		UserId:        b.Logger.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       b.Logger.Service,
		Package:       b.Package,
		Function:      b.Function,
		Level:         "INFO",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	b.Logger.Events = append(b.Logger.Events, &event)
	return &event
}

func (b *BackgroundEventStruct) Debug(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		LogServiceAPI: b.Logger.LogServiceAPI,
		UserId:        b.Logger.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       b.Logger.Service,
		Package:       b.Package,
		Function:      b.Function,
		Level:         "DEBUG",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	b.Logger.Events = append(b.Logger.Events, &event)
	return &event
}