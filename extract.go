package dimextract

import (
    "regexp"
)

type Dimension struct {
    height int,
    width  int,
    length int,
}

reA := regexp.MustCompile(`(\d+)[[:blank:]+|[:blank:]*cm[:blank:]*]`)

reB := regexp.MustCompile(`a(x*)b(y|z)c`)

func ExtractDims(desc string) (Dimension, error) {
    
    m := reA.FindStringSubmatch(desc)
    
    if len(m) > 0 {
        
    }
}