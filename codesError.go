package logger

// ERROR (1000 - 5999)

const (
	Code_ErrorSendingGETrequest            = 1001 // Ошибка отправки GET запроса.
	Code_ErrorSendingPOSTrequest           = 1002 // Ошибка отправки POST запроса.
	Code_ErrorSendingPUTrequest            = 1003 // Ошибка отправки PUT запроса.
	Code_ErrorSendingPATCHrequest          = 1004 // Ошибка отправки PATCH запроса.
	Code_ErrorSendingDELETErequest         = 1005 // Ошибка отправки DELETE запроса.
	Code_ErrorSendingCOPYrequest           = 1006 // Ошибка отправки COPY запроса.
	Code_ErrorSendingHEADrequest           = 1007 // Ошибка отправки HEAD запроса.
	Code_ErrorSendingOPTIONSrequest        = 1008 // Ошибка отправки OPTIONS запроса.
	Code_ErrorSendingLINKrequest           = 1009 // Ошибка отправки LINK запроса.
	Code_ErrorSendingUNLINKrequest         = 1010 // Ошибка отправки UNLINK запроса.
	Code_ErrorSendingPURGErequest          = 1011 // Ошибка отправки PURGE запроса.
	Code_ErrorSendingLOCKrequest           = 1012 // Ошибка отправки LOCK запроса.
	Code_ErrorSendingUNLOCKrequest         = 1013 // Ошибка отправки UNLOCK запроса.
	Code_ErrorSendingPROPFINDrequest       = 1014 // Ошибка отправки PROPFIND запроса.
	Code_ErrorSendingVIEWrequest           = 1015 // Ошибка отправки VIEW запроса.
	Code_ErrorWritingResponseToHTTPrequest = 1016 // Ошибка записи ответа в http запрос.

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
	Code_ErrorConvertingStringToLowerCase       = 1113 // Ошибка конвертации строки в нижний регистр.

	Code_TheSelectorChannelIsClosed = 1150 // Канал Селектора закрыт.
	Code_TheTrendChannelIsClosed    = 1151 // Канал трендов закрыт.

	// Validation
	Code_TheProjectIdParameterIsInvalid                   = 1201 // Параметр projectId невалиден.
	Code_TheUserIdParameterIsInvalid                      = 1202 // Параметр userId невалиден.
	Code_TheUserIdHeaderIsInvalid                         = 1203 // Заголовок userId невалиден.
	Code_TheServiceKeyParameterIsInvalid                  = 1204 // Параметр serviceKey невалиден.
	Code_LoginIsAlreadyInUse                              = 1205 // Логин уже в использовании.
	Code_ThePhoneIsAlreadyInUse                           = 1206 // Телефон уже в использовании.
	Code_EmailIsAlreadyInUse                              = 1207 // Электронная почта уже в использовании.
	Code_LoginCannotBeEmpty                               = 1208 // Логин не может быть пустым.
	Code_ThePhoneCannotBeEmpty                            = 1209 // Телефон не может быть пустым.
	Code_EmailCannotBeEmpty                               = 1210 // Электронная почта не может быть пуста.
	Code_ErrorParsingAddressToRedirectRequest             = 1211 // Ошибка парсинга адресса для перенаправления запроса.
	Code_ErrorConvertingStringToInt                       = 1212 // Ошибка конвертации string в int.
	Code_TheCountryCodeOfThePhoneNumberMustBeginWithAPlus = 1213 // Код страны номера телефона должен начинаться с +.
	Code_InvalidPhoneFormat                               = 1214 // Неверный формат телефона.
	Code_ThePasswordCannotBeEmpty                         = 1215 // Пароль не может быть пустым.
	Code_TheAccountTypeCannotBeEmpty                      = 1216 // Тип аккаунта не может быть пустым.
	Code_TheUsernameCannotBeEmpty                         = 1217 // Имя пользователя не может быть пустым.
	Code_TheFullCompanyNameCannotBeEmpty                  = 1218 // Полное наименование компании не может быть пустым.
	Code_LegalAddressCannotBeEmpty                        = 1219 // Юридический адрес не может быть пустым.
	Code_PostalAddressCannotBeEmpty                       = 1220 // Почтовый адрес не может быть пустым.
	Code_CurrentAccountCannotBeEmpty                      = 1221 // Расчетный счет не может быть пустым.
	Code_BankNameCannotBeEmpty                            = 1222 // Наименование банка не может быть пустым.
	Code_BICcannotBeEmpty                                 = 1223 // БИК не может быть пустым.
	Code_INNcannotBeEmpty                                 = 1224 // ИНН банка не может быть пустым.
	Code_KPPcannotBeEmpty                                 = 1225 // КПП банка не может быть пустым.
	Code_CorrespondentAccountCannotBeEmpty                = 1226 // Корреспондентский счет не может быть пустым.
	Code_PasswordsDoNotMatch                              = 1227 // Пароли не совпадают.
	Code_OGRNIPcannotBeEmpty                              = 1228 // ОГРНИП не может быть пустым.
	Code_OGRNcannotBeEmpty                                = 1229 // ОГРН не может быть пустым.
	Code_KIOcannotBeEmpty                                 = 1230 // КИО не может быть пустым.
	Code_TheLicenseCannotBeEmpty                          = 1231 // Лицензия не может быть пустым.
	Code_TheLicenseDateCannotBeEmpty                      = 1232 // Дата лицензии не может быть пустым.
	Code_AccountTypeIsNotSupported                        = 1233 // Тип аккаунта не поддерживается.

	// DataBase
	Code_ErrorWritingToDatabase   = 3001 // Ошибка записи в БД.
	Code_ErrorReadingFromDatabase = 3002 // Ошибка чтения из БД.
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
		1113: "Error converting string to lower case.",         // Ошибка конвертации строки в нижний регистр.

		1150: "The Selector channel is closed.", // Канал Селектора закрыт.
		1151: "The Trend Channel is closed.",    // Канал трендов закрыт.

		1201: "The projectId parameter is invalid",                      // Параметр projectId невалиден.
		1202: "The userId parameter is invalid",                         // Параметр userId невалиден.
		1203: "The userId header is invalid",                            // Заголовок userId невалиден.
		1204: "The serviceKey parameter is invalid",                     // Параметр serviceKey невалиден.
		1205: "Login is already in use.",                                // Логин уже в использовании.
		1206: "The phone is already in use.",                            // Телефон уже в использовании.
		1207: "Email is already in use.",                                // Электронная почта уже в использовании.
		1208: "Login cannot be empty.",                                  // Логин не может быть пустым.
		1209: "The phone cannot be empty.",                              // Телефон не может быть пустым.
		1210: "Email cannot be empty.",                                  // Электронная почта не может быть пуста.
		1211: "Error parsing address to redirect request.",              // Ошибка парсинга адресса для перенаправления запроса.
		1212: "Error converting string to int.",                         // Ошибка конвертации string в int.
		1213: "The country code of the phone number must begin with +.", // Код страны номера телефона должен начинаться с +.
		1214: "Invalid phone format.",                                   // Неверный формат телефона.
		1215: "The password cannot be empty.",                           // Пароль не может быть пустым.
		1216: "The account type cannot be empty.",                       // Тип аккаунта не может быть пустым.
		1217: "The username cannot be empty.",                           // Имя пользователя не может быть пустым.
		1218: "The full company name cannot be empty.",                  // Полное наименование компании не может быть пустым.
		1219: "Legal address cannot be empty.",                          // Юридический адрес не может быть пустым.
		1220: "Postal address cannot be empty.",                         // Почтовый адрес не может быть пустым.
		1221: "Current account cannot be empty.",                        // Расчетный счет не может быть пустым.
		1222: "Bank name cannot be empty.",                              // Наименование банка не может быть пустым.
		1223: "BIC cannot be empty.",                                    // БИК не может быть пустым.
		1224: "INN cannot be empty.",                                    // ИНН банка не может быть пустым.
		1225: "KPP cannot be empty.",                                    // КПП банка не может быть пустым.
		1226: "Correspondent account cannot be empty.",                  // Корреспондентский счет не может быть пустым.
		1227: "Passwords do not match.",                                 // Пароли не совпадают.
		1228: "OGRNIP cannot be empty.",                                 // ОГРНИП не может быть пустым.
		1229: "OGRN cannot be empty.",                                   // ОГРН не может быть пустым.
		1230: "KIO cannot be empty.",                                    // КИО не может быть пустым.
		1231: "The license cannot be empty.",                            // Лицензия не может быть пустым.
		1232: "The license date cannot be empty.",                       // Дата лицензии не может быть пустым.
		1233: "Account type is not supported.",                          // Тип аккаунта не поддерживается.

		//DataBase
		3000: "Reserve",
		3001: "Error writing to database.",   // Ошибка записи в БД.
		3002: "Error reading from database.", // Ошибка чтения из БД.
	}

	return mapCodesError
}
