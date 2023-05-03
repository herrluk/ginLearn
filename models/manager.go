package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   int
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
}

// TableName 实现 tableName()方法后，不需要任何操作就会自动将迁移的表名改为 manager
// 默认是 struct名 + s
func (Manager) TableName() string {
	return "manager"
}
