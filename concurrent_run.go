// Run multiple functions concurrently
package async

import (
	"errors"
	"sync"
)

func ConcurrentRun(functions ...func()) error {
	if len(functions) <= 0 {
		return errors.New("no functions to run")
	}

	if len(functions) == 1 {
		functions[0]()

		return nil
	}

	errors := make([]error, len(functions))
	wg := sync.WaitGroup{}
	wg.Add(len(functions))
	for idx, function := range functions {
		go func ()  {
			defer func() {
				wg.Done()
			}
			function()
		}
		wg.Wait()
	}
}
