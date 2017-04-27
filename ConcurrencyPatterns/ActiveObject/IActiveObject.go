package ActiveObject

//IActiveObject defines the operations for ActiveObject
type IActiveObject interface {
	SetWorkerFunction(workerFunction func(params ...interface{}))

	Run(params ...interface{}) error

	ForceStop()
}
