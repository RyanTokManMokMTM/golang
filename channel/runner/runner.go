//Simple concurrent code
package runner
import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//define custom error type
var (
	ErrTimeOut = errors.New("can't finished the task within given time")
	ErrOSInterrupt = errors.New("received interrupt")
)

type Runner struct {
	interrupt chan os.Signal //a channel for os single
	complete chan error //a channel for error , completed:nil else err msg
	timeout <- chan time.Time //channal for received a time
	tasks []func(int) //a list of task with id
}

func New(t time.Duration) *Runner{
	return &Runner{
		interrupt: make(chan os.Signal,1),
		complete: make(chan error),
		timeout: time.After(t), //received a time signal after the time
		tasks: make([]func(int),0),//define as a slice
	}
}

func (r *Runner) AddTask(task ...func(int)){
	r.tasks = append(r.tasks,task...) //add our multiple task to tasks slice
}

func (r *Runner) run() error{
	//run each task
	for id, task := range r.tasks{
		//here ,will check the channel first.Timeout ?
		//here listening to each case of channel signal
		select {
		case <- r.interrupt:
			//it received some interrupt signal, stop Listening and return the error
			signal.Stop(r.interrupt) //stop relaying incoming signals to c.
			return ErrOSInterrupt
		default:
			task(id) //run the task if there are not any interrupt,	run our task with id
		}
	}
	return nil
}

func (r *Runner) Start() error{
	//if os interrupted ,send it to runner interrupt channel
	signal.Notify(r.interrupt,os.Interrupt) //listening to os,if it has some interrupt

	//run all the task with goroutine
	go func(){
		r.complete <- r.run() // received a signal if it is
	}()

	select {
	case <-r.timeout:
		return ErrTimeOut //time is out and no tasks have been completed
	case <-r.complete:
		return nil //mission completed ,not error
	}
}

