package engine

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/kaluginivann/Aegis/internal/aggregator"
	"github.com/kaluginivann/Aegis/internal/configs"
	"github.com/kaluginivann/Aegis/internal/detector"
	"github.com/kaluginivann/Aegis/internal/files"
	"github.com/kaluginivann/Aegis/internal/workers"
)

type Engine struct {
	conf       *configs.Config
	Detector   detector.Interface
	Aggregator aggregator.Interface
}

func NewEngine(conf *configs.Config) *Engine {
	return &Engine{
		conf:       conf,
		Detector:   detector.NewDetector(),
		Aggregator: aggregator.NewAggregator(),
	}
}

func (e *Engine) Run() {
	FileInfo, err := files.CheckExistsFile(e.conf)
	if err != nil {
		panic(err)
	}

	WorkerPool := workers.NewWrokerPool(runtime.NumCPU(), e.conf.Logger)
	BufferSize := e.GetBufferSize(FileInfo)

	file, err := os.Open(e.conf.FilePath)
	if err != nil {
		e.conf.Logger.Error("Error open file", "error", err)
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	WorkerPool.Start()
	e.Aggregator.Start()

	e.ReadFile(buffer, file, WorkerPool)

	WorkerPool.Wait()
	e.Aggregator.Stop()
	WorkerPool.Stop()

	fmt.Println(e.Aggregator.Result())
}

func (e *Engine) ReadFile(buffer []byte, file *os.File, WorkerPool *workers.WorkerPool) {
	var offset int64

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			chunk := make([]byte, n)
			copy(chunk, buffer[:n])

			currentOffset := offset

			WorkerPool.Add(func() {
				result := e.Detector.Scan(chunk)
				if len(result) > 0 {
					e.Aggregator.Add(aggregator.ScanResult{
						Matches: result,
						Offset:  currentOffset,
					})
				}
			})
		}
		offset += int64(n)

		if err == io.EOF {
			return
		} else if err != nil {
			e.conf.Logger.Error("Error read file", "error", err)
			panic(err)
		}
	}
}

func (e *Engine) GetBufferSize(FileInfo os.FileInfo) int {
	size := int(FileInfo.Size())
	switch {
	case size < SmallFileThreshold:
		return SmallChunkSize
	case size < MediumFileThreshold:
		return MediumChunkSize
	default:
		return LargeChunkSize
	}
}
