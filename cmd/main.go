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
		fmt.Println("CounterErr: ", logger.GetNumberOfErrors())
	}()
	pEvent := logger.InitParentEvent("main", "main")

	var err error = nil
	if err != nil {
		event := pEvent.Critical(go_logger.Code_ErrorInitializingRepository)
		event.AddError(err).Print()
		return
	}

	if err := newFunc(logger); err != nil {
		pEvent.Debug(go_logger.Code_NoCode)
		return
	}

	event := pEvent.Info(go_logger.Code_SuccessfulRepositoryInitialization)
	event.AddContext(map[string]interface{}{"context": "test context"}).Print()
}

func newFunc(logger go_logger.Logger) error {
	pEvent := logger.InitParentEvent("main", "newFunc")

	err := fmt.Errorf("error newFunc")
	if err != nil {
		pEvent.Error(go_logger.Code_NoCode).AddError(err)
		return err
	}
	return nil
}
