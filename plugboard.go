package enigma

import (
	"errors"
)

type Steckerbrett Plugboard

type Plugboard struct {
	Mapping string
}

func (p *Plugboard) SetReciprocalMapping(mapping string) error {
	settings := Sanitize(mapping)
	final := make([]rune, 26)

	if len(settings)%2 != 0 {
		return errors.New("ERROR: Mapping must be pairs of A-Z characters")
	}

	for i := 0; i < len(settings)-1; i += 2 {
		a := rune(settings[i])
		b := rune(settings[i+1])
		if final[a-'A'] != 0 || final[b-'A'] != 0 {
			return errors.New("ERROR: Mapping must be distinct")
		} else {
			final[a-'A'] = b
			final[b-'A'] = a
		}
	}

	for i, v := range final {
		if v == 0 {
			final[i] = rune(i + 'A')
		}
	}

	p.Mapping = string(final)
	return nil
}
