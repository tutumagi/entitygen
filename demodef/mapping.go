//go:generate go run github.com/tutumagi/entitygen github.com/tutumagi/entitygen/domain
package demodef

import attr "github.com/tutumagi/entitygen/attr"

type IField interface {
	setParent(k string, parent attr.Field)
}
