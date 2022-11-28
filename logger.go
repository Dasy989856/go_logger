package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ParseLogLevel(loglevel string) Level {
	switch strings.ToLower(loglevel) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "critical":
		return CriticalLevel
	default:
		return DebugLevel
	}
}

func (l *LoggerStruct) ChildLogger() Logger {
	return &LoggerStruct{
		LogLevel:      l.LogLevel,
		LogServiceAPI: l.LogServiceAPI,
		UserId:        l.UserId,
		NameService:   l.NameService,
	}
}

// Установка конфигурации logger.
func (l *LoggerStruct) SetConfig(config *Config) {
	if config == nil {
		log.Print(fmt.Errorf("nil config"))
		return
	}

	if config.LogLevel != "" {
		l.LogLevel = ParseLogLevel(config.LogLevel)
	}

	if config.LogServiceAPI != "" {
		l.LogServiceAPI = config.LogServiceAPI
	}

	if config.UserId != 0 {
		l.UserId = config.UserId
	}

	if config.NameService != "" {
		l.NameService = config.NameService
	}
}

// Установка конфигурации logger.
func (l *LoggerStruct) SetUserId(userId int) {
	l.UserId = userId
}

// Создание родительского события.
func (l *LoggerStruct) InitParentEvent(packet, function string) ParentEvent {
	if l == nil {
		log.Print(fmt.Errorf("nil logger"))
		return nil
	}

	pEvent := ParentEventStruct{
		Logger:   l,
		Package:  packet,
		Function: function,
	}

	return &pEvent
}

// Отправка в сервис логирования.
func (l *LoggerStruct) SendToLogService() error {
	if l == nil {
		return fmt.Errorf("nil logger")
	}

	if l.LogServiceAPI == "" {
		return fmt.Errorf("empty Log service API")
	}

	for i, event := range l.Events {
		if event.Level == "warning" || event.Level == "error" || event.Level == "critical" {
			event.Print()
		}

		respLog, err := http.Post(l.LogServiceAPI, "application/json", bytes.NewBuffer(event.ToJson()))
		if err != nil || respLog.StatusCode != http.StatusOK {
			return err
		}
		l.Events = l.Events[i:]
	}

	return nil
}

// Вывод logger в StdOut в формате Json.
func (l *LoggerStruct) Print() {
	if l == nil {
		log.Print(fmt.Errorf("nil logger"))
		return
	}

	if len(l.Events) == 0 {
		return
	}

	fmt.Println(strings.Repeat("=", 25), "LOGGER", strings.Repeat("=", 25))
	fmt.Println(string(l.ToJson()))
	fmt.Println(strings.Repeat("=", 60))
}

// Получение logger с трасировкой событий в формате Json.
func (l *LoggerStruct) ToJson() []byte {
	if l == nil {
		log.Print("nil logger")
		return nil
	}

	arrEventStructs := make([]*EventStruct, 0, 5)
	// Вывод в правильной последовтельности трасировки ошибок.
	for i := len(l.Events) - 1; i >= 0; i-- {
		arrEventStructs = append(arrEventStructs, l.Events[i])
	}

	logger := struct {
		LogLevel     Level          `json:"logLevel"`
		EventStructs []*EventStruct `json:"EventStructs"`
	}{
		LogLevel:     l.LogLevel,
		EventStructs: arrEventStructs,
	}

	jsonLogger, err := json.Marshal(logger)
	if err != nil {
		context := map[string]interface{}{
			"logger": logger,
			"error":  err.Error(),
		}
		log.Print(context)
		return nil
	}
	return jsonLogger
}

// Форматирование событий в формат приемлемый для фронтенда.
func (l *LoggerStruct) ToJsonForFrontend() []byte {
	if l == nil {
		log.Print("nil logger")
		return nil
	}

	// Проверка логера на ошибки.
	for _, event := range l.Events {
		if event.Level == "critical" || event.Level == "error" || event.Level == "warning" {
			return toJsonForFrontendErrorResponse(l)
		}
	}
	return toJsonForFrontendResponse(l)
}

func toJsonForFrontendErrorResponse(logger *LoggerStruct) []byte {
	bodyResponse := struct {
		Events []FrontendEvent `json:"events"`
	}{}

	for _, event := range logger.Events {
		if event.Level == "critical" || event.Level == "error" || event.Level == "warning" {
			var frontendEvent FrontendEvent
			frontendEvent.Code = event.Code
			frontendEvent.Message = MapCodes[event.Code]
			frontendEvent.Params = event.ParamsMessage
			frontendEvent.Field = event.Context["field"]
			bodyResponse.Events = append(bodyResponse.Events, frontendEvent)
		}
	}

	jsonResp, err := json.Marshal(bodyResponse)
	if err != nil {
		context := map[string]interface{}{
			"func": "toJsonForFrontendErrorResponse",
			"logger": logger,
			"error":  err.Error(),
		}
		log.Print(context)
		return nil
	}
	return jsonResp
}

func toJsonForFrontendResponse(logger *LoggerStruct) []byte {
	bodyResponse := struct {
		Events []FrontendEvent `json:"events"`
	}{}

	for _, event := range logger.Events {
		if event.Level == "info" {
			var frontendEvent FrontendEvent
			frontendEvent.Code = event.Code
			frontendEvent.Message = MapCodes[event.Code]
			frontendEvent.Params = event.ParamsMessage
			frontendEvent.Field = event.Context["field"]
			bodyResponse.Events = append(bodyResponse.Events, frontendEvent)
		}
	}

	jsonResp, err := json.Marshal(bodyResponse)
	if err != nil {
		context := map[string]interface{}{
			"func": "toJsonForFrontendResponse",
			"logger": logger,
			"error":  err.Error(),
		}
		log.Print(context)
		return nil
	}
	return jsonResp
}

// Получение HTTP статуса logger (статус крайнего события).
func (l *LoggerStruct) GetStatusHTTP() int {
	for _, ev := range l.Events {
		if ev.Level == "warning" || ev.Level == "error" || ev.Level == "critical" {
			return ev.StatusHTTP
		}
	}
	return http.StatusOK
}
