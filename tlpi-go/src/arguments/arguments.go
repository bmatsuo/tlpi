package arguments

import (
	"fmt"
	"os"
	"strconv"
	"regexp"
)

type InvalidArgumentError struct {
	arg     string
	message string
}

func (err InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid argument: %v -- %v", err.arg, err.message)
}

type MissingArgumentError struct {
	arg string
}

func (err MissingArgumentError) Error() string {
	return fmt.Sprintf("missing argument: %v", err.arg)
}

type Args []string

func Arguments() Args {
	return os.Args[1:]
}

func (as Args) Exist() bool {
	return len(as) > 0
}

func (as Args) String(name string, pattern *regexp.Regexp) (string, Args, error) {
	if !as.Exist() {
		return "", as, MissingArgumentError{name}
	}

	a, _as := as[0], as[1:]

	if pattern != nil{
		if !pattern.MatchString(a) {
			return "", as, InvalidArgumentError{
				name,
				fmt.Sprintf("must have the form %v", pattern),
			}
		}
	}

	return a, _as, nil
}

func (as Args) Int(name string, min, max int) (int, Args, error) {
	if !as.Exist() {
		return 0, as, MissingArgumentError{name}
	}

	a, _as := as[0], as[1:]
	x, err := strconv.ParseInt(a, 10, 0)

	if err != nil {
		return 0, as, err
	} else if x < int64(min) {
		return 0, as, InvalidArgumentError{
			name,
			fmt.Sprintf("minimum value exceeded (%d)", min),
		}
	} else if x > int64(max) {
		return 0, as, InvalidArgumentError{
			name,
			fmt.Sprintf("maximum value exceeded (%d)", min),
		}
	}

	return int(x), _as, err
}

func (as Args) Int64(name string, min, max int64) (int64, Args, error) {
	if !as.Exist() {
		return 0, nil, MissingArgumentError{name}
	}

	a, as := as[0], as[1:]
	x, err := strconv.ParseInt(a, 10, 64)

	if err != nil {
		return 0, as, err
	} else if x < min {
		return 0, as, InvalidArgumentError{
			name,
			fmt.Sprintf("minimum value exceeded (%d)", min),
		}
	} else if x > max {
		return 0, as, InvalidArgumentError{
			name,
			fmt.Sprintf("maximum value exceeded (%d)", min),
		}
	}

	return x, as, err
}
