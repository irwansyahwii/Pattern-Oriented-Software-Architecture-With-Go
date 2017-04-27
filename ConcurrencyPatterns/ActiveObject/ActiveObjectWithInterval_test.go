package ActiveObject

import (
	"sync"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestActiveObject(t *testing.T) {
	var activeObject IActiveObject

	var wait sync.WaitGroup
	wait.Add(1)

	activeObject = NewActiveObjectWithInterval(time.Millisecond * 500)

	counter := 0
	activeObject.SetWorkerFunction(func(params ...interface{}) {
		fmt.Println("Counter")
		counter++

		if counter > 3 {
			wait.Done()
		}
	})

	activeObject.Run()

	wait.Wait()

	activeObject.ForceStop()

	assert.Equal(t, counter, 4, "counter is wrong")
}
