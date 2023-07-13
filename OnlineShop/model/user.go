package model

import (
	"OnlineShop/model/request"
	"OnlineShop/model/response"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"log"
)

/*
		type User struct {
		TimeModel
		UserID         uint   `gorm:"primarykey"` // 使用 UserID 而非 UserId
		Name           string `gorm:"type:varchar(20);unique"`
		Nickname       string `gorm:"type:varchar(20)"`
		Email          string `gorm:"type:varchar(100);unique"`
		Tel            string `gorm:"type:varchar(20);unique"`
		HashedPassword string `gorm:"size:60"`
		IsAdmin        bool
		Addresses      []Address
		//Addresses []Address `gorm:"foreignKey:UserID"`
	}
*/
type User struct {
	TimeModel
	UserID         uint   `gorm:"primarykey"` // 使用 UserID 而非 UserId
	Name           string `gorm:"type:varchar(20);unique" validate:"required,max=20"`
	Nickname       string `gorm:"type:varchar(20)" validate:"omitempty,max=20"`
	Email          string `gorm:"type:varchar(100);unique" validate:"required,email,max=100"`
	Tel            string `gorm:"type:varchar(20);unique" validate:"omitempty,e164,max=20"`
	HashedPassword string `gorm:"size:60" validate:"required,len=60"`
	IsAdmin        bool
	Addresses      []Address
	//Addresses []Address `gorm:"foreignKey:UserID"`
	validate *validator.Validate `gorm:"-"`
}

func (u *User) Validate() error {
	if u.validate == nil {
		u.validate = validator.New()
	}
	return u.validate.Struct(u)
}

// Register 注册用户 - 调用查重函数和创建用户函数
func Register(userReq *request.UserRequest) (err error) {
	// 检验数据是否合规
	validate := validator.New()
	if err := validate.Struct(userReq); err != nil {
		return err
	}
	// 查询该用户是否注册过(名称/邮箱)
	if IsExistUser(userReq.Name, userReq.Email) {
		return errors.New("该用户名/邮箱已注册")
	}
	// 注册用户
	if err := CreateUser(userReq); err != nil {
		return err
	}
	return nil
}

// IsExistUser 检查用户名和邮箱是否注册过
func IsExistUser(name, email string) bool {
	var user User
	result := DBSql.Where("name = ? OR email = ?", name, email).First(&user)
	// 如果没有找到记录，返回 false
	// 静态检查建议：直接返回布尔值，不进行 if 判断
	return result.RowsAffected != 0
}

// CreateUser 创建用户
func CreateUser(userReq *request.UserRequest) (err error) {
	//user := Convertrequest.UserRequestToUser(userReq)
	// 改用 copier 进行结构体转换
	user := &User{}
	errCopier := copier.Copy(&user, &userReq)
	//log.Printf("[info]model-用户注册时,结构体转换%v\n===>%v\n", userReq, user)
	if errCopier != nil {
		log.Printf("[error]model-用户注册时,结构体转换错误\n")
		return errCopier
	}
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		log.Printf("[error]密码加密错误\n")
		return err
	}
	user.HashedPassword = string(hashedPassword)
	if err := DBSql.Create(user).Error; err != nil {
		log.Printf("[error]model-新建用户错误\n")

		return err
	}
	userReq.UserID = user.UserID
	return nil
}

// CheckUser 用户登录时 检查邮箱是否存在 邮箱密码是否一致 **是否是管理员**
func CheckUser(userSReq *request.UserSignInRequest) (userResponse *response.UserSignInResponse, err error) {
	// 检查邮箱是否存在
	var user User
	result := DBSql.Where("email = ?", userSReq.Email).First(&user)
	if result.RowsAffected == 0 {
		log.Printf("[info]model-邮箱不存在%v\n", userSReq.Email)
		return nil, errors.New("该邮箱未注册")
	}
	// 检查密码是否一致
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userSReq.Password)); err != nil {
		log.Printf("[info]model-密码不一致%v\n", userSReq.Email)
		return nil, errors.New("密码不一致")
	}
	log.Printf("[info]model-CheckUser,user:%v\n", user)
	userResponse = &response.UserSignInResponse{ID: user.UserID, IsAdmin: user.IsAdmin}
	return userResponse, nil
}

func UpdateUser(userReq *request.UserUpdateRequest) {
	user := &User{}
	user.UserID = userReq.UserID
	log.Printf("[info]model-UpdateUser,user:%v\n", user)
	if len(userReq.Password) != 0 {
		hashedPasswordByte, _ := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
		userReq.Password = string(hashedPasswordByte)
	}
	DBSql.Model(user).Updates(userReq)
}

// DeleteUser 删除用户
func DeleteUser(userId uint) (err error) {
	err = DBSql.Delete(&User{}, userId).Error
	if err != nil {
		log.Printf("[error]model-DeleteUser,%v\n", err.Error())
		return err
	}
	log.Printf("[info]model-DeleteUser,delete userid:%v\n", userId)

	return nil
}

// QueryOneUser 查询某个用户的信息
func QueryOneUser(userId uint) (*response.UserQueryResponse, error) {
	var userResponse response.UserQueryResponse
	// 查询单条记录要注意添加 Where 条件
	// 使用 Select 方法，避免进行模型间的转化
	if err := DBSql.Table("users").Where("id=?", userId).Select("name,nickname,email,tel,is_admin,id").First(&userResponse).Error; err != nil {
		log.Printf("[error]model-QueryOneUser,%v\n", err.Error())
		return nil, err
	}

	return &userResponse, nil
}
