package main

import (
	"fmt"
	"os"

	"github.com/janwillies/libplattentests-go/pkg/mainpage"
	"github.com/janwillies/libplattentests-go/pkg/reviews"

	"golang.org/x/exp/slog"
)

func main() {
	mc, err := mainpage.New()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	mainpage, err := mc.GetMainpage()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v\n", mainpage)

	rc, err := reviews.New("https://www.plattentests.de/rezi.php?show=18380")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	review, err := rc.GetReview()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v\n", review)
}
