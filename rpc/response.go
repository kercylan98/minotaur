package rpc

type Response interface {
	ReadTo(dst any) error

	MustReadTo(dst any)
}

func NewResponse(reader contextReader) Response {
	return &response{
		reader: reader,
	}
}

type response struct {
	reader contextReader
}

func (r *response) ReadTo(dst any) error {
	return r.reader(dst)
}

func (r *response) MustReadTo(dst any) {
	if err := r.reader(dst); err != nil {
		panic(err)
	}
}
