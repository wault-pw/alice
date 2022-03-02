package validator

import "github.com/oka-is/alice/pkg/domain"

type IValidator interface {
	ValidateUser(user domain.User) error
	ValidateTerminate(opts ValidateTerminateOpts) error
}

type Validator struct{}
type NoOptValidator struct{}

func New() *Validator {
	return &Validator{}
}

func NewNoOpt() *NoOptValidator {
	return &NoOptValidator{}
}

func one(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}

	return nil
}
