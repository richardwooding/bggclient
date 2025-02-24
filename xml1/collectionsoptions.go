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
