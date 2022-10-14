package error

// Error - define Error as constant
type Error string

func (re Error) Error() string {
	return string(re)
}
