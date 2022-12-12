package logger

const (
	Code_NoCode = 0
)

// Соотношение кода события к сообщению.
var MapCodes = initMapCodes()

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