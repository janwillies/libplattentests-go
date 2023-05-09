package reviews

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestGetReview(t *testing.T) {

	rc, url := testServer("../../test/testdata/reviews/18380-witt-no-json.html", t)

	t.Run("json-not-found", func(t *testing.T) {
		_, err := rc.getJson()
		if err != err.(*JsonNotFoundError) {
			t.Errorf("got %q want %q", err, err.(*JsonNotFoundError))
		}
	})
	t.Run("get-review-fallback", func(t *testing.T) {
		got, err := rc.GetReview()
		if err != nil {
			t.Errorf("unable to get review: %v", err)
		}

		releaseDate, _ := time.Parse("02.01.2006", "25.02.2022")
		want := Review{
			Artist:      "Joachim Witt",
			Album:       "Rübezahls Reise",
			ReleaseDate: releaseDate,
			Rating:      1,
			Author:      "Christopher Sennfelder",
			Highlights:  []string{"-"},
			Tracklist: []string{
				"Rübezahl",
				"In Einsamkeit (feat. Chris Harms)",
				"So fern",
				"Shandai ya (feat. The Mystery Of The Bulgarian Voices)",
				"Die Wölfe ziehen",
				"Abendwind",
				"Das Leben in mir",
				"Stern (feat. Claudia Uhle)",
				"Die Seele",
				"Bernstein",
				"Ich spür die Liebe in mir",
			},
			SpotifyID: "1r0u9npH7PmJtkI3Dhqusc",
			URL:       url,
			CoverURL:  "https://www.plattentests.de/cover/witt13.jpg",
			References: []string{
				"Witt", "Witt & Orchester", "Unheilig", "Mono Inc.", "Lord Of The Lost", "Hubert Kah & Die Terroristen Der Liebe", "Herren", "Wunder", "Subway To Sally", "Deine Lakaien", "Wolfsheim", "Peter Heppner", "Megaherz", "Eisbrecher", "In Extremo", "Eisblume", "Blutengel", "Goethes Erben", "Oomph!", "Rammstein", "Richard Wagner", "HIM", "Unzucht", "Lacrimosa", "Samsas Traum", "Nu Pagadi", "Schandmaul", "Saltatio Mortis", "Tanzwut", "Oonagh", "Santiano", "Staubkind", "Funker Vogt", "Rosenstolz", "Nina Hagen", "X-Perience", "Angelzoom", "Apoptygma Berzerk", "Welle:Erdball", "Knorkator", "Heilung", "Faun", "Matthias Reim", "Le Mystère Des Voix Bulgares", "Helene Fischer",
			},
			Text: "Die meisten Menschen mögen Musik. Für alle anderen gibt es Joachim Witt. Sein neues Album \"Rübezahls Reise\" existiert, weil zwei Platten noch keine Trilogie bilden. Immerhin starrt dem Käufer diesmal kein bärtiger Schrat auf dem Cover entgegen. Wobei zur Diskussion steht, inwiefern sich Käufer eines Witt-Albums überhaupt von irgendetwas abschrecken lassen. Kompromisslos sind sie zweifelsohne. Kühne Recken, die für eine Welt ohne Geschmack kämpfen und ihrem Idol die Nibelungentreue halten, koste es, was es wolle. Man könnte ihnen jetzt Unvernunft attestieren, aber das wäre gemein. Nicht gemein ist es, die Klänge auf \"Rübezahls Reise\" als unerträglich zu bezeichnen. Bei allem Verständnis für Musik als Ausdruck von Individualität gibt es Grenzen. Da helfen auch keine bulgarischen Chöre. Dabei existieren sie, diese Momente, in denen aufblitzt, dass durchaus fähige Musiker am Werk waren. Das instrumentale Finale von \"So fern\" ist zwar unglaublich pathetisch, eine gewisse Eingängigkeit kann man ihm aber nicht absprechen. Wenn da nur nicht die schreckliche Lyrik des Herrn Witt wäre, welcher sich in erschütternden Versen der kritischen Selbstbetrachtung widmet. Überhaupt, die Texte. Immer wieder greift Witt tief in den Metapherntopf und fördert dabei erstaunliche Schöpfungen zutage. Besonders ergreifend geraten seine Worte in \"Abendwind\". Aus Gründen wird hier auf ein Zitat verzichtet. Man stelle sich einfach eine Kreuzung aus Joseph von Eichendorff und einem Bluterguss vor. Die Musik gibt ihr Bestes, um die tiefschürfenden Ergüsse des Sängers ins rechte Licht zu rücken und scheitert krachend. Apropos krachend: Es wemmst gewaltig im Hause Witt. Tiefstgestimmte Gitarren, tribalistisches Getrommel und allerhand orchestrales Gedöns bilden einen Klangteppich, der auf dem Sperrmüll am besten aufgehoben wäre. Vielleicht ist Witts stoisches Festhalten an seinem Sound ein Statement zur Nachhaltigkeit. Vielleicht ist er aber auch taub. Dies hindert ihn jedoch nicht daran, die ganz großen Botschaften anzupacken. In \"Ich spür die Liebe in mir\" geht es um die Waldeinsamkeit. Das ist ein schönes Wort, dem der Künstler mit seiner grässlichen Musik nicht im Geringsten gerecht wird. Anders gesagt: Das alles klingt, als hätte jemand Zarathustra auf Wish bestellt. Der Preis für das unwahrscheinlichste Feature des Jahres geht indessen an Claudia Uhle, ihres Zeichens Sängerin von X-Perience, die in \"Stern\" wahrlich Meisterhaftes zum Besten gibt: \"Durch die Nacht, da führt mich dein Stern / Scheint so nah und doch noch so fern / Ist der Mut, der mich trägt / Und die Hoffnung, die lebt / Gib mir Kraft, für mich leuchtet dein Stern\". Dazu ertönen allen Ernstes Dudelsäcke im Hintergrund. Das kann, das darf doch nicht ernst gemeint sein. Ist es vielleicht auch nicht. Das ist ja das Perfide an Joachim Witt. Er bringt es in \"Die Seele\" sogar selbst auf den Punkt: \"Die Zeichen der Zeit veröden im Dickicht der Verantwortlichkeit.\" So ist das. Am Ende war dann wieder keiner zuständig.",
		}

		// assertCorrectMessage(t, got.Artist, want.Artist)
		// assertCorrectMessage(t, got.Album, want.Album)
		// assertCorrectMessage(t, got.Author, want.Author)
		// assertCorrectMessage(t, got.CoverURL, want.CoverURL)
		// assertCorrectMessage(t, got.SpotifyID, want.SpotifyID)
		// assertCorrectMessage(t, got.Text, want.Text)
		// assertCorrectMessage(t, got.Highlights[0], want.Highlights[0])
		// assertCorrectMessage(t, got.Tracklist[3], want.Tracklist[3])
		// assertCorrectMessage(t, got.References[5], want.References[5])

		if !cmp.Equal(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("get-review", func(t *testing.T) {
		rc, url := testServer("../../test/testdata/reviews/18380-witt.html", t)

		got, err := rc.GetReview()
		if err != nil {
			t.Errorf("unable to get review: %v", err)
		}

		releaseDate, _ := time.Parse("02.01.2006", "25.02.2022")
		want := Review{
			Artist:      "Joachim Witt",
			Album:       "Rübezahls Reise",
			ReleaseDate: releaseDate,
			Rating:      1,
			Author:      "Christopher Sennfelder",
			Highlights:  []string{"-"},
			Tracklist: []string{
				"Rübezahl",
				"In Einsamkeit (feat. Chris Harms)",
				"So fern",
				"Shandai ya (feat. The Mystery Of The Bulgarian Voices)",
				"Die Wölfe ziehen",
				"Abendwind",
				"Das Leben in mir",
				"Stern (feat. Claudia Uhle)",
				"Die Seele",
				"Bernstein",
				"Ich spür die Liebe in mir",
			},
			SpotifyID: "1r0u9npH7PmJtkI3Dhqusc",
			URL:       url,
			CoverURL:  "https://www.plattentests.de/witt13.jpg",
			References: []string{
				"Witt", "Witt & Orchester", "Unheilig", "Mono Inc.", "Lord Of The Lost", "Hubert Kah & Die Terroristen Der Liebe", "Herren", "Wunder", "Subway To Sally", "Deine Lakaien", "Wolfsheim", "Peter Heppner", "Megaherz", "Eisbrecher", "In Extremo", "Eisblume", "Blutengel", "Goethes Erben", "Oomph!", "Rammstein", "Richard Wagner", "HIM", "Unzucht", "Lacrimosa", "Samsas Traum", "Nu Pagadi", "Schandmaul", "Saltatio Mortis", "Tanzwut", "Oonagh", "Santiano", "Staubkind", "Funker Vogt", "Rosenstolz", "Nina Hagen", "X-Perience", "Angelzoom", "Apoptygma Berzerk", "Welle:Erdball", "Knorkator", "Heilung", "Faun", "Matthias Reim", "Le Mystère Des Voix Bulgares", "Helene Fischer",
			},
			Text: "Die meisten Menschen mögen Musik. Für alle anderen gibt es Joachim Witt. Sein neues Album \"Rübezahls Reise\" existiert, weil zwei Platten noch keine Trilogie bilden. Immerhin starrt dem Käufer diesmal kein bärtiger Schrat auf dem Cover entgegen. Wobei zur Diskussion steht, inwiefern sich Käufer eines Witt-Albums überhaupt von irgendetwas abschrecken lassen. Kompromisslos sind sie zweifelsohne. Kühne Recken, die für eine Welt ohne Geschmack kämpfen und ihrem Idol die Nibelungentreue halten, koste es, was es wolle. Man könnte ihnen jetzt Unvernunft attestieren, aber das wäre gemein. Nicht gemein ist es, die Klänge auf \"Rübezahls Reise\" als unerträglich zu bezeichnen. Bei allem Verständnis für Musik als Ausdruck von Individualität gibt es Grenzen. Da helfen auch keine bulgarischen Chöre. Dabei existieren sie, diese Momente, in denen aufblitzt, dass durchaus fähige Musiker am Werk waren. Das instrumentale Finale von \"So fern\" ist zwar unglaublich pathetisch, eine gewisse Eingängigkeit kann man ihm aber nicht absprechen. Wenn da nur nicht die schreckliche Lyrik des Herrn Witt wäre, welcher sich in erschütternden Versen der kritischen Selbstbetrachtung widmet. Überhaupt, die Texte. Immer wieder greift Witt tief in den Metapherntopf und fördert dabei erstaunliche Schöpfungen zutage. Besonders ergreifend geraten seine Worte in \"Abendwind\". Aus Gründen wird hier auf ein Zitat verzichtet. Man stelle sich einfach eine Kreuzung aus Joseph von Eichendorff und einem Bluterguss vor. Die Musik gibt ihr Bestes, um die tiefschürfenden Ergüsse des Sängers ins rechte Licht zu rücken und scheitert krachend. Apropos krachend: Es wemmst gewaltig im Hause Witt. Tiefstgestimmte Gitarren, tribalistisches Getrommel und allerhand orchestrales Gedöns bilden einen Klangteppich, der auf dem Sperrmüll am besten aufgehoben wäre. Vielleicht ist Witts stoisches Festhalten an seinem Sound ein Statement zur Nachhaltigkeit. Vielleicht ist er aber auch taub. Dies hindert ihn jedoch nicht daran, die ganz großen Botschaften anzupacken. In \"Ich spür die Liebe in mir\" geht es um die Waldeinsamkeit. Das ist ein schönes Wort, dem der Künstler mit seiner grässlichen Musik nicht im Geringsten gerecht wird. Anders gesagt: Das alles klingt, als hätte jemand Zarathustra auf Wish bestellt. Der Preis für das unwahrscheinlichste Feature des Jahres geht indessen an Claudia Uhle, ihres Zeichens Sängerin von X-Perience, die in \"Stern\" wahrlich Meisterhaftes zum Besten gibt: \"Durch die Nacht, da führt mich dein Stern / Scheint so nah und doch noch so fern / Ist der Mut, der mich trägt / Und die Hoffnung, die lebt / Gib mir Kraft, für mich leuchtet dein Stern\". Dazu ertönen allen Ernstes Dudelsäcke im Hintergrund. Das kann, das darf doch nicht ernst gemeint sein. Ist es vielleicht auch nicht. Das ist ja das Perfide an Joachim Witt. Er bringt es in \"Die Seele\" sogar selbst auf den Punkt: \"Die Zeichen der Zeit veröden im Dickicht der Verantwortlichkeit.\" So ist das. Am Ende war dann wieder keiner zuständig.",
		}

		assertCorrectMessage(t, got.Artist, want.Artist)
		assertCorrectMessage(t, got.Album, want.Album)
		assertCorrectMessage(t, got.Author, want.Author)
		assertCorrectMessage(t, got.CoverURL, want.CoverURL)
		assertCorrectMessage(t, got.SpotifyID, want.SpotifyID)
		assertCorrectMessage(t, got.Text, want.Text)
		assertCorrectMessage(t, got.Highlights[0], want.Highlights[0])
		assertCorrectMessage(t, got.Tracklist[3], want.Tracklist[3])
		assertCorrectMessage(t, got.References[5], want.References[5])

		// if !cmp.Equal(*got, want) {
		// 	t.Errorf("got %q want %q", *got, want)
		// }
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func testServer(testFile string, t testing.TB) (*ReviewClient, string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		htmlByte, err := os.ReadFile(testFile)
		if err != nil {
			t.Error("unable to read html file: " + err.Error())
		}
		fmt.Fprint(w, string(htmlByte))
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	rc, err := New(srv.URL)
	if err != nil {
		t.Errorf("unable to setup review client: %v", err)
	}

	return rc, srv.URL
}
