package repos

type Container struct {
	PostRepo *PostRepo
}

func NewContainer() *Container {
	p := &PostRepo{}

	return &Container{
		PostRepo: p,
	}
}
