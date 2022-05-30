package repositories

import "gorm.io/gorm"

type User struct {
	ID       int
	Username string `gorm:"unique"`
	Password string
	Email    string
	LineID   string
}

type UpdateUserInformation struct {
	Username string
	Email    string
	LineID   string
}

type UpdateUserPassword struct {
	Username string
	Password string
}

type UserRepository interface {
	Create(User) error
	UpdateInformation(UpdateUserInformation) error
	UpdatePassword(UpdateUserPassword) error
	GetUsers() ([]User, error)
	GetUserByUsername(username string) (User, error)
	Delete(username string) error
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepostoryDB(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return userRepositoryDB{db}
}

func (r userRepositoryDB) Create(user User) error {
	return r.db.Create(&user).Error
}

func (r userRepositoryDB) UpdateInformation(userUpdate UpdateUserInformation) error {
	return r.db.Model(&User{}).Where("username=?", userUpdate.Username).Updates(User{
		Email:  userUpdate.Email,
		LineID: userUpdate.LineID,
	}).Error
}

func (r userRepositoryDB) UpdatePassword(userUpdate UpdateUserPassword) error {
	return r.db.Model(&User{}).Where("username=?", userUpdate.Username).Updates(User{
		Password: userUpdate.Password,
	}).Error
}

func (r userRepositoryDB) GetUsers() (users []User, err error) {
	err = r.db.Find(&users).Error
	return users, err
}

func (r userRepositoryDB) GetUserByUsername(username string) (user User, err error) {
	err = r.db.Where("username=?", username).First(&user).Error
	return user, err
}

func (r userRepositoryDB) Delete(username string) error {
	return r.db.Where("username=?", username).Delete(&User{}).Error
}
