package services

import "time"

func (w *wikipediaService) StartService() {
	if w.running {
		w.logger.Warn("Service is already running")
		return
	}

	w.running = true
}

func (w *wikipediaService) StopService() {
	if !w.running {
		w.logger.Warn("Service is already not running")
		return
	}

	w.running = false
}

func (w *wikipediaService) RestartService() {
	w.StopService()
	time.Sleep(500 * time.Millisecond)
	w.StartService()
}
