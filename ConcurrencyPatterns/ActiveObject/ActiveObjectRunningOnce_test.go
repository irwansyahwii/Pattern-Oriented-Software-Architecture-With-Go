package ActiveObject

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestActiveObjectRunningOnce(t *testing.T) {
	var activeObject IActiveObject

	var wait sync.WaitGroup
	wait.Add(1)

	counter := 0

	activeObject = NewActiveObjectRunningOnce(time.Millisecond*50, func(param interface{}) {
		assert.Equal(t, param, 10, "param is incorrect")

		counter++

		wait.Done()
	})

	activeObject.Run(10)

	wait.Wait()

	activeObject.ForceStop()

	time.Sleep(time.Millisecond * 1000)

	assert.Equal(t, counter, 1, "counter is wrong")
}
