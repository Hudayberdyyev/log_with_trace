package domain

type Phrase struct {
	phrase string
	author string
}

func NewPhrase(phrase, author string) *Phrase {
	return &Phrase{
		phrase: phrase,
		author: author,
	}
}

func (p *Phrase) GetPhrase() string {
	return p.phrase
}

func (p *Phrase) GetAuthor() string {
	return p.author
}
