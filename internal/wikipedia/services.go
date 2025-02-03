package wikipedia

type WikipediaService interface {
	StreamReaderService
}

type StreamReaderService interface {
	ReadStream()
}
