package mandelbrot

type Task struct {
	startLineIndex, numberOfLines int
}

func NewTask(startLineIndex, numberOfLines int) *Task {
	return &Task{startLineIndex:startLineIndex, numberOfLines:numberOfLines}
}