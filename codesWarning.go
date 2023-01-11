package logger

// WARNING (6000 - 6999)

const (
	Code_ThereIsNoFreeVirtualMachine = 6001 // Нет свободной виртуальной машины.

	Code_TariffResourceExhausted           = 6500 // Исчерпан ресурс тарифа.
	Code_TheMaximumNumberOfRunningProjects = 6501 // Максимальное количество запущенных проектов.

)

func initCodesWarning() map[int]string {
	mapCodesWarning := map[int]string{
		6000: "Reserve",
		6001: "There is no free virtual machine.", // Нет свободной виртуальной машины.

		6500: "Tariff resource exhausted.",              // Исчерпан ресурс тарифа.
		6501: "The maximum number of running projects.", // Максимальное количество запущенных проектов.
	}
	return mapCodesWarning
}
