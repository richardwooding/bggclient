package xml1

import (
	"github.com/richardwooding/bggclient/xml1/customerrors"
	"strconv"
)

type CollectionOption func(m map[string]string) (map[string]string, error)

type includeOption func(include bool) CollectionOption

func paramInclude(param string) includeOption {
	return func(include bool) CollectionOption {
		return func(m map[string]string) (map[string]string, error) {
			if include {
				m[param] = "1"
			} else {
				m[param] = "0"
			}
			return m, nil
		}
	}
}

var Own = paramInclude("own")
var Rated = paramInclude("rated")
var Comment = paramInclude("comment")
var Trade = paramInclude("trade")
var Want = paramInclude("want")
var WantinTrade = paramInclude("wantintrade")
var Wishlist = paramInclude("wishlist")
var WantToPlay = paramInclude("wanttoplay")
var WantToBuy = paramInclude("wanttobuy")
var PrevOwned = paramInclude("prevowned")
var PreOrdered = paramInclude("preordered")
var HasParts = paramInclude("hasparts")
var WantParts = paramInclude("wantparts")
var NotifyContent = paramInclude("notifycontent")
var NotifySale = paramInclude("notifysale")
var NotifyAuxtion = paramInclude("notifyauction")
var ShowPrivate = paramInclude("showprivate")

func WishlistPriority(priority int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		if priority < 1 || priority > 5 {
			return nil, customerrors.PriorityOutOfBoundsError{Priority: priority}
		}
		m["wishlistpriority"] = strconv.Itoa(priority)
		return m, nil
	}
}

func MinRating(rating int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		if rating < 1 || rating > 10 {
			return nil, customerrors.RatingOutOfBoundsError{Rating: rating}
		}
		m["minrating"] = strconv.Itoa(rating)
		return m, nil
	}
}

func MaxRating(rating int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		if rating < 1 || rating > 10 {
			return nil, customerrors.RatingOutOfBoundsError{Rating: rating}
		}
		m["maxrating"] = strconv.Itoa(rating)
		return m, nil
	}
}

func MinBGGRating(rating int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		if rating < 1 || rating > 10 {
			return nil, customerrors.RatingOutOfBoundsError{Rating: rating}
		}
		m["minbggrating"] = strconv.Itoa(rating)
		return m, nil
	}
}

func MaxBGGRating(rating int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		if rating < 1 || rating > 10 {
			return nil, customerrors.RatingOutOfBoundsError{Rating: rating}
		}
		m["maxbggrating"] = strconv.Itoa(rating)
		return m, nil
	}
}

func MinPlays(plays int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		m["minplays"] = strconv.Itoa(plays)
		return m, nil
	}
}

func MaxPlays(plays int) CollectionOption {
	return func(m map[string]string) (map[string]string, error) {
		m["maxplays"] = strconv.Itoa(plays)
		return m, nil
	}
}

func filterString(name, value string) CollectionOption {
	switch name {
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, customerrors.New("Invalid filter name specified")
		}
	}
}

func filterInt(name string, value int) CollectionOption {
	switch name {
	case "wishlist priority":
		return WishlistPriority(value)
	case "min rating", "minimum rating":
		return MinRating(value)
	case "max rating", "maximum rating":
		return MaxRating(value)
	case "min bgg rating", "minimum bgg rating":
		return MinBGGRating(value)
	case "max bgg rating", "maximum bgg rating":
		return MaxBGGRating(value)
	case "min plays", "minimum plays":
		return MinPlays(value)
	case "max plays", "maximum plays":
		return MaxPlays(value)
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, customerrors.New("Invalid filter name specified")
		}
	}
}

func filterBool(name string, value bool) CollectionOption {
	switch name {
	case "own", "own games":
		return Own(value)
	case "rated":
		return Rated(value)
	case "comment":
		return Comment(value)
	case "trade":
		return Trade(value)
	case "want":
		return Want(value)
	case "want in trade":
		return WantinTrade(value)
	case "wishlist":
		return Wishlist(value)
	case "want to play":
		return WantToPlay(value)
	case "want to buy":
		return WantToBuy(value)
	case "prev owned", "previously owned":
		return PrevOwned(value)
	case "pre ordered", "preordered":
		return PreOrdered(value)
	case "has parts":
		return HasParts(value)
	case "want parts":
		return WantParts(value)
	case "notify content":
		return NotifyContent(value)
	case "notify sale":
		return NotifySale(value)
	case "notify auction":
		return NotifyAuxtion(value)
	case "show private":
		return ShowPrivate(value)
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, customerrors.New("Invalid filter name specified")
		}
	}
}

func Filter(name string, value any) CollectionOption {
	switch v := value.(type) {
	case string:
		return filterString(name, v)
	case int:
		return filterInt(name, v)
	case bool:
		return filterBool(name, v)
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, customerrors.New("Invalid filter value specified")
		}
	}
}
