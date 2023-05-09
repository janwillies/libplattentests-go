package mainpage

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGetRezis(t *testing.T) {

	mc := testServer("../../test/testdata/mainpage/index-1.html", t)

	t.Run("get-adw", func(t *testing.T) {
		got, err := mc.getRezis(selAlbumDerWoche)
		if err != nil {
			t.Errorf("unable to get Album der Woche: %q", err)
		}
		want := []int{19274}
		if !equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get-highlights", func(t *testing.T) {
		got, err := mc.getRezis(selAktuelleHighlights)
		if err != nil {
			t.Errorf("unable to get Aktuelle Highlights: %q", err)
		}
		want := []int{
			19259, 19262, 19269, 19266, 19261, 19271, 19267, 19265,
		}
		if !equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get-neue-rezis", func(t *testing.T) {
		got, err := mc.getRezis(selNeueRezis)
		if err != nil {
			t.Errorf("unable to get Neue Rezensionen: %q", err)
		}
		want := []int{
			19260, 19261, 19262, 19263, 19264, 19265, 19266, 19267, 19268, 19259, 19269, 19270, 19271, 19272, 19273, 19274,
		}
		if !equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("get-weiter-aktuell", func(t *testing.T) {
		got, err := mc.getRezis(selWeiterAktuell)
		if err != nil {
			t.Errorf("unable to get Weiter Aktuell: %q", err)
		}
		want := []int{
			19243, 19254, 19241, 19245, 19216, 19226, 19242, 19247, 19258, 19230, 19231, 19219, 19250, 19212, 19235, 19251, 19252, 19253, 19238, 19257, 19256, 19255,
		}
		if !equal(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
func TestGetPreviews(t *testing.T) {

	mc := testServer("../../test/testdata/mainpage/index-1.html", t)

	t.Run("get-preview", func(t *testing.T) {
		got, err := mc.getPreviews(selVorschau)
		if err != nil {
			t.Errorf("unable to get Vorschau: %q", err)
		}
		relDate1, _ := time.Parse("02.01.2006", "14.04.2023")
		relDate2, _ := time.Parse("02.01.2006", "28.04.2023")
		relDate3, _ := time.Parse("02.01.2006", "05.05.2023")
		want := []Preview{
			{Artist: "Feist", Album: "Multitudes", ReleaseDate: relDate1},
			{Artist: "Philip Bradatsch", Album: "Philip Bradatsch", ReleaseDate: relDate2},
			{Artist: "Ed Sheeran", Album: "-", ReleaseDate: relDate3},
		}
		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestGetMainpage(t *testing.T) {

	mc := testServer("../../test/testdata/mainpage/index-1.html", t)

	t.Run("get-mainpage", func(t *testing.T) {
		got, err := mc.GetMainpage()
		if err != nil {
			t.Errorf("unable to get mainpage: %q", err)
		}
		relDate1, _ := time.Parse("02.01.2006", "14.04.2023")
		relDate2, _ := time.Parse("02.01.2006", "28.04.2023")
		relDate3, _ := time.Parse("02.01.2006", "05.05.2023")
		want := Mainpage{
			Hashsum: 731575,
			Previews: []Preview{
				{Artist: "Feist", Album: "Multitudes", ReleaseDate: relDate1},
				{Artist: "Philip Bradatsch", Album: "Philip Bradatsch", ReleaseDate: relDate2},
				{Artist: "Ed Sheeran", Album: "-", ReleaseDate: relDate3},
			},
			ReviewIdCat: map[int]string{
				19212: "weiteraktuell",
				19216: "weiteraktuell",
				19219: "weiteraktuell",
				19226: "weiteraktuell",
				19230: "weiteraktuell",
				19231: "weiteraktuell",
				19235: "weiteraktuell",
				19238: "weiteraktuell",
				19241: "weiteraktuell",
				19242: "weiteraktuell",
				19243: "weiteraktuell",
				19245: "weiteraktuell",
				19247: "weiteraktuell",
				19250: "weiteraktuell",
				19251: "weiteraktuell",
				19252: "weiteraktuell",
				19253: "weiteraktuell",
				19254: "weiteraktuell",
				19255: "weiteraktuell",
				19256: "weiteraktuell",
				19257: "weiteraktuell",
				19258: "weiteraktuell",
				19259: "akt_highlights",
				19260: "neuerezis",
				19261: "akt_highlights",
				19262: "akt_highlights",
				19263: "neuerezis",
				19264: "neuerezis",
				19265: "akt_highlights",
				19266: "akt_highlights",
				19267: "akt_highlights",
				19268: "neuerezis",
				19269: "akt_highlights",
				19270: "neuerezis",
				19271: "akt_highlights",
				19272: "neuerezis",
				19273: "neuerezis",
				19274: "adw",
			},
		}
		if !cmp.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func testServer(testFile string, t testing.TB) *MainpageClient {
	handler := func(w http.ResponseWriter, r *http.Request) {
		htmlByte, err := os.ReadFile(testFile)
		if err != nil {
			t.Error("unable to read html file: " + err.Error())
		}
		fmt.Fprint(w, string(htmlByte))
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	mc, err := NewWithUrl(srv.URL)
	if err != nil {
		t.Errorf("unable to setup mainpage client: %v", err)
	}

	return mc
}
