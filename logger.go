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

// Создание родительского события. )
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

// Отправка в сервис логирования все события. При ошибки выводит logger в StdOut.
// (Временно - дублирует события типа error и critical в stdOut)
func (l *LoggerStruct) SendToLogService() error {
	if l == nil {
		return fmt.Errorf("nil logger")
	}

	if l.LogServiceAPI == "" {
		l.Print()
		return fmt.Errorf("empty Log service API")
	}

	for _, event := range l.Events {
		if event.Level == "error" || event.Level == "critical" {
			event.Print()
		}

		respLog, err := http.Post(l.LogServiceAPI, "application/json", bytes.NewBuffer(event.ToJson()))
		if err != nil || respLog.StatusCode != http.StatusOK {
			l.Print()
			return err
		}
		defer respLog.Body.Close()
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
		LogLevel Level          `json:"logLevel"`
		Events   []*EventStruct `json:"events"`
	}{
		LogLevel: l.LogLevel,
		Events:   arrEventStructs,
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
// Если в логере есть ошибки, возварается структура с ошибками.
// Если ошибок нет - возвращаются события типа info.
func (l *LoggerStruct) ToJsonForFrontend() []byte {
	if l == nil {
		log.Print("nil logger")
		return nil
	}

	// Проверка логера на ошибки.
	if l.GetNumberOfErrors() != 0 || l.GetNumberOfWarning() != 0 {
		return l.ToJsonForFrontendErrorResponse(l)
	}
	return l.ToJson()
}

func (l *LoggerStruct) ToJsonForFrontendErrorResponse(logger *LoggerStruct) []byte {
	if len(logger.Events) <= 0 {
		return []byte(`{"status":"error", "message":"logger is empty"}`)
	}

	bodyResponse := struct {
		Errors []FrontendEvent `json:"errors"`
	}{}

	for _, event := range logger.Events {
		if event.Context["forFrontend"] != nil && (event.Level == "critical" || event.Level == "error" || event.Level == "warning") {
			var frontendEvent FrontendEvent
			frontendEvent.CreatedAt = event.CreatedAt
			frontendEvent.Code = event.Code
			frontendEvent.Message = event.Message
			if event.Message == "" {
				if msg, ok := event.Context["message"].(string); ok {
					frontendEvent.Message = msg
				}
				if msg, ok := MapCodes[event.Code]; ok {
					frontendEvent.Message = msg
				}
			}
			frontendEvent.Params = event.ParamsMessage
			frontendEvent.Field = event.Context["field"]
			bodyResponse.Errors = append(bodyResponse.Errors, frontendEvent)
		}
	}

	if len(bodyResponse.Errors) == 0 {
		return []byte(`{"status":"not found error for frontend"}`)
	}

	jsonResp, err := json.Marshal(bodyResponse)
	if err != nil {
		context := map[string]interface{}{
			"func":   "toJsonForFrontendErrorResponse",
			"logger": logger,
			"error":  err.Error(),
		}
		log.Print(context)
		return nil
	}
	return jsonResp
}

// Получение HTTP статуса logger (статус крайнего события).
// Если статусы не установлены - вернется 200.
func (l *LoggerStruct) GetStatusHTTP() int {
	for i := len(l.Events) - 1; i >= 0; i-- {
		if l.Events[i].StatusHTTP != 0 {
			return l.Events[i].StatusHTTP
		}
	}
	return http.StatusOK
}

// Возвращает количество событий уровень которых выше или равен error. (error, critical)
func (l *LoggerStruct) GetNumberOfErrors() (counterError int) {
	for _, e := range l.Events {
		if ParseLogLevel(e.Level) >= ErrorLevel {
			counterError++
		}
	}
	return
}

// Возвращает количество событий уровень которых равен warning.
func (l *LoggerStruct) GetNumberOfWarning() (counterWarning int) {
	for _, e := range l.Events {
		if ParseLogLevel(e.Level) == WarningLevel {
			counterWarning++
		}
	}
	return
}

// Возвращает количество событий уровень которых равен info.
func (l *LoggerStruct) GetNumberOfInfo() (counterInfo int) {
	for _, e := range l.Events {
		if ParseLogLevel(e.Level) == InfoLevel {
			counterInfo++
		}
	}
	return
}

// Возвращение всех событий.
func (l *LoggerStruct) GetAllEvents() ([]Event) {
	var sliceEvents []Event
	for _, e := range l.Events {
		sliceEvents = append(sliceEvents, e)
	}
	return sliceEvents
}