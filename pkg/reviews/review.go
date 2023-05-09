package reviews

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

// return a single review (e.g. plattentests.de/rezi.php?show=19238)
type Review struct {
	Artist      string
	Album       string
	ReleaseDate time.Time
	SpotifyID   string
	Rating      int
	References  []string
	CoverURL    string
	Author      string
	Text        string
	Highlights  []string
	Tracklist   []string
	URL         string
}

// our client struct
type ReviewClient struct {
	Doc *goquery.Document
	URL string
}

func New(url string) (*ReviewClient, error) {
	// Request the HTML page.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("received bad http statuscode " + strconv.Itoa(resp.StatusCode))
	}
	// go works on utf-8 only
	utf8body, err := charset.NewReader(resp.Body, "iso-8859-1")
	if err != nil {
		return nil, err
	}
	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(utf8body)
	if err != nil {
		return nil, err
	}

	return &ReviewClient{
		Doc: doc,
		URL: url,
	}, nil

}

func (r *ReviewClient) GetReview() (Review, error) {
	var review Review
	var fallback bool

	review.URL = r.URL

	reviewJson, err := r.getJson()
	if err != nil || reviewJson == nil {
		// slog.Warn(errorParseJson + ", falling back")
		fallback = true
	}

	if !fallback {
		review.Artist = reviewJson.ItemReviewed.ByArtist.Name
		review.Album = reviewJson.ItemReviewed.Name
		review.Rating = reviewJson.ReviewRating.RatingValue
		review.CoverURL = reviewJson.ItemReviewed.Image
		review.Author = reviewJson.Author.Name
	} else {
		review.Artist, err = r.getArtist()
		if err != nil {
			return review, err
		}
		review.Album, err = r.getAlbum()
		if err != nil {
			return review, err
		}
		review.CoverURL = r.getCoverURL()
		review.Author = r.getAuthor()
		review.Rating, err = r.getRating()
		if err != nil {
			return review, err
		}

	}
	review.References = r.getReferences()
	review.SpotifyID = r.getSpotifyID()
	review.Text = r.getText()
	review.Highlights = r.getHighlights()
	review.Tracklist = r.getTracklist()
	review.ReleaseDate, err = r.getReleaseDate()
	if err != nil {
		return review, err
	}

	return review, nil
}
