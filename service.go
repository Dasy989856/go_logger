package logger

// Инициализация logger. 
func NewLogger(config *Confing) Logger {
	var logger *LoggerStruct
	if config != nil {
		logger.SetConfig(config)
	}
	return logger
}

type Logger interface {
	// Установка конфигурации logger.
	SetConfig(*Confing)
	// Инициализация BackgroundEventStruct. (Родитель событий).
	InitParentEvent(packet, function string) ParentEvent
	// Отправка в сервис логирования.
	SendToLogService() error
	// Вывод logger в StdOut в формате Json.
	WriteToStdOut()
	// Получение logger с трасировкой событий в формате Json.
	ToJson() ([]byte, error)
	// Форматирование logger в формат приемлемый для фронтенда.
	ToFrontendJson() ([]byte, error)
	// Получение HTTP статуса logger (статус крайнего события).
	GetStatusHTTP() int
}

type ParentEvent interface {
	Critical(codeMessage int, paramsMessage ...string) Event
	Error(codeMessage int, paramsMessage ...string) Event
	Warning(codeMessage int, paramsMessage ...string) Event
	Info(codeMessage int, paramsMessage ...string) Event
	Debug(codeMessage int, paramsMessage ...string) Event
}

type Event interface {
	// Добавление информацию об оригинальной ошибке и http статусе.
	AddError(err error, httpStatus int) Event
	// Добавление контекста.
	AddContext(context map[string]interface{}) Event
	// Получение оригинальной ошибки в формате string.
	Error() string
	// Получение события в формате Json.
	ToJson() ([]byte, error)
	// Вывод EventStruct в StdOut в формате Json.
	WriteToStdOut()
}
