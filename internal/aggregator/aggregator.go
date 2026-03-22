package aggregator

import "sync"

type Interface interface {
	Result() map[string]int
	Add(res ScanResult)
	Stop()
	Start()
}

type Aggregator struct {
	results     map[string]int
	resultsChan chan ScanResult
	wg          sync.WaitGroup
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		results:     make(map[string]int),
		resultsChan: make(chan ScanResult, 100),
	}
}

func (a *Aggregator) Start() {
	a.wg.Go(func() {
		for res := range a.resultsChan {
			for _, match := range res.Matches {
				a.results[match]++
			}
		}
	})
}

func (a *Aggregator) Result() map[string]int {
	return a.results
}

func (a *Aggregator) Add(res ScanResult) {
	a.resultsChan <- res
}

func (a *Aggregator) Stop() {
	close(a.resultsChan)
	a.wg.Wait()
}
