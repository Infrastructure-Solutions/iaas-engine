package domain

type Hiera struct {
	Classes []string
}

func (hiera Hiera) GetClasses() []string {
	return hiera.Classes
}
