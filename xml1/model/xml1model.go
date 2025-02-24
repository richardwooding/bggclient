package model

// XML1Model is an interface that all XML1 models must implement.package model
type XML1Model interface {
	xml1Model()
}

//nolint:golint,unused
func (b *Boardgames) xml1Model() {}

//nolint:golint,unused
func (b *Boardgame) xml1Model() {}

//nolint:golint,unused
func (n *Name) xml1Model() {}

//nolint:golint,unused
func (p *BoardgamePublisher) xml1Model() {}

//nolint:golint,unused
func (f *BoardgameFamily) xml1Model() {}

//nolint:golint,unused
func (c *BoardgameCategory) xml1Model() {}

//nolint:golint,unused
func (d *BoardgameDesigner) xml1Model() {}

//nolint:golint,unused
func (a *BoardgameArtist) xml1Model() {}

//nolint:golint,unused
func (e *BoardgameExpansion) xml1Model() {}

//nolint:golint,unused
func (p *Poll) xmlModel() {}

//nolint:golint,unused
func (p *PollSummary) xmlModel() {}

//nolint:golint,unused
func (p *Result) xmlModel() {}

//nolint:golint,unused
func (s *SummaryResult) xmlModel() {}
