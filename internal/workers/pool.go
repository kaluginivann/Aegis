package workers

import (
	"sync"

	"github.com/kaluginivann/Aegis/internal/logger"
)

type Job func()

type WorkerPool struct {
	WorkersCount int
	Logger       logger.Interface
	jobs         chan Job
	wg           sync.WaitGroup
}

func NewWrokerPool(WorkersCount int, logger logger.Interface) *WorkerPool {
	return &WorkerPool{
		WorkersCount: WorkersCount,
		Logger:       logger,
		jobs:         make(chan Job),
	}
}

func (w *WorkerPool) Start() {
	for range w.WorkersCount {
		go Worker(&w.wg, w.jobs, w.Logger)
	}
}

func (w *WorkerPool) Wait() {
	w.wg.Wait()
}

func (w *WorkerPool) Stop() {
	close(w.jobs)
}

func (w *WorkerPool) Add(job Job) {
	w.wg.Add(1)
	w.jobs <- job
}
