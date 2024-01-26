package server

type Clue struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Indicator struct {
    Word     string `json:"word"`
    ClueType *Clue  `json:"clue_type"`
}

type Indicators struct {
    Indicators []*Indicator `json:"indicators"`
}

// TEMP CODE TO CHECK WORKS
func (i *Indicators) Search(s string) *Indicators {
    if s == "askew" {
        return &Indicators{Indicators: []*Indicator{&Indicator{
            Word: "askew",
            ClueType: &Clue{
                Name:        "Anagram",
                Description: "An anagram is wordplay in which letters from a part of the clue are reordered to give the solution.",
            },
        }}}
    }
    return nil
}
