package orchestrator

type Msg interface{}

// Runner is any type that will run in a goroutine
// It will get its information via the channel returned by
// GetChanIn() and send the run informations back via the return
// the Process() function
type Runner interface {
	GetChanIn() <-chan Msg
	Process(done <-chan struct{}) <-chan Msg
	Execute()
}

// Runs a set of runner (one per goroutine) and make them communicate with each other
func Run(runner ...Runner) {

}
