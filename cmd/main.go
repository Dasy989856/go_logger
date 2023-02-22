package main

import (
	"fmt"

	go_logger "github.com/dasy989856/go_logger"
)

func main() {
	logger := go_logger.NewLogger(nil)
	defer func() {
		if err := logger.SendToLogService(); err != nil {
			logger.Print()
		}
		fmt.Println(logger.GetStatusHTTP())
		fmt.Println("CounterErr: ", logger.GetNumberOfErrors())
	}()
	pEvent := logger.InitParentEvent("main", "main")


	if err := newFunc(logger); err != nil {
		pEvent.Debug(go_logger.Code_NoCode).SetStatusHTTP(109999)
		return
	}


	event := pEvent.Info(go_logger.Code_SuccessfulRepositoryInitialization)
	event.AddContext(map[string]interface{}{"context": "test context"}).Print()
}

func newFunc(logger go_logger.Logger) error {
	pEvent := logger.InitParentEvent("main", "newFunc")

	err := fmt.Errorf("error newFunc")
	if err != nil {
		pEvent.Error(go_logger.Code_NoCode).AddError(err).SetStatusHTTP(999)
		return err
	}
	return nil
}
