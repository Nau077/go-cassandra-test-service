package repos

type Container struct {
	PostRepo *PostRepo
}

func NewContainer(p *PostRepo) *Container {
	return &Container{
		PostRepo: p,
	}
}
