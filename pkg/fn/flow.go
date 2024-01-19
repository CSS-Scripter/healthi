package fn

type Flow struct {
	Slice[func() error]
}

func (f Flow) Run() error {
	var err error
	f.ForEach(func(fn func() error) bool {
		err = fn()
		return err == nil
	})
	return err
}
