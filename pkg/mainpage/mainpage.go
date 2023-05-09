package mainpage

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type Mainpage struct {
	Hashsum     int
	ReviewIdCat map[int]string
	Previews    []Preview
}

type Preview struct {
	Artist      string
	Album       string
	ReleaseDate time.Time
}

type MainpageClient struct {
	Doc *goquery.Document
}

const (
	baseUrl               = "https://www.plattentests.de/index.php"
	selAlbumDerWoche      = "div.adw h3"
	selAktuelleHighlights = "#akt_highlights div div"
	selNeueRezis          = ".neuerezis li"
	selWeiterAktuell      = "#weiteraktuell li"
	selVorschau           = "#vorschau li"
	catAlbumDerWoche      = "adw"
	catAktuelleHighlights = "akt_highlights"
	catNeueRezis          = "neuerezis"
	catWeiterAktuell      = "weiteraktuell"
)

func New() (*MainpageClient, error) {
	return NewWithUrl(baseUrl)
}

func NewWithUrl(url string) (*MainpageClient, error) {
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

	return &MainpageClient{
		Doc: doc,
	}, nil

}

func (mc *MainpageClient) GetMainpage() (Mainpage, error) {
	var mainpage Mainpage

	mainpage.ReviewIdCat = make(map[int]string)

	// Find the other Review items (bottom right)
	others, err := mc.getRezis(selWeiterAktuell)
	if err != nil {
		return mainpage, err
	}
	for _, v := range others {
		mainpage.ReviewIdCat[v] = catWeiterAktuell
	}

	// Find the newest Review items (left menu)
	newest, err := mc.getRezis(selNeueRezis)
	if err != nil {
		return mainpage, err
	}
	for _, v := range newest {
		mainpage.ReviewIdCat[v] = catNeueRezis
	}

	// Find the current highlights (center)
	highlights, err := mc.getRezis(selAktuelleHighlights)
	if err != nil {
		return mainpage, err
	}
	for _, v := range highlights {
		mainpage.ReviewIdCat[v] = catAktuelleHighlights
	}

	// Find the Review of the week (top)
	adw, err := mc.getRezis(selAlbumDerWoche)
	if err != nil {
		return mainpage, err
	}
	for _, v := range adw {
		mainpage.ReviewIdCat[v] = catAlbumDerWoche
	}

	// Find upcoming reviews
	previews, err := mc.getPreviews(selVorschau)
	if err != nil {
		return mainpage, err
	}
	mainpage.Previews = previews

	// create hashsum so we can later check whether we already have a copy
	for k := range mainpage.ReviewIdCat {
		mainpage.Hashsum += k
	}

	return mainpage, nil
}

func (mc *MainpageClient) getRezis(selector string) ([]int, error) {
	var ids []int
	var err error
	mc.Doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		link, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}
		id, errIn := findId(link)
		if errIn != nil {
			err = errIn
		}
		ids = append(ids, id)
	})
	return ids, err
}

func (mc *MainpageClient) getPreviews(selector string) ([]Preview, error) {
	var previews []Preview
	var err error
	mc.Doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		var preview Preview

		artistSlice := strings.Split(s.Text(), " - ")
		preview.Artist = strings.TrimSpace(artistSlice[0])

		albumSlice := strings.Split(artistSlice[1], "(")
		preview.Album = strings.TrimSpace(albumSlice[0])

		releaseSlice := strings.Split(albumSlice[1], ")")
		releaseSlice = strings.Split(releaseSlice[0], " ")
		releaseDate, errIn := time.Parse("02.01.2006", releaseSlice[1])
		if errIn != nil {
			err = errIn
		}
		preview.ReleaseDate = releaseDate

		previews = append(previews, preview)
	})
	return previews, err
}

// extract id from rezi.php?show=18373
func findId(link string) (int, error) {
	idSlice := strings.SplitAfter(link, "=")
	id, err := strconv.Atoi(idSlice[len(idSlice)-1])
	if err != nil {
		return 0, err
	}
	return id, nil
}
