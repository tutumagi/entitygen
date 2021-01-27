//go:generate go run entitygen entitygen/domain
package entitydef

import attr "entitygen/attr"

type IField interface {
	setParent(k string, parent attr.Field)
}
