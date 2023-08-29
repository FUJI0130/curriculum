package userdm

type Careers struct {
	from   CareersFrom
	to     CareersTo
	detail string
}

func NewCareers(from int, to int, detail string) (*Careers, error) {
	cf, err := NewCareersFrom(from)
	if err != nil {
		return nil, err
	}

	ct, err := NewCareersTo(to)
	if err != nil {
		return nil, err
	}

	return &Careers{
		from:   *cf,
		to:     *ct,
		detail: detail,
	}, nil
}

func (c *Careers) From() CareersFrom {
	return c.from
}

func (c *Careers) To() CareersTo {
	return c.to
}

func (c *Careers) Detail() string {
	return c.detail
}
