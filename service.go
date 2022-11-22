package logger

// Инициализация logger.
func NewLogger(userId int, serviceName, loglevel, logServiceAPI string) Logger {
	logger := LoggerStruct{
		UserId:  userId,
		Service: serviceName,
	}

	switch loglevel {
	case "":
		logger.LogLevel = DebugLevel
	case "DEBUG":
		logger.LogLevel = DebugLevel
	case "INFO":
		logger.LogLevel = InfoLevel
	case "WARNING":
		logger.LogLevel = WarningLevel
	case "ERROR":
		logger.LogLevel = ErrorLevel
	case "CRITICAL":
		logger.LogLevel = CriticalLevel
	}

	return &logger
}

type Logger interface {
	// Добавление API сервиса логирования.
	AddLogServiceAPI(urlAPI string) Logger
	// Добавление UserID в logger.
	AddUserId(userId int) Logger
	// Добавление имени сервиса в logger.
	AddServiceName(serviceName string) Logger
	// Инициализация BackgroundEventStruct. (Родитель событий).
	InitBackgroundEvent(packet, function string) BackgroundEvent
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

type BackgroundEvent interface {
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
	// Отправка события в сервис логирования, при неудаче - запись в StdOut.
	// Дублирование событий тип WARNING, ERROR, CRITICAL в StdOut.
	WriteToDataBase() error
	// Вывод EventStruct в StdOut в формате Json.
	WriteToStdOut()
}
