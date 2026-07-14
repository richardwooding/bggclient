package bggopts

import (
	"maps"
	"testing"

	"github.com/richardwooding/bggclient/xml1"
)

func applyBoardgame(t *testing.T, opts []xml1.BoardgameOption) map[string]string {
	t.Helper()
	params := map[string]string{}
	var err error
	for _, opt := range opts {
		params, err = opt(params)
		if err != nil {
			t.Fatalf("applying option: %v", err)
		}
	}
	return params
}

func applyCollection(t *testing.T, opts []xml1.CollectionOption) map[string]string {
	t.Helper()
	params := map[string]string{}
	var err error
	for _, opt := range opts {
		params, err = opt(params)
		if err != nil {
			t.Fatalf("applying option: %v", err)
		}
	}
	return params
}

func TestBoardgameFlags(t *testing.T) {
	tests := []struct {
		name  string
		flags BoardgameFlags
		want  map[string]string
	}{
		{"empty", BoardgameFlags{}, map[string]string{}},
		{"comments on", BoardgameFlags{Comments: new(true)}, map[string]string{"comments": "1"}},
		{"comments off", BoardgameFlags{Comments: new(false)}, map[string]string{"comments": "0"}},
		{"stats and historical", BoardgameFlags{Stats: new(true), Historical: new(true)}, map[string]string{"stats": "1", "historical": "1"}},
		{"date range", BoardgameFlags{From: new("2024-01-01"), To: new("2024-12-31")}, map[string]string{"from": "2024-01-01", "to": "2024-12-31"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts, err := tt.flags.Options()
			if err != nil {
				t.Fatalf("Options() error: %v", err)
			}
			got := applyBoardgame(t, opts)
			if !maps.Equal(got, tt.want) {
				t.Errorf("params = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardgameFlagsInvalidDate(t *testing.T) {
	for _, flags := range []BoardgameFlags{
		{From: new("not-a-date")},
		{To: new("31-12-2024")},
	} {
		if _, err := flags.Options(); err == nil {
			t.Errorf("Options() with %+v: expected error, got nil", flags)
		}
	}
}

func TestCollectionFlags(t *testing.T) {
	tests := []struct {
		name  string
		flags CollectionFlags
		want  map[string]string
	}{
		{"empty", CollectionFlags{}, map[string]string{}},
		{"own", CollectionFlags{Own: new(true)}, map[string]string{"own": "1"}},
		{"own excluded", CollectionFlags{Own: new(false)}, map[string]string{"own": "0"}},
		{
			"bools",
			CollectionFlags{Rated: new(true), Trade: new(true), WantToPlay: new(true), PrevOwned: new(true)},
			map[string]string{"rated": "1", "trade": "1", "wanttoplay": "1", "prevowned": "1"},
		},
		{
			"ints",
			CollectionFlags{WishlistPriority: new(3), MinRating: new(5), MaxPlays: new(10)},
			map[string]string{"wishlistpriority": "3", "minrating": "5", "maxplays": "10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyCollection(t, tt.flags.Options())
			if !maps.Equal(got, tt.want) {
				t.Errorf("params = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollectionFlagsOutOfBounds(t *testing.T) {
	for name, flags := range map[string]CollectionFlags{
		"priority too high": {WishlistPriority: new(6)},
		"rating too low":    {MinRating: new(0)},
		"rating too high":   {MaxRating: new(11)},
	} {
		t.Run(name, func(t *testing.T) {
			params := map[string]string{}
			var err error
			for _, opt := range flags.Options() {
				params, err = opt(params)
				if err != nil {
					return
				}
			}
			t.Error("expected out-of-bounds error, got nil")
		})
	}
}
