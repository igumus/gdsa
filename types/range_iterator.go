package types

type RangeOption func(*rangeOptions)

type rangeOptions struct {
	start  int
	end    int
	step   func(int) int
	finite bool
}

func applyRangeOptions(opts ...RangeOption) *rangeOptions {
	ret := &rangeOptions{
		start:  0,
		end:    100,
		step:   func(i int) int { return i + 1 },
		finite: false,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func WithStart(s int) RangeOption {
	return func(ro *rangeOptions) {
		ro.start = s
	}
}

func WithEnd(e int) RangeOption {
	return func(ro *rangeOptions) {
		ro.end = e
	}
}

func WithStepFunction(f func(int) int) RangeOption {
	return func(ro *rangeOptions) {
		ro.step = f
	}
}

type rangeIter struct {
	start   int
	end     int
	current int
	step    func(int) int
	finite  bool
}

func newRange(cfg *rangeOptions) *rangeIter {
	return &rangeIter{
		start:   cfg.start,
		end:     cfg.end,
		current: cfg.start,
		step:    cfg.step,
		finite:  cfg.finite,
	}
}

func NewFiniteRange(opts ...RangeOption) Iterator[int] {
	cfg := applyRangeOptions(opts...)
	cfg.finite = true
	return newRange(cfg)
}

func NewInfiniteRange(opts ...RangeOption) Iterator[int] {
	cfg := applyRangeOptions(opts...)
	cfg.finite = false
	return newRange(cfg)
}

func (r *rangeIter) HasNext() bool {
	if !r.finite {
		return true
	}

	if r.finite && r.current < r.end {
		return true
	}
	return false
}

func (r *rangeIter) Next() int {
	temp := r.current
	r.current = r.step(r.current)
	return temp
}
