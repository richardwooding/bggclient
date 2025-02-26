package model

// XML1Model is an interface that all XML1 models must implement.package model
type XML1Model interface {
	xml1Model()
}

//nolint:golint,unused
func (b *Boardgames) xml1Model() {}

//nolint:golint,unused
func (b *Items) xml1Model() {}

//nolint:golint,unused
func (b *Errors) xml1Model() {}

//nolint:golint,unused
func (b *Geeklist) xml1Model() {}
