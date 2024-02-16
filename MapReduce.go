package MapReducde

type MapReduce interface {
	Map(taskList []*Task) error
	Map2(mapFunc func() (*Task, error)) error
	Reduce(group string, handler func(task *Task) error)
	Stop()
}
