package logger

// WARNING (6000 - 6999)

const (
	Code_ThereIsNoFreeVirtualMachine = 6001 // Нет свободной виртуальной машины.
	Code_AutomaticStop               = 6002 // Проект остановился автоматически.

	Code_TariffResourceExhausted           = 6500 // Исчерпан ресурс тарифа.
	Code_TheMaximumNumberOfRunningProjects = 6501 // Максимальное количество запущенных проектов.
	Code_MDCoreInternalErrors              = 6502 // Внутренние ошибки в мдкоре (вывод логов)
	Code_EndOfFile                         = 6503 // EOF от мдкора.

)

func initCodesWarning() map[int]string {
	mapCodesWarning := map[int]string{
		6000: "Reserve",
		6001: "There is no free virtual machine.", // Нет свободной виртуальной машины.
		6002: "Automatic stop",                    // Проект остановился автоматически

		6500: "Tariff resource exhausted.",              // Исчерпан ресурс тарифа.
		6501: "The maximum number of running projects.", // Максимальное количество запущенных проектов.
		6502: "MDCore internal errors.",                 // Внутренние ошибки в мдкоре (вывод логов).
		6503: "End of file.",                            // EOF от мдкора.
	}
	return mapCodesWarning
}
