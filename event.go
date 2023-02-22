package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Добавление оригинальной ошибки.
func (e *EventStruct) AddError(err error) Event {
	if e == nil {
		log.Print(fmt.Errorf("nil event"))
		return e
	}

	if err != nil {
		e.OriginalError = err.Error()
	} else {
		e.OriginalError = "error nil"
	}

	if e.StatusHTTP == 0 {
		e.StatusHTTP = http.StatusInternalServerError
	}

	return e
}

// Установка http статуса.
func (e *EventStruct) SetStatusHTTP(statusHTTP int) Event {
	if e == nil {
		log.Print(fmt.Errorf("nil event"))
		return e
	}

	e.StatusHTTP = statusHTTP
	return e
}

// Получение http статуса.
func (e *EventStruct) GetStatusHTTP() int {
	if e == nil {
		log.Print(fmt.Errorf("nil event"))
		return 500
	}
	return e.StatusHTTP
}

// Добавление контекста.
func (e *EventStruct) AddContext(context map[string]interface{}) Event {
	if e == nil {
		log.Print(fmt.Errorf("nil event"))
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

// Получение кода ошибки
func (e *EventStruct) GetCode() int {
	return e.Code
}

// Получения сообщения кода.
func (e *EventStruct) GetMessage() string {
	return e.Message
}

// Установка кастомного сообщения событию.
func (e *EventStruct) SetMessage(message string) Event {
	e.Message = message
	return e
}

// Получение события в формате Json.
// TODO перевести в ручное форматирование к JSON. `https://go.dev/play/p/SH5bsrjzB06`
func (e *EventStruct) ToJson() []byte {
	if e == nil {
		log.Print("nil event")
		return nil
	}

	jsEvent, err := json.Marshal(e)
	if err != nil {
		log.Print(err)
		return nil
	}
	return jsEvent
}

// Вывод EventStruct в StdOut в формате Json.
func (e *EventStruct) Print() {
	if e == nil {
		log.Print(fmt.Errorf("nil event"))
		return
	}

	fmt.Println(strings.Repeat("=", 25), "EVENT", strings.Repeat("=", 25))
	fmt.Println(string(e.ToJson()))
	fmt.Println(strings.Repeat("=", 60))
}

func (e *EventStruct) SendToLogService() error {
	if e == nil {
		return fmt.Errorf("nil logger")
	}

	if e.LogServiceAPI == "" {
		return fmt.Errorf("empty Log service API")
	}

	if e.Level == "warning" || e.Level == "error" || e.Level == "critical" {
		e.Print()
	}

	respLog, err := http.Post(e.LogServiceAPI, "application/json", bytes.NewBuffer(e.ToJson()))
	if err != nil {
		log.Print(err)
		return err
	}
	defer respLog.Body.Close()
	
	if respLog.StatusCode != http.StatusOK {
		log.Print("response from Log Service not 200")
		return err
	}

	return nil
}
