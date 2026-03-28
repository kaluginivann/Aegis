package workers

import (
	"sync"

	"github.com/kaluginivann/Aegis/internal/logger"
)

func Worker(wg *sync.WaitGroup, jobChannel <-chan Job, logger logger.Interface) {
	logger.Info("Start worker")
	for {
		select {
		case job, ok := <-jobChannel:
			if !ok {
				goto END
			}
			job()
			wg.Done()
		}
	}
END:
	logger.Info("Job is done!")
}
