package pusuparams

import (
	"github.com/nickwells/check.mod/v2/check"
	"github.com/nickwells/param.mod/v6/psetter"
	"github.com/nickwells/pusu.mod/pusu"
)

// TopicSetter is a param.Setter which can be used to set pusu.Topic values
type TopicSetter struct {
	psetter.ValueReqMandatory

	// Value is a pointer to the topic that this setter will set
	Value *pusu.Topic
	// Checks is a slice of checks that the topic must pass
	Checks []check.ValCk[pusu.Topic]
}

// CountChecks returns the number of check functions this setter has
func (s TopicSetter) CountChecks() int {
	return len(s.Checks)
}

// SetWithVal (called with the value following the parameter) checks the .
func (s TopicSetter) SetWithVal(_ string, paramVal string) error {
	topic := pusu.Topic(paramVal)

	if err := topic.Check(); err != nil {
		return err
	}

	for _, check := range s.Checks {
		if err := check(topic); err != nil {
			return err
		}
	}

	*s.Value = topic

	return nil
}

// AllowedValues returns a string describing the allowed values
func (s TopicSetter) AllowedValues() string {
	return "A valid publish/subscribe topic" + psetter.HasChecks(s)
}

// ValDescribe returns a string describing the value that can follow the
// parameter
func (s TopicSetter) ValDescribe() string {
	return "topic"
}

// CurrentValue returns the current setting of the parameter value
func (s TopicSetter) CurrentValue() string {
	return string(*s.Value)
}

// CheckSetter panics if the setter has not been properly created - if the
// Value is nil.
func (s TopicSetter) CheckSetter(name string) {
	intro := name + ": pusuparam.TopicSetter Check failed:"

	if s.Value == nil {
		panic(intro + " the Value to be set is nil")
	}
}
