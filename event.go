package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Добавление информацию об оригинальной ошибке и http статусе.
func (e *EventStruct) AddError(err error, httpStatus int) Event {
	if e == nil {
		log.Print(fmt.Errorf("nil EventStruct"))
		return e
	}

	if err != nil {
		e.OriginalError = err.Error()
		e.StatusHTTP = httpStatus
	} else {
		e.OriginalError = "error nil"
		e.StatusHTTP = http.StatusInternalServerError
	}
	return e
}

// Добавление контекста.
func (e *EventStruct) AddContext(context map[string]interface{}) Event {
	if e == nil {
		log.Print(fmt.Errorf("nil EventStruct"))
		return e
	}

	if context == nil {
		e.Context = make(map[string]interface{})
	}

	e.Context = context
	return e
}

// Получение оригинальной ошибки в формате string.
func (e *EventStruct) Error() string {
	if e == nil {
		return ""
	}
	return e.OriginalError
}

// Получение события в формате Json.
func (e *EventStruct) ToJson() ([]byte, error) {
	if e == nil {
		return nil, fmt.Errorf("nil EventStruct")
	}

	js, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return js, nil
}

// Вывод EventStruct в StdOut в формате Json.
func (e *EventStruct) WriteToStdOut() {
	if e == nil {
		log.Print(fmt.Errorf("nil EventStruct"))
		return
	}

	EventStructJson, err := e.ToJson()
	if err != nil {
		context := map[string]interface{}{
			"fucn":        "WriteToStdOut",
			"EventStruct": e,
			"error":       err.Error(),
		}
		log.Print(context)
		return
	}

	fmt.Println(strings.Repeat("=", 25), "EventStruct", strings.Repeat("=", 25))
	fmt.Println(string(EventStructJson))
	fmt.Println(strings.Repeat("=", 60))
}
