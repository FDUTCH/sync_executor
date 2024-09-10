package executor

import "sync"

// Runnable ...
type Runnable interface {
	Run()
}

// Executor allows you to wrap your code safely in mutex so you will never experience deadlock
type Executor struct {
	mu sync.Mutex
}

// Run runs Runnable synchronously
func (e *Executor) Run(runnable Runnable) {
	e.Exec(runnable.Run)
}

func NewExecutor() *Executor {
	return &Executor{}
}

// Exec runs passed function synchronously
func (e *Executor) Exec(fn func()) {
	e.mu.Lock()
	defer e.mu.Unlock()
	fn()
}

var executor = NewExecutor()

// Exec runs passed function synchronously
func Exec(fn func()) {
	executor.Exec(fn)
}

// Run runs Runnable synchronously
func Run(runnable Runnable) {
	executor.Exec(runnable.Run)
}
