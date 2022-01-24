//go:generate go run gitlab.testkaka.com/usm/game/entitygen gitlab.testkaka.com/usm/game/entitygen/domain
package demodef

import attr "gitlab.testkaka.com/usm/game/entitygen/attr"

type IField interface {
	setParent(k string, parent attr.Field)
}
