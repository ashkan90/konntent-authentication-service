package utils

const (
	DefaultChunkSize = 250
)

func Chunk[T any](data []T, size int) ([][]T, int) {
	ln := len(data)

	if ln == 0 {
		return nil, 0
	}

	chunks := make([][]T, (ln+size-1)/size)
	prev := 0
	i := 0
	till := ln - size

	for prev < till {
		next := prev + size
		chunks[i] = data[prev:next]
		prev = next
		i++
	}
	chunks[i] = data[prev:]

	return chunks, ln
}
