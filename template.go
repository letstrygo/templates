package templates

import (
	"errors"

	"github.com/samber/lo"
)

var (
	ErrInvalidTemplateType = errors.New("invalid template type")
)

type TemplateType string

const (
	TemplateTypeGitRepository TemplateType = "git"
	TemplateTypeLocal         TemplateType = "local"
)

var TemplateTypes = []TemplateType{
	TemplateTypeGitRepository,
	TemplateTypeLocal,
}

func NewTemplateType(val string) (TemplateType, error) {
	var zeroVal TemplateType

	tt := TemplateType(val)
	if !lo.Contains(TemplateTypes, tt) {
		return zeroVal, ErrInvalidTemplateType
	}

	return tt, nil
}

type Template struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Source     string       `json:"source"`
	Type       TemplateType `json:"type"`
	IsOfficial bool         `json:"-"`
}
