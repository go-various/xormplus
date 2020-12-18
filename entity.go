package xormplus

var _ EntityInc = (*Entity)(nil)


type EntityInc interface {
	ID() interface{}
	Detail() (bool, error)
	Exists() (bool, error)
	Create() (int64, error)
	Update() (int64, error)
	Remove() (int64, error)
	Object() interface{}
}

type Entity struct {
	dao     XormPlus
	beanPtr interface{}
}

func NewEntity(dao XormPlus, objectPtr interface{}) *Entity {
	return &Entity{dao: dao, beanPtr: objectPtr}
}

func (e *Entity) Object() interface{} {
	return e.beanPtr
}

func (e *Entity) ID() interface{} {
	return e.beanPtr.(Table).ID()
}

func (e *Entity) Exists() (bool, error) {
	return e.dao.Exist(e.beanPtr)
}

func (e *Entity) Detail() (has bool, err error) {
	return e.dao.ID(e.beanPtr.(Table).ID()).Get(e.beanPtr)
}

func (e *Entity) Create() (int64, error) {
	return e.dao.InsertOne(e.beanPtr)
}

func (e *Entity) Update() (int64, error) {
	return e.dao.ID(e.beanPtr.(Table).ID()).Update(e.beanPtr)
}

func (e *Entity) Remove() (rows int64, err error) {
	return e.dao.ID(e.beanPtr.(Table).ID()).Delete( e.beanPtr)
}
