package model

// XML1Model is an interface that all XML1 models must implement.package model
type XML1Model interface {
	xml1Model()
}

func (b *Boardgames) xml1Model() {}

func (b *Boardgame) xml1Model() {}

func (n *Name) xml1Model() {}

func (p *BoardgamePublisher) xml1Model() {}

func (f *BoardgameFamily) xml1Model() {}

func (c *BoardgameCategory) xml1Model() {}

func (d *BoardgameDesigner) xml1Model() {}

func (a *BoardgameArtist) xml1Model() {}

func (e *BoardgameExpansion) xml1Model() {}

func (p *Poll) xmlModel() {}

func (p *PollSummary) xmlModel() {}

func (p *Result) xmlModel() {}

func (s *SummaryResult) xmlModel() {}
