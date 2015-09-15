//Package webclients contains a
package webclients

type Request struct {
	Headers []string
	Body    string
}

type Response struct {
	Headers    []string
	Body       string
	ReturnCode int
}

type WorkerArgs struct {
	Done     chan<- struct{}
	Request  Request
	Response Response
}

func New() *WorkerArgs {
	wa := new(WorkerArgs)

	wa.Done = make(chan<- struct{}, 1)
	wa.Request.Headers = make([]string, 16)
	wa.Response.Headers = make([]string, 16)

	return wa
}

func Worker(wArgs WorkerArgs) {

}
