package logger

type Level int8

const (
	DebugLevel Level = iota + 1
	InfoLevel
	WarningLevel
	ErrorLevel
	CriticalLevel
)

type Config struct {
	LogLevel      string
	LogServiceAPI string
	UserId        int
	NameService   string
}

// Структура логера.
type LoggerStruct struct {
	LogLevel      Level           `json:"logLevel,omitempty"`      // Уровень логирования.
	LogServiceAPI string          `json:"logServiceAPI,omitempty"` // API сервиса логирования.
	Events        []*EventStruct  `json:"events,omitempty"`        // Массив событий.
	UserId        int             `json:"-"`                       // ID пользователя.
	NameService   string          `json:"-"`                       // Название сервсиса.
	ChildLoggers  []*LoggerStruct `json:"-"`
}

type ParentEventStruct struct {
	Logger   *LoggerStruct
	Package  string // Название пакета.
	Function string // Название функции.
}

// Данные события
type EventStruct struct {
	LogServiceAPI string                 `json:"-"`                       // API сервиса логирования.
	UserId        int                    `json:"userId"`                  // ID пользователя.
	CreatedAt     string                 `json:"createdAt"`               // Дата создания.
	Level         string                 `json:"level"`                   // Уровень.
	Service       string                 `json:"service"`                 // Название сервсиса.
	Package       string                 `json:"package"`                 // Название пакета.
	Function      string                 `json:"function"`                // Название функции.
	Code          int                    `json:"code"`                    // Код события.
	Message       string                 `json:"message"`                 // Сообщение.
	ParamsMessage []string               `json:"paramsMessage,omitempty"` // Параметры для сообщения.
	Context       map[string]interface{} `json:"context,omitempty"`       // Контекст (дополнительный данные).
	OriginalError string                 `json:"originalError,omitempty"` // Оригинальная ошибка от стороних пакетов.
	StatusHTTP    int                    `json:"statusHTTP,omitempty"`    // Код HTTP ответа. (200, 400, 404, 500...)
}

// Структура ошибки в формате для Фронтенд.
type FrontError struct {
	Code    int         `json:"code"`             // Код ошибки.
	Message string      `json:"message"`          // Сообщение ошибки.
	Params  []string    `json:"params,omitempty"` // Параметры сообщения.
	Field   interface{} `json:"field,omitempty"`  // Поле(ключ) Json в котором произошла ошибка.
}
