package services

import "time"

func (w *wikipediaService) Start() {
	w.running = true
}

func (w *wikipediaService) Stop() {
	w.running = false
}

func (w *wikipediaService) Restart() {
	w.Stop()
	time.Sleep(500 * time.Millisecond)
	w.Start()
}
