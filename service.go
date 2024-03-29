package logger

// Инициализация logger.
func NewLogger(config *Config) Logger {
	logger := LoggerStruct{}
	if config != nil {
		logger.SetConfig(config)
	}
	return &logger
}

type Logger interface {
	// Создание дочернего логера.
	ChildLogger() Logger
	// Установка конфигурации logger.
	SetConfig(*Config)
	// Установка UserID.
	SetUserId(userId int)
	// Инициализация BackgroundEventStruct. (Родитель событий).
	InitParentEvent(packet, function string) ParentEvent
	// Отправка в сервис логирования.
	SendToLogService() error
	// Вывод logger в StdOut в формате Json.
	Print()
	// Получение logger с трасировкой событий в формате Json.
	ToJson() []byte
	// Форматирование logger в формат приемлемый для фронтенда.
	ToJsonForFrontend() []byte
	// Получение HTTP статуса logger (статус крайнего события).
	GetStatusHTTP() (statusHTTP int)
	// Возвращает количество ошибок урвоень которых выше или равен error. (error, critical)
	GetNumberOfErrors() (counterError int)
	// Возвращает количество ошибок урвоень которых равен warning.
	GetNumberOfWarning() (counterWarning int)
	// Возвращает количество ошибок урвоень которых равен info.
	GetNumberOfInfo() (counterInfo int)
	// Возвращение всех событий.
	GetAllEvents() ([]Event) 
}

type ParentEvent interface {
	Critical(codeMessage int, paramsMessage ...string) Event
	Error(codeMessage int, paramsMessage ...string) Event
	Warning(codeMessage int, paramsMessage ...string) Event
	Info(codeMessage int, paramsMessage ...string) Event
	Debug(codeMessage int, paramsMessage ...string) Event
}

type Event interface {
	// Добавление информацию об оригинальной ошибке.
	AddError(err error) Event
	// Установка http статуса.
	SetStatusHTTP(statusHTTP int) Event
	// Получение http статуса.
	GetStatusHTTP() int
	// Добавление контекста.
	AddContext(context map[string]interface{}) Event
	// Получение оригинальной ошибки в формате string.
	Error() string
	// Получение кода события.
	GetCode() int
	// Получения сообщения кода.
	GetMessage() string
	// Установка кастомного сообщения событию.
	SetMessage(message string) Event
	// Получение события в формате Json.
	ToJson() []byte
	// Вывод EventStruct в StdOut в формате Json.
	Print()
}
