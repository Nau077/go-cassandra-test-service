package repos

type enterDTO struct {
	userName string
	password string
}

type authSign interface {
	signIn(enterDTO) error
	signUp() error
}
