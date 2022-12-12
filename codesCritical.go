package logger

// CRITICAL (1-999)

const (
	// 1-49 (Server)

	Code_ErrorStarting__NameService__Service = 1 // Ошибка запуска %v сервиса.
	Code_ErrorInitializingConfiguration      = 2 // Ошибка инициализации конфигурации.
	Code_ErrorInitializingRepository         = 3 // Ошибка инициализации репозитория.

	// 50-150 (DataBase)
	Code_ErrorConnectingToDatabase__NameDB__          = 20 // Ошибка подключения к БД %v.
	Code_ConnectionToDataBase__NameDB__IsNotSupported = 21 // Подключение к БД %v не поддерживается.
)

func initCodesCritical() map[int]string {
	mapCodesCritical := map[int]string{
		// 1-49 (Server)
		1: "Error starting %v service.",          // Ошибка запуска %v сервиса.
		2: "Configuration initialization error.", // Ошибка инициализации конфигурации.
		3: "Error initializing repository.",      // Ошибка инициализации репозитория.

		// 50-150 (DataBase)
		20: "Error connecting to database %v.",            // Ошибка подключения к БД %v.
		21: "Connection to database %v is not supported.", // Подключение к БД %v не поддерживается.
	}

	return mapCodesCritical
}
