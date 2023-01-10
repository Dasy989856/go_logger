# GoLogger

## Процесс логирования имеет несколько уровней:
- DEBUG - фиксируются события с уровнями: DEBUG, INFO, WARNING, ERROR, CRITICAL;
- INFO - фиксируются события с уровнями: INFO, WARNING, ERROR, CRITICAL;
- WARNING - фиксируются события с уровнями: WARNING, ERROR, CRITICAL;
- ERROR - фиксируются события с уровнями: ERROR и CRITICAL;
- CRITICAL - фиксируются события с уровнем CRITICAL.

## Настройка уровня логирования:
- В файле конфигурации /confings/config.yml поле `logger.level: debug`
- Параметр выставляется в соответствии с необходимым уровнем логирования: debug, info, warning, error, critical.

## Описание уровней сообщений о событиях, попадающих в лог файлы:
- DEBUG - подробное и детальное логирование всей системной информации для последующего использования в отладке;
- INFO - подтверждение, информация о событиях, не приводящих к ошибкам в работе модулей;
- WARNING - информация о событиях, которые могут привести к ошибкам в работе модулей;
- ERROR - информация об ошибках, возникших в работе модулей;
- CRITICAL - информация о критических ошибках, возникших в работе модулей.

## Примеры использования go_logger:
- **Создание логера на корневом уровне с пустой конфигурацией:** `logger := go_logger.NewLogger(nil)`

- **Создание конфигурации и его установка:**
```
configLogger := go_logger.Config{
		LogLevel:      viper.GetString("logger.level"),
		LogServiceAPI: logServiceApi,
		NameService:   models.NameService,
	}
logger.SetConfig(&configLogger)
```

- **Создание дочернего logger. Он берет свойства конфигурации от родительского logger (Используется для hanlder):**
```
    handlerLogger := logger.ChildLogger()
```

- **Инициализируем родительское событие:** `pEvent := hLogger.InitParentEvent("NamePackage", "NameFunction")`
```
NamePackage - имя пакета.
NameFunction - имя функции.
```

- **Создаем события нужного типа от родительского:** `event := pEvent.Error(go_logger.Code_ErrorDecodingJson)`
- **(Опционально) Добавление оригинальной ошибки и http статуса.**
- **(Опционально) Добавление контекста события.**
- **Example:**
```
if err != nil {
    event := pEvent.Error(go_logger.Code_ErrorDecodingJson) // Создание события.
    event.AddError(err) // Добавление в событие оригинальной ошибки.
    event.SetStatusHTTP(http.StatusBadRequest) // Статус код который будет установлен в ответе на http запрос REST API.
    event.AddContext(map[string]interface{}{"url": url, "password": pass}
    return
}
```
- **Вывод событий:**
```
logger.SendToLogService() - данный метод делает попытку отправки логов в сервис логирования. При неудачи выводит логи в StdOut.
logger.Print() - вывод логов в StdOut.

logger.Code - логер содержит коды (либо дает возможность дать код определеному событию. Каждый код события соответсвует текстовому сообщению.
params - это параметры для сообщения события. Пример: Ошибка подключения к БД %v. - params1 подставиться вместо %v.
```
