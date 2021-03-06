package scago

// Scago is the base object that contains enough information to
// allow sound changes to be performed on words. It contains
// rules to be applied and the categories that are used in the
// rules.
type Scago struct {
	rules      *Rule     // a pointer to the first rule in the list
	categories *Category // a pointer to the first category in the list
}

// Apply applies the Scago's ruleset to the given word, returning
// the changed word and any error that came up. If an error is
// returned, the returned string may be empty.
// TODO: implement this functionally
func (s *Scago) Apply(lemma string) (string, error) {
	var err error
	for r := s.rules; r != nil; r = r.next {
		lemma, err = r.Apply(lemma)
		if err != nil {
			return "", err
		}
	}
	return lemma, nil
}

// New returns a new blank instance of Scago.
func New() *Scago {
	return &Scago{}
}
