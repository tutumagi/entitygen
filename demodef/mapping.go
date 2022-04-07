//go:generate go run gitlab.nftgaga.com/usm/game/entitygen gitlab.nftgaga.com/usm/game/entitygen/domain
package demodef

import attr "gitlab.nftgaga.com/usm/game/entitygen/attr"

type IField interface {
	setParent(k string, parent attr.Field)
}
