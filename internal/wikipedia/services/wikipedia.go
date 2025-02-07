package services

func (w *wikipediaService) RunService() {
	for change := range w.ReadStream() {
		msg, err := w.ConstructMessage(&change)
		if err != nil {
			w.logger.Errorf("Error constructing message: %v", err)
			continue
		}

		if msg != "" {
			if err = w.SendNotification(msg); err != nil {
				w.logger.Errorf("Error sending notification: %v", err)
			}
		} else {
			w.logger.Warn("No message was sent")
		}
	}
}
