package domain

type Manifest struct {
	ClassName string
	Content   string
}

func (manifest Manifest) GetClassName() string {
	return manifest.ClassName
}

func (manifest Manifest) GetContent() string {
	return manifest.Content
}
