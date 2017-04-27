package ActiveObject

import (
	"errors"
	"time"
)

//ActiveObjectWithInterval implements IActiveObject with behavior running in a specified interval
type ActiveObjectWithInterval struct {
	workerFunction func(params interface{})
	ticker         *time.Ticker
	duration       time.Duration
	doneChannel    chan bool
}

func NewActiveObjectWithInterval(duration time.Duration) *ActiveObjectWithInterval {

	return &ActiveObjectWithInterval{duration: duration, doneChannel: make(chan bool)}
}

func (activeObject *ActiveObjectWithInterval) SetWorkerFunction(workerFunction func(param interface{})) {
	activeObject.workerFunction = workerFunction
}

func (activeObject *ActiveObjectWithInterval) Run(param interface{}) error {
	if activeObject.ticker != nil {
		return errors.New("Already running")
	}

	activeObject.ticker = time.NewTicker(activeObject.duration)

	go func() {
		for {
			select {
			case <-activeObject.ticker.C:
				activeObject.workerFunction(param)

			case <-activeObject.doneChannel:
				activeObject.ticker.Stop()
				return
			}
		}
	}()

	return nil
}

func (activeObject *ActiveObjectWithInterval) ForceStop() {
	activeObject.doneChannel <- true
}
