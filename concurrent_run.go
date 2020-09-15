// Run multiple functions concurrently
package async

import (
	"errors"
	"runtime/debug"
	"sync"
)

func ConcurrentRun(functions ...func()) error {
	// There are no functions to run
	if len(functions) <= 0 {
		return errors.New("no functions to run")
	}

	// Just 1 function needs to run, so run as a plain function
	if len(functions) == 1 {
		functions[0]()

		return nil
	}

	var runtimeError error
	runtimeErrors := make([]error, len(functions))
	wg := sync.WaitGroup{}
	wg.Add(len(functions))
	for idx, function := range functions {
		go func(index int, function func()) {
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					runtimeErrors[index] = errors.New(string(debug.Stack()))
					// runtimeError = errors.Unwrap(runtimeErrors[index])
				}
			}()
			function()
		}(idx, function)
	}
	wg.Wait()

	// fmt.Println(runtimeErrors)
	return runtimeError
}
