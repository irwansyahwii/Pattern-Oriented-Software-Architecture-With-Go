package ActiveObject

import (
	"errors"
	"time"
)

//ActiveObjectRunningOnce implements IActiveObject with behavior running in a specified interval
type ActiveObjectRunningOnce struct {
	workerFunction func(params interface{})
	duration       time.Duration
	isStopped      bool
}

func NewActiveObjectRunningOnce(duration time.Duration, workerFunction func(param interface{})) *ActiveObjectRunningOnce {

	return &ActiveObjectRunningOnce{duration: duration, workerFunction: workerFunction}

}
func (activeObject *ActiveObjectRunningOnce) SetWorkerFunction(workerFunction func(param interface{})) {

	activeObject.workerFunction = workerFunction

}

func (activeObject *ActiveObjectRunningOnce) executeWorkerFunctionConcurrently(param interface{}) error {

	if activeObject.workerFunction == nil {
		return errors.New("Please set the worker function first")
	}
	go activeObject.workerFunction(param)

	return nil

}

func (activeObject *ActiveObjectRunningOnce) Run(param interface{}) error {
	activeObject.isStopped = false
	return activeObject.executeWorkerFunctionConcurrently(param)

}

func (activeObject *ActiveObjectRunningOnce) ForceStop() {
	activeObject.isStopped = true
}
