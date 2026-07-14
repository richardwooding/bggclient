// Package bggopts defines option structs shared by the CLI (kong tags) and
// the MCP server (json/jsonschema tags), converting them to the functional
// options of the xml1 package.
package bggopts

import (
	"fmt"
	"time"

	"github.com/richardwooding/bggclient/xml1"
)

// BoardgameFlags are the optional parameters of the boardgame endpoint.
type BoardgameFlags struct {
	Comments   *bool   `help:"Include user comments." json:"comments,omitempty" jsonschema:"include user comments"`
	Stats      *bool   `help:"Include rating statistics." json:"stats,omitempty" jsonschema:"include rating statistics"`
	Historical *bool   `help:"Include historical data." json:"historical,omitempty" jsonschema:"include historical rank and rating data"`
	From       *string `help:"Historical start date (YYYY-MM-DD)." json:"from,omitempty" jsonschema:"historical data start date in YYYY-MM-DD format"`
	To         *string `help:"Historical end date (YYYY-MM-DD)." json:"to,omitempty" jsonschema:"historical data end date in YYYY-MM-DD format"`
}

func (f BoardgameFlags) Options() ([]xml1.BoardgameOption, error) {
	var opts []xml1.BoardgameOption
	if f.Comments != nil {
		opts = append(opts, xml1.Comments(*f.Comments))
	}
	if f.Stats != nil {
		opts = append(opts, xml1.Stats(*f.Stats))
	}
	if f.Historical != nil {
		opts = append(opts, xml1.Historical(*f.Historical))
	}
	if f.From != nil {
		t, err := time.Parse(time.DateOnly, *f.From)
		if err != nil {
			return nil, fmt.Errorf("invalid from date %q: expected YYYY-MM-DD", *f.From)
		}
		opts = append(opts, xml1.From(&t))
	}
	if f.To != nil {
		t, err := time.Parse(time.DateOnly, *f.To)
		if err != nil {
			return nil, fmt.Errorf("invalid to date %q: expected YYYY-MM-DD", *f.To)
		}
		opts = append(opts, xml1.To(&t))
	}
	return opts, nil
}

// CollectionFlags are the optional filters of the collection endpoint.
type CollectionFlags struct {
	Own              *bool `help:"Filter on owned games." json:"own,omitempty" jsonschema:"filter on owned games"`
	Rated            *bool `help:"Filter on rated games." json:"rated,omitempty" jsonschema:"filter on games the user has rated"`
	Comment          *bool `help:"Filter on commented games." json:"comment,omitempty" jsonschema:"filter on games the user has commented on"`
	Trade            *bool `help:"Filter on games for trade." json:"trade,omitempty" jsonschema:"filter on games offered for trade"`
	Want             *bool `help:"Filter on wanted games." json:"want,omitempty" jsonschema:"filter on games wanted in trade"`
	Wishlist         *bool `help:"Filter on wishlisted games." json:"wishlist,omitempty" jsonschema:"filter on wishlisted games"`
	WantToPlay       *bool `help:"Filter on games the user wants to play." json:"want_to_play,omitempty" jsonschema:"filter on games the user wants to play"`
	WantToBuy        *bool `help:"Filter on games the user wants to buy." json:"want_to_buy,omitempty" jsonschema:"filter on games the user wants to buy"`
	PrevOwned        *bool `help:"Filter on previously owned games." json:"prev_owned,omitempty" jsonschema:"filter on previously owned games"`
	PreOrdered       *bool `help:"Filter on pre-ordered games." json:"pre_ordered,omitempty" jsonschema:"filter on pre-ordered games"`
	HasParts         *bool `help:"Filter on games with spare parts." json:"has_parts,omitempty" jsonschema:"filter on games for which the user has spare parts"`
	WantParts        *bool `help:"Filter on games needing parts." json:"want_parts,omitempty" jsonschema:"filter on games for which the user wants parts"`
	WishlistPriority *int  `help:"Filter wishlist priority (1-5)." json:"wishlist_priority,omitempty" jsonschema:"filter on wishlist priority, 1 (must have) to 5 (don't buy)"`
	MinRating        *int  `help:"Minimum personal rating (1-10)." json:"min_rating,omitempty" jsonschema:"minimum personal rating, 1 to 10"`
	MaxRating        *int  `help:"Maximum personal rating (1-10)." json:"max_rating,omitempty" jsonschema:"maximum personal rating, 1 to 10"`
	MinBGGRating     *int  `help:"Minimum BGG rating (1-10)." json:"min_bgg_rating,omitempty" jsonschema:"minimum BGG community rating, 1 to 10"`
	MaxBGGRating     *int  `help:"Maximum BGG rating (1-10)." json:"max_bgg_rating,omitempty" jsonschema:"maximum BGG community rating, 1 to 10"`
	MinPlays         *int  `help:"Minimum recorded plays." json:"min_plays,omitempty" jsonschema:"minimum number of recorded plays"`
	MaxPlays         *int  `help:"Maximum recorded plays." json:"max_plays,omitempty" jsonschema:"maximum number of recorded plays"`
}

func (f CollectionFlags) Options() []xml1.CollectionOption {
	var opts []xml1.CollectionOption
	boolOpts := []struct {
		value  *bool
		option func(bool) xml1.CollectionOption
	}{
		{f.Own, xml1.Own},
		{f.Rated, xml1.Rated},
		{f.Comment, xml1.Comment},
		{f.Trade, xml1.Trade},
		{f.Want, xml1.Want},
		{f.Wishlist, xml1.Wishlist},
		{f.WantToPlay, xml1.WantToPlay},
		{f.WantToBuy, xml1.WantToBuy},
		{f.PrevOwned, xml1.PrevOwned},
		{f.PreOrdered, xml1.PreOrdered},
		{f.HasParts, xml1.HasParts},
		{f.WantParts, xml1.WantParts},
	}
	for _, o := range boolOpts {
		if o.value != nil {
			opts = append(opts, o.option(*o.value))
		}
	}
	intOpts := []struct {
		value  *int
		option func(int) xml1.CollectionOption
	}{
		{f.WishlistPriority, xml1.WishlistPriority},
		{f.MinRating, xml1.MinRating},
		{f.MaxRating, xml1.MaxRating},
		{f.MinBGGRating, xml1.MinBGGRating},
		{f.MaxBGGRating, xml1.MaxBGGRating},
		{f.MinPlays, xml1.MinPlays},
		{f.MaxPlays, xml1.MaxPlays},
	}
	for _, o := range intOpts {
		if o.value != nil {
			opts = append(opts, o.option(*o.value))
		}
	}
	return opts
}
