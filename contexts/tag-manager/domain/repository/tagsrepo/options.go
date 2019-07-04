package tagsrepo

type Options []Option

func (opts Options) Stmp() (stmp *Stmp, err error) {
	stmp = new(Stmp)
	for _, opt := range opts {
		if err = opt(stmp); err != nil {
			return nil, err
		}
	}
	return
}
