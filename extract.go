package dimextract

import (
    "fmt"
    "regexp"
    "sort"
    "strconv"
    "strings"
)

// Can't deal with length/width/height identification consistently
// I will try to find 3 dimensions where d1 is the highest
type Dimension struct {
    D1 int
    D2 int
    D3 int
}

func (d *Dimension) str() string {
    return fmt.Sprintf("(%d / %d / %d )", d.D1, d.D2, d.D3)
}

func prepareAllRegexp() (*regexp.Regexp, *regexp.Regexp, *regexp.Regexp) {
    reX := regexp.MustCompile(`((?:\d+ *(?:\*|x|X|-|,|/)?? *){2,4}) *((?:c|m|C|M)*)`)
    reX.Longest()

    reHLL := regexp.MustCompile(`(?:L|l|D|d|H|w|Haut|Long|Larg|Diam|Longueur|Largeur|Hauteur|Prof|Profondeur) *(?:\.|\:)* *(\d+)`)
    reHLL.Longest()

    reOneD := regexp.MustCompile(`(\d+) (?:cm|CM)`)
    reOneD.Longest()

    return reX, reHLL, reOneD
}

var reX, reHLL, reOneD = prepareAllRegexp()

func Split(r rune) bool {
    return r == ' ' || r == 'x' || r == 'X' || r == '*' || r == '-' || r == ',' || r == '/'
}

func extractPrefixedDims(desc string) Dimension {
    matchs := reHLL.FindAllStringSubmatch(desc, -1)

    var d Dimension
    if len(matchs) == 0 {
        return d
    }

    var dims []int
    for _, match := range matchs {
        dimsStr := match[1]
        dim, _ := strconv.Atoi(dimsStr)
        dims = append(dims, dim)
    }

    sort.Slice(dims, func(i, j int) bool {
        return dims[i] > dims[j]
    })
    d.D1 = dims[0]
    if len(dims) > 1 {
        d.D2 = dims[1]
        if len(dims) > 2 {
            d.D3 = dims[2]
        }
    }
    return d
}

func extractXedDims(desc string) Dimension {
    matchs := reX.FindAllStringSubmatch(desc, -1)

    var d Dimension

    for _, match := range matchs {
        dimsStr := match[1]
        splittedDims := strings.FieldsFunc(dimsStr, Split)
        if len(splittedDims) < 2 {
            continue
        }
        ints := make([]int, len(splittedDims))
        for i, s := range splittedDims {
            ints[i], _ = strconv.Atoi(s)
        }
        sort.Slice(ints, func(i, j int) bool {
            return ints[i] > ints[j]
        })
        d.D1 = ints[0]
        d.D2 = ints[1]
        if len(ints) > 2 {
            d.D3 = ints[2]
        }
    }
    return d
}

func extractOneDim(desc string) Dimension {
    matchs := reOneD.FindStringSubmatch(desc)

    var d Dimension

    if len(matchs) == 0 {
        return d
    }
    d.D1, _ = strconv.Atoi(matchs[1])
    return d
}

func ExtractDims(desc string) (Dimension, error) {

    d := extractXedDims(desc)

    if d.D1 != 0 {
        return d, nil
    }

    d = extractPrefixedDims(desc)
    if d.D1 != 0 {
        return d, nil
    }

    return extractOneDim(desc), nil
}