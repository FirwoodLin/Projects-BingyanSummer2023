package model

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Register 注册用户 - 调用查重函数和创建用户函数
func Register(userReq *UserRequest) (err error) {
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
	result := DB.Where("name = ? OR email = ?", name, email).First(&user)
	// 如果没有找到记录，返回 false
	// 静态检查建议：直接返回布尔值，不进行 if 判断
	return result.RowsAffected != 0
}

// CreateUser 创建用户
func CreateUser(userReq *UserRequest) (err error) {
	user := ConvertUserRequestToUser(userReq)
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		log.Printf("[error]密码加密错误\n")
		return err
	}
	user.HashedPassword = string(hashedPassword)
	if err := DB.Create(user).Error; err != nil {
		log.Printf("[error]model-新建用户错误\n")

		return err
	}
	userReq.ID = user.ID
	return nil
}

// CheckUser 用户登录时 检查邮箱是否存在 邮箱密码是否一致 **是否是管理员**
func CheckUser(userSReq *UserSignInRequest) (userResponse *UserSignInResponse, err error) {
	// 检查邮箱是否存在
	var user User
	// user = new(User)
	// user.Email = userSReq.Email
	result := DB.Where("email = ?", userSReq.Email).First(&user)
	if result.RowsAffected == 0 {
		log.Printf("[info]model-邮箱不存在%v\n", userSReq.Email)
		return nil, errors.New("该邮箱未注册")
	}
	// userSReq.ID = user.ID
	// 检查密码是否一致
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(userSReq.Password)); err != nil {
		log.Printf("[info]model-密码不一致%v\n", userSReq.Email)
		return nil, errors.New("密码不一致")
	}
	log.Printf("[info]model-CheckUser,user:%v\n", user)
	userResponse = &UserSignInResponse{ID: user.ID, IsAdmin: user.IsAdmin}
	return userResponse, nil
}

func UpdateUser(userReq *UserUpdateRequest) {
	// user := ConvertUserUpdateRequestToUser(userReq)
	user := &User{}
	user.ID = userReq.ID
	log.Printf("[info]model-UpdateUser,user:%v\n", user)
	// user.Name = userReq.Name
	// user.Nickname = userReq.Nickname
	// user.Email = userReq.Email
	// user.Tel = userReq.Tel
	if len(userReq.Password) != 0 {
		hashedPasswordByte, _ := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
		userReq.Password = string(hashedPasswordByte)
	}
	DB.Model(user).Updates(userReq)
}

// DeleteUser 删除用户
func DeleteUser(userId uint) (err error) {
	err = DB.Delete(&User{}, userId).Error
	if err != nil {
		log.Printf("[error]model-DeleteUser,%v\n", err.Error())
		return err
	}
	log.Printf("[info]model-DeleteUser,delete userid:%v\n", userId)

	return nil
}

// QueryOneUser 查询某个用户的信息
func QueryOneUser(userId uint) (*UserQueryResponse, error) {
	var userResponse UserQueryResponse
	// 查询单条记录要注意添加 Where 条件
	// 使用 Select 方法，避免进行模型间的转化
	if err := DB.Table("users").Where("id=?", userId).Select("name,nickname,email,tel,is_admin,id").First(&userResponse).Error; err != nil {
		log.Printf("[error]model-QueryOneUser,%v\n", err.Error())
		return nil, err
	}

	return &userResponse, nil
}

// QueryUsers 查询所有人的信息
func QueryUsers() (*[]UserQueryResponse, error) {
	var users []UserQueryResponse
	//if err := DB.Find(&users).Error; err != nil {
	//	log.Printf("[error]model-QueryUsers,%v", err.Error())
	//	return nil, err
	//}
	if err := DB.Table("users").Select("name,nickname,email,tel,is_admin,id").Scan(&users).Error; err != nil {
		log.Printf("[error]model-QueryUsers,%v", err.Error())
		return nil, err
	}
	return &users, nil
}
