# libplattentests-go
library to interact with [plattentests.de](https://www.plattentests.de) from go

currently supports scraping the following:
- index page (aka "mainpage" [plattentests.de/index.php](https://www.plattentests.de/index.php))
- detailed review pages (e.g. [plattentests.de/rezi.php?show=3073](https://www.plattentests.de/rezi.php?show=3073))

## usage
### mainpage
```go
import (
	"fmt"

	"github.com/janwillies/libplattentests-go/pkg/mainpage"
)

func main() {
    mc, err := mainpage.New()
    if err != nil {
        log.Fatal(err)
    }
    
    mainpage, err := mc.GetMainpage()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("%+v\n", mainpage)
}
```
### review
```go
import (
	"fmt"

	"github.com/janwillies/libplattentests-go/pkg/reviews"
)
func main() {
    rc, err := reviews.New("https://www.plattentests.de/rezi.php?show=3073")
    if err != nil {
        log.Fatal(err)
    }
    
    review, err := rc.GetReview()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", review)
}
```

## tests
some initial tests, I'm sure there could be more. PRs welcome
```bash
$ go test -coverpkg=./... ./...

?   	github.com/janwillies/libplattentests-go	[no test files]
ok  	github.com/janwillies/libplattentests-go/pkg/mainpage	0.294s	coverage: 81.3% of statements in ./...
ok  	github.com/janwillies/libplattentests-go/pkg/reviews	0.434s	coverage: 86.0% of statements in ./...
```

## credit
all credit goes to plattentests.de for maintaining such a cool project for over 20 years now!