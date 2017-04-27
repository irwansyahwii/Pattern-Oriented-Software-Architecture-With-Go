package ActiveObject

//IActiveObject defines the operations for ActiveObject
type IActiveObject interface {
	SetWorkerFunction(workerFunction func(param interface{}))

	Run(param interface{}) error

	ForceStop()
}
