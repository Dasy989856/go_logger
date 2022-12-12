package logger

// ERROR (1000 - 5999)

const (
	Code_ErrorSendingGETrequest      = 1001 // Ошибка отправки GET запроса.
	Code_ErrorSendingPOSTrequest     = 1002 // Ошибка отправки POST запроса.
	Code_ErrorSendingPUTrequest      = 1003 // Ошибка отправки PUT запроса.
	Code_ErrorSendingPATCHrequest    = 1004 // Ошибка отправки PATCH запроса.
	Code_ErrorSendingDELETErequest   = 1005 // Ошибка отправки DELETE запроса.
	Code_ErrorSendingCOPYrequest     = 1006 // Ошибка отправки COPY запроса.
	Code_ErrorSendingHEADrequest     = 1007 // Ошибка отправки HEAD запроса.
	Code_ErrorSendingOPTIONSrequest  = 1008 // Ошибка отправки OPTIONS запроса.
	Code_ErrorSendingLINKrequest     = 1009 // Ошибка отправки LINK запроса.
	Code_ErrorSendingUNLINKrequest   = 1010 // Ошибка отправки UNLINK запроса.
	Code_ErrorSendingPURGErequest    = 1011 // Ошибка отправки PURGE запроса.
	Code_ErrorSendingLOCKrequest     = 1012 // Ошибка отправки LOCK запроса.
	Code_ErrorSendingUNLOCKrequest   = 1013 // Ошибка отправки UNLOCK запроса.
	Code_ErrorSendingPROPFINDrequest = 1014 // Ошибка отправки PROPFIND запроса.
	Code_ErrorSendingVIEWrequest     = 1015 // Ошибка отправки VIEW запроса.

	Code_ProxyError = 1020 // Ошибка проксирования.

	Code_JsonEncodingError       = 1050 // Ошибка кодирования в JSON.
	Code_ErrorDecodingJson       = 1051 // Ошибка декодирования JSON.
	Code_ErrorReadingRequestBody = 1052 // Ошибка чтения тела запроса.

	Code_ErrorUpgradingHTTPconnectionToWebSocket    = 1054 // Ошибка обновления соединения HTTP до Web-Socket.
	Code_ErrorReadingMessageFromWebSocketConnection = 1055 // Ошибка чтения сообщения из Web-Socket соединения.
	Code_WebSocketConnectionNotFound                = 1056 // Web-Socket соединение не найдено.
	Code_ErrorClosingWebSocketConnection            = 1057 // Ошибка закрытия Web-Socket соединения.

	Code_ErrorCreatingDirectory = 1078 // Ошибка создания каталога.
	Code_FileDeletionError      = 1079 // Ошибка удаления файла.
	Code_ErrorOpeningFile       = 1080 // Ошибка открытия файла.
	Code_ProjectNotFound        = 1081 // Проект не найден.
	Code_ErrorGettingLaunchPath = 1082 // Ошибка получения пути запуска.
	Code_ErrorTransferringFile  = 1083 // Ошибка при передаче файла.

	Code_InvalidRefreshToken                    = 1097 // Некорректный Refresh токен.
	Code_InvalidAccessToken                     = 1098 // Некорректный Access токен.
	Code_EmptyRefreshToken                      = 1099 // Пустой Refresh токен.
	Code_EmptyAccessToken                       = 1100 // Пустой Access токен.
	Code_TheProjectIdParameterIsEmpty           = 1101 // Параметр projectId пуст.
	Code_TheUserIdParameterIsEmpty              = 1102 // Параметр userId пуст.
	Code_TheUserIdHeaderIsEmpty                 = 1103 // Заголовок userId пуст.
	Code_TheServiceKeyParameterIsEmpty          = 1104 // Параметр serviceKey пуст.
	Code_ProjectItemsAreEmpty                   = 1105 // Элементы проекта пусты.
	Code_TheSchemaHasAnUnsupportedElementType   = 1106 // В схеме есть неподдерживаемый тип элемента.
	Code_AtTheInputPortTheOutputInformation     = 1107 // У входного порта информация выходного.
	Code_TheLinkRefersToANonExistentElement     = 1108 // Связь обращается к несуществующему элементу.
	Code_TheLinkAccessesANonExistentElementPort = 1109 // Связь обращается к несуществующему порту элемента.
	Code_AnInputPortHasAnOutputElement          = 1110 // У входного порта есть выходной элемент.
	Code_AnUnsupportedLibraryTypeWasSpecified   = 1111 // Указан неподдерживаемый тип библиотеки.
	Code_UnknownSolverType                      = 1112 // Неизвестный тип решателя.

	Code_TheSelectorChannelIsClosed = 1150 // Канал Селектора закрыт.
	Code_TheTrendChannelIsClosed    = 1151 // Канал трендов закрыт.

	Code_TheProjectIdParameterIsInvalid  = 1201 // Параметр projectId невалиден.
	Code_TheUserIdParameterIsInvalid     = 1202 // Параметр userId невалиден.
	Code_TheUserIdHeaderIsInvalid        = 1203 // Заголовок userId невалиден.
	Code_TheServiceKeyParameterIsInvalid = 1203 // Параметр serviceKey невалиден.

	Code_ErrorConvertingStringToInt           = 1210 // Ошибка конвертации string в int.
	Code_ErrorParsingAddressToRedirectRequest = 1211 // Ошибка парсинга адресса для перенаправления запроса.

	Code_ErrorWritingToDatabase = 3001 // Ошибка записи в БД.
)

func initCodesError() map[int]string {
	mapCodesError := map[int]string{
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
	}

	return mapCodesError
}
