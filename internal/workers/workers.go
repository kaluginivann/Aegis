package workers

import (
	"sync"

	"github.com/kaluginivann/Aegis/internal/logger"
)

func Worker(wg *sync.WaitGroup, jobChannel <-chan Job, logger logger.Interface, i int) {
	logger.Info("Start worker", "Number", i)
LOOP:
	for {
		select {
		case job, ok := <-jobChannel:
			if !ok {
				break LOOP
			}
			job()
			wg.Done()
		}
	}
}
