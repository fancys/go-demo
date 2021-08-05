package main

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

func main() {
	material := Material{
		Id:         12356,
		ObjectCode: "MATERIAL",
		Code:       "1",
		Name:       "2",
		ManageType: "B",
		Status:     "Y",
		Creator:    1,
		CreateTime: time.Now()}

	fmt.Println(material.Code)

	dsn := "sqlserver://sa:AVAtech@2020@192.168.0.216:1433?database=dahupt_stocktask_BAT"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Add(&material, db)

	defer sqlDB.Close()
}

type Material struct {
	Id int64 `gorm "ids"`

	ObjectCode string `gorm object_code`

	Code string `gorm code`

	Name string `gorm name`

	ManageType string `gorm manage_type`

	Status string `gorm status`

	Creator int64 `gorm creator`

	CreateTime time.Time `gorm create_time`
}

func (Material) TableName() string {
	return "ava_mm_material"
}

func Add(material *Material, db *gorm.DB) {
	db.Create(material)
}

func BatchAdd(materials *[]Material, db *gorm.DB) {
	db.CreateInBatches(materials, 10)
}

//
//func Find(db *gorm.DB) []Material{
//	var material Material
// 	result := db.Find(&material)
//  	//return result.Rows()
//}
