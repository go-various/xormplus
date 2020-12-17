package xormplus

var _ Entity = (*entity)(nil)


type Entity interface {
	Id() interface{}
	Detail() (bool, error)
	Exists() (bool, error)
	Create() (int64, error)
	Update() (int64, error)
	Remove() (int64, error)
	Entity() interface{}
}

type entity struct {
	dao     XormPlus
	beanPtr interface{}
}

func NewEntity(dao XormPlus, objectPtr interface{}) *entity {
	return &entity{dao: dao, beanPtr: objectPtr}
}

func (e *entity) Entity() interface{} {
	return e.beanPtr
}

func (e *entity) Id() interface{} {
	return e.beanPtr.(Table).ID()
}

func (e *entity) Exists() (bool, error) {
	return e.dao.Exist(e.beanPtr)
}

func (e *entity) Detail() (has bool, err error) {
	return e.dao.ID(e.beanPtr.(Table).ID()).Get(e.beanPtr)
}

func (e *entity) Create() (int64, error) {
	return e.dao.InsertOne(e.beanPtr)
}

func (e *entity) Update() (int64, error) {
	return e.dao.ID(e.beanPtr.(Table).ID()).Update(e.beanPtr)
}

func (e *entity) Remove() (rows int64, err error) {
	return e.dao.ID(e.beanPtr.(Table).ID()).Delete( e.beanPtr)
}
