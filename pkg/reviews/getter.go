package reviews

import (
	"encoding/json"
	"errors"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Review "schema" as found in plattentests header
type ReviewSchema struct {
	Context      string       `json:"@context"`
	Type         string       `json:"@type"`
	ItemReviewed ItemReviewed `json:"itemReviewed"`
	ReviewRating ReviewRating `json:"reviewRating"`
	Author       Author       `json:"author"`
	ReviewBody   string       `json:"reviewBody"`
}
type ByArtist struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}
type ItemReviewed struct {
	Type     string   `json:"@type"`
	Name     string   `json:"name"`
	Image    string   `json:"image"`
	ByArtist ByArtist `json:"byArtist"`
}
type ReviewRating struct {
	Type        string `json:"@type"`
	RatingValue int    `json:"ratingValue"`
	BestRating  string `json:"bestRating"`
}
type Author struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}

type JsonNotFoundError struct {
}
type JsonParseError struct {
	Err error
}

const (
	errorNotFoundJson     = "unable to find json schema"
	errorParseJson        = "unable to parse json schema"
	errorParseArtist      = "unable to parse artist"
	errorParseAlbum       = "unable to parse album"
	errorParseRating      = "unable to parse rating"
	errorParseReleaseDate = "unable to parse release date"
)

func (e *JsonNotFoundError) Error() string {
	return errorNotFoundJson
}
func (e *JsonParseError) Error() string {
	return errorParseJson
}

// in newer reviews there's a json schema
func (r *ReviewClient) getJson() (*ReviewSchema, error) {
	var reviewSchema *ReviewSchema

	reviewJSON := r.Doc.Find("script[type='application/ld+json']").Text()
	if reviewJSON == "" {
		return reviewSchema, &JsonNotFoundError{}
	}

	decodedReviewJSON := html.UnescapeString(reviewJSON)

	err := json.Unmarshal([]byte(decodedReviewJSON), &reviewSchema)
	if err != nil {
		return reviewSchema, &JsonParseError{Err: err}
	}

	return reviewSchema, nil
}

func (r *ReviewClient) getTracklist() []string {
	var tracklist []string
	r.Doc.Find("div#rezitracklist li").Each(func(i int, s *goquery.Selection) {
		tracklist = append(tracklist, strings.TrimSpace(s.Text()))
	})
	return tracklist
}

func (r *ReviewClient) getHighlights() []string {
	var highlights []string
	r.Doc.Find("div#rezihighlights li").Each(func(i int, s *goquery.Selection) {
		highlights = append(highlights, strings.TrimSpace(s.Text()))
	})
	return highlights
}

func (r *ReviewClient) getAuthor() string {
	autor := r.Doc.Find("div#rezitext p.autor a").Text()
	return strings.TrimSpace(autor)
}

func (r *ReviewClient) getText() string {
	var text string
	r.Doc.Find("div#rezitext p").Each(func(i int, s *goquery.Selection) {
		_, exists := s.Attr("class")
		if !exists {
			text += " " + s.Text()
		}
	})
	return strings.TrimSpace(text)
}

func (r *ReviewClient) getSpotifyID() string {
	var spotifyId string
	// <iframe src="https://open.spotify.com/embed/album/2q0yeivzk1b2UUdtHf8mcC"
	spotifyURL, exists := r.Doc.Find("div#reziforum div iframe").Attr("src")
	if exists {
		spotifyId = strings.SplitAfter(spotifyURL, "album/")[1]
	}
	return spotifyId
}

func (r *ReviewClient) getReferences() []string {
	var references []string

	originalRefs := r.Doc.Find("div#reziref p").First().Text()
	// remove duplicate whitespaces from Review references
	reInsideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	originalRefs = reInsideWhitespace.ReplaceAllString(originalRefs, " ")
	// TODO: artist with ;
	references = strings.Split(originalRefs, ";")
	for i := range references {
		references[i] = strings.TrimSpace(references[i])
	}

	return references
}

func (r *ReviewClient) getRating() (int, error) {
	rString := r.Doc.Find("p.bewertung strong").First().Text()
	rExtract := strings.Split(rString, "/")[0]
	rating, err := strconv.Atoi(rExtract)
	if err != nil {
		return 0, err
	}
	return rating, nil
}

func (r *ReviewClient) getReleaseDate() (time.Time, error) {
	var releaseDate time.Time

	releaseDatePT := strings.Split(r.Doc.Find("div.headerbox.rezi p").First().Text(), ": ")[1]
	releaseDate, err := time.Parse("02.01.2006", releaseDatePT)
	if err != nil {
		return releaseDate, err
	}

	return releaseDate, nil
}

// TODO: Artist Name mit " - "?
func (r *ReviewClient) getArtist() (string, error) {
	slice := strings.Split(r.Doc.Find("div.headerbox.rezi h1").Text(), " - ")
	if len(slice) != 2 {
		return "", errors.New(errorParseArtist)
	}
	return strings.TrimSpace(slice[0]), nil
}

// TODO: Album Name mit " - "?
func (r *ReviewClient) getAlbum() (string, error) {
	slice := strings.Split(r.Doc.Find("div.headerbox.rezi h1").Text(), " - ")
	if len(slice) != 2 {
		return "", errors.New(errorParseAlbum)
	}
	return strings.TrimSpace(slice[1]), nil
}

func (r *ReviewClient) getCoverURL() string {
	coverURL, _ := r.Doc.Find("meta[property='og:image']").Attr("content")
	return coverURL
}
