package worker

import (
	"sync"

	"gopkg.in/eapache/channels.v1"
)

type Job interface {
	Process()
}

type WorkerPool struct {
	queueSize   int
	workerNum   int
	jobChannels *channels.RingChannel
	quit        chan bool
	wg          *sync.WaitGroup
}

func NewWorker(poolNum int, queueSize int) *WorkerPool {
	worker := WorkerPool{
		queueSize:   queueSize,
		workerNum:   poolNum,
		jobChannels: channels.NewRingChannel(channels.BufferCap(queueSize)),
		quit:        make(chan bool, poolNum),
		wg:          new(sync.WaitGroup),
	}

	return &worker
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.workerNum; i++ {
		go func() {
			w.wg.Add(1)
			for {
				select {
				case job := <-w.jobChannels.Out():
					job.(Job).Process()
				case <-w.quit:
					w.wg.Done()
					return
				}
			}
		}()
	}
}

func (w *WorkerPool) Stop() {
	for i := 0; i < w.workerNum; i++ {
		w.quit <- true
	}

	w.wg.Wait()
	w.jobChannels.Close()
	close(w.quit)
}

func (w *WorkerPool) GetQueueSize() int {
	return w.jobChannels.Len()
}

func (w *WorkerPool) AddJob(j Job) {
	w.jobChannels.In() <- j
}
