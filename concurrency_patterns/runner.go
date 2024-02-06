package main

// type Runner struct {
// 	interrupt chan os.Signal
// 	complete  chan error
// 	timeout   <-chan time.Time
// 	tasks     []func(int)
// }

// var ErrTimeout = errors.New("Receieved timeout")
// var ErrInterrupt = errors.New("receieved interrupt")

// func New(d time.Duration) *Runner {
// 	return &Runner{
// 		interrupt: make(chan os.Signal, 1),
// 		complete:  make(chan error),
// 		timeout:   time.After(d),
// 	}
// }

// func (r *Runner) Add(tasks ...func(int)) {
// 	r.tasks = append(r.tasks, tasks...)
// }

// func (r *Runner) Start() error {
// 	signal.Notify(r.interrupt, os.Interrupt)

// }

// func main() {

// }
