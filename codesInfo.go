package logger

// INFO (7000 - 8999)

const (
	Code_SuccessfulConfigurationInitialization       = 7001 // Успешная инициализация конфигурации.
	Code_SuccessfulRepositoryInitialization          = 7002 // Успешная инициализации репозитория.
	Code_SuccessfulRoutingInitialization             = 7003 // Успешная инициализации роутинга.
	Code_SuccessfulConnectionToTheDatabase__NameDB__ = 7004 // Успешное подключение к БД %v.
	Code_SuccessfulWriteToTheDataBase                = 7005 // Успешная запись в БД.
	Code_SuccessfulLaunchOfTheProjectForCalculation  = 7006 // Успешный запуск проекта на расчет.
	Code_SuccessfulLaunchOfTheSelector               = 7007 // Успешный запуск Селектора.
	Code_SuccessfulLaunchOfTrends                    = 7008 // Успешный запуск Трендов.
	Code_SuccessfulClosingOfTheSelectorChannel       = 7009 // Успешное закрытие канала Селектора.
	Code_SuccessfulClosingOfTheTrendChannel          = 7010 // Успешное закрытие канала Трендов.
	Code_SuccessfulSettingOfTrendParameters          = 7011 // Успешная установка параметров трендов.
	Code_SuccessfullySettingANewValueToTheParameter  = 7012 // Успешная установка нового значения параметру.
	Code_SuccessfulDecodingOfTheRequest              = 7013 // Успешное декодирование запроса.
	Code_SuccessfullySettingANewVariableValue        = 7014 // Успешная установка нового значения переменной.

	Code_SuccessfulShutdownOfTheProject = 7050 // Успешная остановка проекта.

	Code_SuccessfulCreationOfTheNoNameFile = 7100 // Успешное создание NoName файла.
	Code_SuccessfulFileTransfer            = 7101 // Успешная передача файлов.

	Code_SuccessfulCompanyCreationInTheDataBase = 7200 // Успешное создание компании в базе данных.

	Code_SuccessfulSendingOfAnEmailFromPasswordRecoveryToken = 7300 // Успешная отправка письма с токеном восстановления пароля.

	Code_StartingAWebSocketConnection = 8001 // Старт подключения Web-Socket соединения.
)

func initCodesInfo() map[int]string {
	mapCodesInfo := map[int]string{
		7000: "Reserve",
		7001: "Successful configuration initialization.",           // Успешная инициализация конфигурации.
		7002: "Successful repository initialization.",              // Успешная инициализации репозитория.
		7003: "Successful routing initialization.",                 // Успешная инициализации роутинга.
		7004: "Successful connection to the database %v.",          // Успешное подключение к БД %v.
		7005: "Successful write to the database.",                  // Успешная запись в БД.
		7006: "Successful launch of the project for calculation.",  // Успешный запуск проекта на расчет.
		7007: "Successful launch of the Selector.",                 // Успешный запуск Селектора.
		7008: "Successful launch of Trends.",                       // Успешный запуск Трендов.
		7009: "Successful closing of the Selector channel.",        // Успешное закрытие канала Селектора.
		7010: "Successful closing of the Trend channel.",           // Успешное закрытие канала Трендов.
		7011: "Successful setting of trend parameters.",            // Успешная установка параметров трендов.
		7012: "Successfully setting a new value to the parameter.", // Успешная установка нового значения параметру.
		7013: "Successful decoding of the request.",                // Успешное декодирование запроса.
		7014: "Successfully setting a new variable value.",         // Успешная установка нового значения переменной.

		7050: "Successful shutdown of the project.", // Успешная остановка проекта.

		7100: "Successful creation of the NoName file.", // Успешное создание NoName файла.
		7101: "Successful file transfer.",               // Успешная передача файлов.

		7200: "Successful company creation in the database.", // Успешное создание компании в базе данных.

		7300: "Successful sending of an email from password recovery token.", // Успешная отправка письма с токеном восстановления пароля.

		8001: "Starting a Web-Socket connection.", // Старт подключения Web-Socket соединения.
	}

	return mapCodesInfo
}
