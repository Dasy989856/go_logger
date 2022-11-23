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

// Установка конфигурации logger.
func (l *LoggerStruct) SetConfig(config *Config) {
	if config == nil {
		return
	}

	if config.UserId != 0 {
		l.UserId = config.UserId
	}

	if config.LogLevel != "" {
		l.LogLevel = ParseLogLevel(config.LogLevel)
	}
}

// Создание родительского события.
func (l *LoggerStruct) InitParentEvent(packet, function string) ParentEvent {
	if l == nil {
		log.Print(fmt.Errorf("nil logger"))
		return l
	}

	l.Package = packet
	l.Function = function
	return l
}

// Отправка в сервис логирования.
func (l *LoggerStruct) SendToLogService() error {
	if l == nil {
		return fmt.Errorf("nil logger")
	}

	if l.LogServiceAPI == "" {
		return fmt.Errorf("empty Log service API")
	}

	for _, EventStruct := range l.Events {
		jsonEventStruct, err := EventStruct.ToJson()
		if err != nil {
			return err
		}
		// Отправка лога и получение ответа от сервиса логирования.
		respLog, err := http.Post(l.LogServiceAPI, "application/json", bytes.NewBuffer(jsonEventStruct))
		if err != nil || respLog.StatusCode != http.StatusOK {
			return err
		}
	}

	return nil
}

// Вывод logger в StdOut в формате Json.
func (l *LoggerStruct) WriteToStdOut() {
	if l == nil {
		log.Print(fmt.Errorf("nil logger"))
		return
	}

	logerJson, err := l.ToJson()
	if err != nil {
		log.Print(err)
	}

	fmt.Println(strings.Repeat("=", 25), "LOGGER", strings.Repeat("=", 25))
	fmt.Println(string(logerJson))
	fmt.Println(strings.Repeat("=", 60))
}

// Получение logger с трасировкой событий в формате Json.
func (l *LoggerStruct) ToJson() ([]byte, error) {
	if l == nil {
		return nil, fmt.Errorf("nil logger")
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
	}
	return jsonLogger, nil
}

// Форматирование событий в формат приемлемый для фронтенда.
func (l *LoggerStruct) ToFrontendJson() ([]byte, error) {
	if l == nil {
		return nil, fmt.Errorf("nil logger")
	}

	bodyResponse := struct {
		Errors []FrontError `json:"errors"`
	}{}

	for _, log := range l.Events {
		if log.Level != "ERROR" {
			continue
		}

		var frontLog FrontError
		frontLog.Code = log.Code
		frontLog.Message = log.Message
		frontLog.Params = log.ParamsMessage
		frontLog.Field = log.Context["field"]

		bodyResponse.Errors = append(bodyResponse.Errors, frontLog)
	}

	jsonResp, err := json.Marshal(bodyResponse)
	if err != nil {
		return nil, err
	}
	return jsonResp, nil
}

// Получение HTTP статуса logger (статус крайнего события).
func (l *LoggerStruct) GetStatusHTTP() int {
	for _, ev := range l.Events {
		if ev.Level == "WARNING" || ev.Level == "ERROR" || ev.Level == "CRITICAL" {
			return ev.StatusHTTP
		}
	}
	return http.StatusOK
}
