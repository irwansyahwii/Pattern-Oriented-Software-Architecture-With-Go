package ActiveObject

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestActiveObject(t *testing.T) {
	var activeObject IActiveObject

	var wait sync.WaitGroup
	wait.Add(1)

	activeObject = NewActiveObjectWithInterval(time.Millisecond * 50)

	counter := 0
	activeObject.SetWorkerFunction(func(param interface{}) {
		counter++

		if counter > 3 {
			wait.Done()
		}
	})

	activeObject.Run(10)

	wait.Wait()

	activeObject.ForceStop()

	time.Sleep(time.Millisecond * 1000)

	assert.Equal(t, counter, 4, "counter is wrong")
}
