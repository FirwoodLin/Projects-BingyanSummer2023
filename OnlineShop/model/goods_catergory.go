package model

type Category struct {
	TimeModel
	CategoryID uint   `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100)"`
	Weight     int    // 权重，用于排序
	// 进行外键关联 = 暂时取消外键关联
	ParentID uint
	// 0 代表没有父类，即为根类
	//Parent *Category
}

// GetRootCategories 查询所有的根类，并返回
func GetRootCategories() (categories []Category) {
	DBSql.Where("parent_id = ?", 0).Find(&categories)
	return
}

// GetAllCategories 查询所有的分类，并返回
func GetAllCategories() (categories []Category, err error) {
	err = DBSql.Find(&categories).Error
	return
}

// GetCategoryByID 根据 ID 查询分类，并返回
func GetCategoryByID(id uint) (category Category, err error) {
	err = DBSql.Find(&category, id).Error
	// golang 可以在定义函数时，指定返回值的名称，这样可以直接 return
	return
}
