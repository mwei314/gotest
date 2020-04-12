package models

// User 用户表Model
type User struct {
	ID   uint   `gorm:"column:id;primary_key;auto_increment"`
	Name string `gorm:"column:name;type:varchar(100);not null;default:''"`
	Pass string `gorm:"column:pass;type:varchar(100);not null"`
}

// TableName 设置表名
func (User) TableName() string {
	return "tbl_user"
}

// Users 用户列表
type Users []User

// Insert 新增用户
func (user *User) Insert() error {
	result := GormDB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	user = result.Value.(*User)
	return nil
}

// List 获取用户列表
func (user User) List() {

}

// Update 修改用户
func (user User) Update() error {
	result := GormDB.Update(&user)
	return result.Error
}

// FindUserByName 根据名称查找用户
func FindUserByName(userName string) User {
	var user User
	GormDB.Where(&User{Name: userName}).First(&user)
	return user
}
