package reviews

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetJson(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/18346.html", t)

	t.Run("artist-with-dash", func(t *testing.T) {
		got, err := rc.getJson()
		if err != nil {
			t.Errorf("unable to get json: %q", err)
		}
		want := "Super - chunk"
		assertCorrectMessage(t, got.ItemReviewed.ByArtist.Name, want)
	})

	t.Run("album-with-dash", func(t *testing.T) {
		got, err := rc.getJson()
		if err != nil {
			t.Errorf("unable to get json: %q", err)
		}
		want := "Wild -  loneliness"
		assertCorrectMessage(t, got.ItemReviewed.Name, want)
	})

	t.Run("ancient-review", func(t *testing.T) {
		rc, _ := testServer("../../test/testdata/reviews/1.html", t)
		got, err := rc.getJson()
		if err != nil {
			t.Errorf("unable to get json: %q", err)
		}
		if got == nil {
			t.Errorf("nil pointer")
			t.FailNow()
		}
		want := ReviewSchema{
			Context: "http://schema.org/",
			Type:    "Review",
			ItemReviewed: ItemReviewed{
				Type:  "MusicAlbum",
				Name:  "Everything, everything",
				Image: "https://www.plattentests.de/underw.jpg",
				ByArtist: ByArtist{
					Type: "MusicGroup",
					Name: "Underworld",
				},
			},
			ReviewRating: ReviewRating{
				Type:        "Rating",
				RatingValue: 8,
				BestRating:  "10",
			},
			Author: Author{
				Type: "Person",
				Name: "Gerd Bezold",
			},
			ReviewBody: "Underworld, live and everything",
		}
		if !cmp.Equal(*got, want) {
			t.Errorf("got %q want %q", *got, want)
		}
	})
	t.Run("newer-review", func(t *testing.T) {
		rc, _ := testServer("../../test/testdata/reviews/18381.html", t)
		got, err := rc.getJson()
		if err != nil {
			t.Errorf("unable to get json: %q", err)
		}
		if got == nil {
			t.Errorf("nil pointer")
			t.FailNow()
		}
		want := ReviewSchema{
			Context: "http://schema.org/",
			Type:    "Review",
			ItemReviewed: ItemReviewed{
				Type:  "MusicAlbum",
				Name:  "Hygiene",
				Image: "https://www.plattentests.de/drugchur.jpg",
				ByArtist: ByArtist{
					Type: "MusicGroup",
					Name: "Drug Church",
				},
			},
			ReviewRating: ReviewRating{
				Type:        "Rating",
				RatingValue: 8,
				BestRating:  "10",
			},
			Author: Author{
				Type: "Person",
				Name: "Kevin Holtmann",
			},
			ReviewBody: "Heilig's Krächle",
		}
		if !cmp.Equal(*got, want) {
			t.Errorf("got %q want %q", *got, want)
		}
	})
}
func TestGetAlbum(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/18346.html", t)

	t.Run("fallback-album-with-dash", func(t *testing.T) {
		got, _ := rc.getAlbum()
		want := "Wild -loneliness"
		assertCorrectMessage(t, got, want)
	})

	rc, _ = testServer("../../test/testdata/reviews/1.html", t)
	t.Run("album", func(t *testing.T) {
		got, _ := rc.getAlbum()
		want := "Everything, everything"
		assertCorrectMessage(t, got, want)
	})
}
func TestGetArtist(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/18346.html", t)
	t.Run("fallback-artist-with-dash", func(t *testing.T) {
		got, _ := rc.getArtist()
		want := "Super- chunk"
		assertCorrectMessage(t, got, want)
	})

	rc, _ = testServer("../../test/testdata/reviews/1.html", t)
	t.Run("artist", func(t *testing.T) {
		got, _ := rc.getArtist()
		want := "Underworld"
		assertCorrectMessage(t, got, want)
	})
}
func TestGetTracklist(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/1.html", t)
	t.Run("tracklist", func(t *testing.T) {
		got := rc.getTracklist()
		want := "Rez/Cowgirl"
		assertCorrectMessage(t, got[6], want)
	})

	rc, _ = testServer("../../test/testdata/reviews/18381.html", t)
	t.Run("tracklist", func(t *testing.T) {
		got := rc.getTracklist()
		want := "Athlete on bench"
		assertCorrectMessage(t, got[9], want)
	})
}
func TestGetHighlights(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/1.html", t)
	t.Run("highlights", func(t *testing.T) {
		got := rc.getHighlights()
		want := "Rez/Cowgirl"
		assertCorrectMessage(t, got[2], want)
	})
}
func TestGetAuthor(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/1.html", t)
	t.Run("autor", func(t *testing.T) {
		got := rc.getAuthor()
		want := "Gerd Bezold"
		assertCorrectMessage(t, got, want)
	})
}
func TestGetSpotifyID(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/1.html", t)
	t.Run("spotify", func(t *testing.T) {
		got := rc.getSpotifyID()
		want := "2QDm4fRgSCnQF5kLUiI2kS"
		assertCorrectMessage(t, got, want)
	})
}
func TestGetCoverURL(t *testing.T) {

	rc, _ := testServer("../../test/testdata/reviews/1.html", t)
	t.Run("coverURL", func(t *testing.T) {
		got := rc.getCoverURL()
		want := "https://www.plattentests.de/cover/underw.jpg"
		assertCorrectMessage(t, got, want)
	})
}

func TestNewerReview(t *testing.T) {

	// srv := httptest.NewServer(testHandler("../../test/testdata/reviews/18381.html"))
	// rc, err := New(srv.URL)
	// if err != nil {
	// 	t.Errorf("unable to setup review client: %v", err)
	// }

	// t.Run("tracklist", func(t *testing.T) {
	// 	got := rc.getTracklist()
	// 	want := "Athlete on bench"
	// 	assertCorrectMessage(t, got[9], want)
	// })

	// t.Run("highlights", func(t *testing.T) {
	// 	got := rc.getHighlights()
	// 	want := "World impact"
	// 	assertCorrectMessage(t, got[3], want)
	// })

	// t.Run("rezitext", func(t *testing.T) {
	// 	got := rc.getText()
	// 	want := "\"Ja, heilig's Blechle, was ist denn hier los?\", möchte man nach einigen Durchgängen dieser Platte fragen. \"Ein Heidenlärm, der Laune macht!\", so lautet die korrekte Antwort. Dafür verantwortlich: Drug Church, eine Post-Hardcore-Band aus Albany im US-Bundesstaat New York. Die Band startete einst als Nebenprojekt von Frontmann Patrick Kindlon, der bereits mit seiner Hauptband Self Defense Family energiegeladenen Post-Hardcore und schwerlidrigen New Wave verkuppelte. Mit Drug Church schlägt er in eine ähnliche Kerbe, wobei die Band bislang größtenteils unter dem Radar segelte. Ob sich das mit \"Hygiene\", der vierten Albumveröffentlichung, wesentlich ändert? Es darf bezweifelt werden. Verdient wäre es allemal, denn schlüssiger als hier kann man Melodien und Krachkomponenten ja kaum miteinander verschweißen. Drug Church bleiben bei all dem schwer zu fassen: Ihr Sound speist sich freilich aus Hardcore, Post-Hardcore und der eher robusten Emo-Schiene, umfasst aber auch Elemente aus Grunge, Pop, Wave und – wenn man so möchte – sogar Nu Metal. Natürlich wird es hier nicht peinlich, sondern bleibt stets der großen und hehren Idee verpflichet, einen Klang zu kreieren, der energetisch und stürmisch ist, der in guten Momenten gar euphorische Höhen erreicht, der aber beizeiten auch den Abgrund kennt und davon eindrucksvoll berichtet. Wer sich eine von Andrew W. K. organisierte 80s-Party vorstellt, auf der zwischen Joy Division und Bauhaus auch mal Turnstile läuft, der bekommt eine Idee davon, wie sich \"Hygiene\" in manchen Momenten anfühlt. Und doch greift diese Metapher zu kurz, denn natürlich fehlt ein Querverweis zu den Speerspitzen der Emo- und Hardcore-Szene, zu Touché Amoré und La Dispute, die hier durchaus auch als Referenzen genannt werden sollten. Der Opener \"Fun's over\" führt mit seinem Titel natürlich auf die falsche Fährte, denn der Spaß beginnt ja gerade erst: Knackige Gitarren spielen sich in diesen knappen zwei Minuten in einen furiosen Rausch, Kindlon spricht eher, als dass er singt, aber für ein einleitendes Kapitel legt der Song schon ziemlich eindrucksvoll offen, wohin die Reise hier geht. Drug Church spielen sich hier durch sehnigen, kraftvollen Hi-Energy-Rock, der an den Rändern dunkel schimmert. \"Millions miles of fun\" ist dann so etwas wie der Hit der Platte: Die Gitarren fahren einen halsbrecherischen Schlingerkurs, ihr Post-Hardcore klingt metallisch, bei aller Härte und Prägnanz verzichten die US-Amerikaner aber nicht auf maximale Eingängigkeit. Klar wird dies auch im melancholischen Post-Punk von \"Detective lieutenant\", der den Sound der späten Siebziger- und frühen Achtzigerjahre ganz unverhohlen würdigt – insbesondere der Bass träumt sich in eine Zeit zurück, in der Ian Curtis noch seine traurigen Lieder sang. Drug Church liefern auf \"Hygiene\" eine fein austarierte Idee davon, wie Post-Hardcore im Jahr 2022 klingen kann: scharfkantig und aggressiv wie im Dampfwalzenkommando von \"Tiresome\". Oder an Fucked Up erinnernd wie im adrenalingeschwängerten Punkrock von \"World impact\". Kein Bullshit, kein Chichi. Was hier los ist? Der Teufel."
	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("autor", func(t *testing.T) {
	// 	got := rc.getAuthor()
	// 	want := "Kevin Holtmann"
	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("spotify", func(t *testing.T) {
	// 	got := rc.getSpotifyID()
	// 	want := ""
	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("album", func(t *testing.T) {
	// 	got, _ := rc.getAlbum()
	// 	want := "Hygiene"
	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("artist", func(t *testing.T) {
	// 	got, _ := rc.getArtist()
	// 	want := "Drug Church"
	// 	assertCorrectMessage(t, got, want)
	// })

	// t.Run("coverURL", func(t *testing.T) {
	// 	got := rc.getCoverURL()
	// 	want := "https://www.plattentests.de/cover/drugchur.jpg"
	// 	assertCorrectMessage(t, got, want)
	// })
}
