package engine

const (
	KB int = 1024
	MB int = 1024 * KB

	SmallFileThreshold  int = 1 * MB
	MediumFileThreshold int = 100 * MB

	SmallChunkSize  int = 4 * KB
	MediumChunkSize int = 64 * KB
	LargeChunkSize  int = 1 * MB
)
