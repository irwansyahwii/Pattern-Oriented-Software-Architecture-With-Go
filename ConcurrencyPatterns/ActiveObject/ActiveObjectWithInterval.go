package ActiveObject

import (
	"errors"
	"time"
)

//ActiveObjectWithInterval implements IActiveObject with behavior running in a specified interval
type ActiveObjectWithInterval struct {
	ticker      *time.Ticker
	doneChannel chan bool

	ActiveObjectRunningOnce
}

func NewActiveObjectWithInterval(duration time.Duration, workerFunction func(param interface{})) *ActiveObjectWithInterval {

	activeObject := &ActiveObjectWithInterval{}

	activeObject.duration = duration
	activeObject.workerFunction = workerFunction
	activeObject.doneChannel = make(chan bool)

	return activeObject
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
