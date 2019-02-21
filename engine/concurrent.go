package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	//WorkerReady(chan Request)
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.Run()

	//e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
