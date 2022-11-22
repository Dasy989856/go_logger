package logger

import (
	"strings"
	"time"
)

func (l *LoggerStruct) Critical(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:        l.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       l.NameService,
		Package:       l.Package,
		Function:      l.Function,
		Level:         "critical",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	l.Events = append(l.Events, &event)
	return &event
}

func (l *LoggerStruct) Error(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:        l.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       l.NameService,
		Package:       l.Package,
		Function:      l.Function,
		Level:         "error",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	l.Events = append(l.Events, &event)
	return &event
}

func (l *LoggerStruct) Warning(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:        l.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       l.NameService,
		Package:       l.Package,
		Function:      l.Function,
		Level:         "warning",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	l.Events = append(l.Events, &event)
	return &event
}

func (l *LoggerStruct) Info(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:        l.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       l.NameService,
		Package:       l.Package,
		Function:      l.Function,
		Level:         "info",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	l.Events = append(l.Events, &event)
	return &event
}

func (l *LoggerStruct) Debug(codeMessage int, paramsMessage ...string) Event {
	event := EventStruct{
		UserId:        l.UserId,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05.999 Z0700"),
		Service:       l.NameService,
		Package:       l.Package,
		Function:      l.Function,
		Level:         "debug",
		Code:          codeMessage,
		Message:       MapCodes[codeMessage],
	}

	for _, param := range paramsMessage {
		event.Message = strings.Replace(event.Message, "%v", param, 1)
	}

	l.Events = append(l.Events, &event)
	return &event
}