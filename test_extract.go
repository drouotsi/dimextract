package dimextract

import "testing"

type TestCase {
    desc string
    dim Dimension
}

func TestExtract(t *testing.T) {
    testCases := []TestCase{
        {
            "Blabla bla 13 * 12 * 11 cm",
            {
                13, 12, 11
            }
        },
        {
            "Blabla bla 13 * 12 * 11 cm",
            {
                13, 12, 11
            }
        },
    }
    
    for _, tc := range testCases {
        d, err := ExtractDims(tc.desc)
        if err != nil {
            t.Errorf("Eeeee", err)
        }
        if d != tc.dim {
            t.Error("Fail !")
        }
    }
}