package logger

const (
	Code_NoCode = 0
)

// Соотношение кода к сообщению.
var MapCodess = initMapCodes()
var MapCodes = map[int]string{
	0: "No code.",
	// CRITICAL (1-999)
	// 1-49 (Server)
	1: "Error starting %v service.",          // Ошибка запуска %v сервиса.
	2: "Configuration initialization error.", // Ошибка инициализации конфигурации.
	3: "Error initializing repository.",      // Ошибка инициализации репозитория.

	// 50-150 (DataBase)
	20: "Error connecting to database %v.",            // Ошибка подключения к БД %v.
	21: "Connection to database %v is not supported.", // Подключение к БД %v не поддерживается.

	// ERROR (1000 - 4999)
	1000: "Reserve",
	1001: "Error sending GET request.",      // Ошибка отправки GET запроса.
	1002: "Error sending POST request.",     // Ошибка отправки POST запроса.
	1003: "Error sending PUT request.",      // Ошибка отправки PUT запроса.
	1004: "Error sending PATCH request.",    // Ошибка отправки PATCH запроса.
	1005: "Error sending DELETE request.",   // Ошибка отправки DELETE запроса.
	1006: "Error sending COPY request.",     // Ошибка отправки COPY запроса.
	1007: "Error sending HEAD request.",     // Ошибка отправки HEAD запроса.
	1008: "Error sending OPTIONS request.",  // Ошибка отправки OPTIONS запроса.
	1009: "Error sending LINK request.",     // Ошибка отправки LINK запроса.
	1010: "Error sending UNLINK request.",   // Ошибка отправки UNLINK запроса.
	1011: "Error sending PURGE request.",    // Ошибка отправки PURGE запроса.
	1012: "Error sending LOCK request.",     // Ошибка отправки LOCK запроса.
	1013: "Error sending UNLOCK request.",   // Ошибка отправки UNLOCK запроса.
	1014: "Error sending PROPFIND request.", // Ошибка отправки PROPFIND запроса.
	1015: "Error sending VIEW request.",     // Ошибка отправки VIEW запроса.

	1020: "Proxy error.", // Ошибка проксирования.

	1050: "JSON encoding error.",        // Ошибка кодирования в JSON.
	1051: "Error decoding JSON.",        // Ошибка декодирования JSON.
	1052: "Error reading request body.", // Ошибка чтения тела запроса.

	1054: "Error upgrading HTTP connection to Web-Socket.",    // Ошибка обновления соединения HTTP до Web-Socket.
	1055: "Error reading message from Web-Socket connection.", // Ошибка чтения сообщения из Web-Socket соединения.
	1056: "Web-Socket connection not found.",                  // Web-Socket соединение не найдено.
	1057: "Error closing Web-Socket connection.",              // Ошибка закрытия Web-Socket соединения.

	1078: "Error creating directory.", // Ошибка создания каталога.
	1079: "File deletion error.",      // Ошибка удаления файла.
	1080: "Error opening file.",       // Ошибка открытия файла.
	1081: "Project not found.",        // Проект не найден.
	1082: "Error getting launch path", // Ошибка получения пути запуска,
	1083: "Error transferring file.",  // Ошибка при передаче файла.

	1097: "",                                               //
	1099: "Empty Refresh token.",                           // Пустой Refresh токен.
	1100: "Empty Access token.",                            // Пустой Access токен.
	1101: "The projectId parameter is empty.",              // Параметр projectId пуст.
	1102: "The userId parameter is empty.",                 // Параметр userId пуст.
	1103: "The userId header is empty.",                    // Заголовок userId пуст.
	1104: "The serviceKey parameter is empty.",             // Параметр serviceKey пуст.
	1105: "Project items are empty.",                       // Элементы проекта пусты.
	1106: "The schema has an unsupported element type.",    // В схеме есть неподдерживаемый тип элемента.
	1107: "At the input port, the output information",      // У входного порта информация выходного.
	1108: "The link refers to a non-existent element.",     // Связь обращается к несуществующему элементу.
	1109: "The link accesses a non-existent element port.", // Связь обращается к несуществующему порту элемента.
	1110: "An input port has an output element.",           // У входного порта есть выходной элемент.
	1111: "An unsupported library type was specified.",     // Указан неподдерживаемый тип библиотеки.
	1112: "Unknown solver type.",                           // Неизвестный тип решателя.

	1150: "The Selector channel is closed.", // Канал Селектора закрыт.
	1151: "The Trend Channel is closed.",    // Канал трендов закрыт.

	1201: "The projectId parameter is invalid",  // Параметр projectId невалиден.
	1202: "The userId parameter is invalid",     // Параметр userId невалиден.
	1203: "The userId header is invalid",        // Заголовок userId невалиден.
	1204: "The serviceKey parameter is invalid", // Параметр serviceKey невалиден.

	1210: "Error converting string to int.",            // Ошибка конвертации string в int.
	1211: "Error parsing address to redirect request.", // Ошибка парсинга адресса для перенаправления запроса.
	// 3000-3999 (DataBase)
	3000: "Reserve",
	3001: "Error writing to database.", // Ошибка записи в БД.

	// WARNING (5000 - 5999)
	4000: "Reserve",

	// INFO (5000 - 7999)
	5000: "Reserve",
	5001: "Successful configuration initialization.",           // Успешная инициализация конфигурации.
	5002: "Successful repository initialization.",              // Успешная инициализации репозитория.
	5003: "Successful routing initialization.",                 // Успешная инициализации роутинга.
	5004: "Successful connection to the database %v.",          // Успешное подключение к БД %v.
	5005: "Successful write to the database.",                  // Успешная запись в БД.
	5006: "Successful launch of the project for calculation.",  // Успешный запуск проекта на расчет.
	5007: "Successful launch of the Selector.",                 // Успешный запуск Селектора.
	5008: "Successful launch of Trends.",                       // Успешный запуск Трендов.
	5009: "Successful closing of the Selector channel.",        // Успешное закрытие канала Селектора.
	5010: "Successful closing of the Trend channel.",           // Успешное закрытие канала Трендов.
	5011: "Successful setting of trend parameters.",            // Успешная установка параметров трендов.
	5012: "Successfully setting a new value to the parameter.", // Успешная установка нового значения параметру.
	5013: "Successful decoding of the request.",                // Успешное декодирование запроса.
	5014: "Successfully setting a new variable value.",         // Успешная установка нового значения переменной.

	5050: "Successful shutdown of the project.", // Успешная остановка проекта.

	5100: "Successful creation of the NoName file.", // Успешное создание NoName файла.
	5101: "Successful file transfer.",               // Успешная передача файлов.

	8001: "Starting a Web-Socket connection.", // Старт подключения Web-Socket соединения.
}

func initMapCodes() map[int]string {
	mapCodes := make(map[int]string, 1000)
	mapCodes[0] = "No code."

	for k, v := range initCodesCritical() {
		mapCodes[k] = v
	}

	for k, v := range initCodesError() {
		mapCodes[k] = v
	}

	for k, v := range initCodesWarning() {
		mapCodes[k] = v
	}

	for k, v := range initCodesInfo() {
		mapCodes[k] = v
	}

	for k, v := range initCodesDebug() {
		mapCodes[k] = v
	}

	return mapCodes
}

// Установка новых соотношений 'код события: сообщение события'
func SetNewMapCodes(newMapCodes map[int]string){
	MapCodes = newMapCodes
}