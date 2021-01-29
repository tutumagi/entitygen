//go:generate go run gitlab.gamesword.com/nut/entitygen gitlab.gamesword.com/nut/entitygen/domain
package entitydef

import attr "gitlab.gamesword.com/nut/entitygen/attr"

type IField interface {
	setParent(k string, parent attr.Field)
}
