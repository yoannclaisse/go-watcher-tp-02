package checker

import "fmt"

type UnreachableURLError struct {
	URL string
	Err error
}

func (e *UnreachableURLError) Error() string {
	return fmt.Sprintf("URL innacessible : %s (%v)", e.URL, e.Err)
}

func (e *UnreachableURLError) Unwrap() error {
	return e.Err
}
