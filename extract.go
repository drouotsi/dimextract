package dimextract

import (
    "fmt"
    "regexp"
)

// Can't deal with length/width/height identification consistently
type Dimension struct {
    d1 int
    d2 int
    d3 int
}

func (d *Dimension) str() string {
    return fmt.Sprintf("(%d / %d / %d )", d.d1, d.d2, d.d3)
}

func prepareAllRegexp() *regexp.Regexp {
    re := regexp.MustCompile(`(?:(\d+) *(?:\*|x|X|-|,|/)?? *)+`)
    re.Longest()
    return re
}
//var re = regexp.MustCompile(`(?P<unit>(?:c|m)+)`)

var re = prepareAllRegexp()

//var re = regexp.MustCompile(`(?:(?P<dim>\d+)[:blank:]*(?:x|X|-|,|/)??[:blank:]*)+(?P<unit>(?:c|m)+)*?`)

func ExtractDims(desc string) (Dimension, error) {

    fmt.Println(desc)
    match := re.FindStringSubmatch(desc)

    fmt.Println(match)

    result := make(map[string]string)
    for i, name := range re.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
    }

    fmt.Println(result)
    var d Dimension
    return d, nil
}