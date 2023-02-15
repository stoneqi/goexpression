package src

import "errors"

type MapParameters map[string]any

func (p MapParameters) Get(name string) (any, error) {

	value, found := p[name]

	if !found {
		errorMessage := "No parameter '" + name + "' found."
		return nil, errors.New(errorMessage)
	}

	return value, nil
}
