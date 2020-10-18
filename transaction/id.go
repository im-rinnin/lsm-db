package transaction

import "github.com/PatchouliG/lsm-db/id"

var idGenerator *id.Generator

type Id struct {
	id.Id
}

func NextId() Id {
	return Id{idGenerator.Next()}
}

func (id *Id) Cmp(i Id) int64 {
	return id.Id.Cmp(i.Id)
}
