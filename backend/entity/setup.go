package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("ระบบโปรแกรมบริหารจัดการสินค้าคงคลัง.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} 

	// Migrate the schema
	database.AutoMigrate(
		&Employee{},
		&Customer{},
		&Product{},
		&Pcode{},
		&Detail{},
		&Ccode{},
		&Cusdetail{},
	)

	db = database

	passwordEmployee1, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	Employee1 := Employee{
		FirstName:     "thna",
		LastName:	"srimeang",
		Email:    "thana@gmail.com",
		Password: string(passwordEmployee1),
	}
	db.Raw("SELECT * FROM employees WHERE email = ?", "thanaponkhanoon1123@gmail.com").Scan(&Employee1)
	passwordEmployee2, err := bcrypt.GenerateFromPassword([]byte("1234512126"), 14)
	Employee2 := Employee{
		FirstName:     "thanathon",
		LastName:	"pongpak",
		Email:    "thanathon@gmail.com",
		Password: string(passwordEmployee2),
	}
	db.Raw("SELECT * FROM employees WHERE email = ?", "thanaponkhanoon1123@gmail.com").Scan(&Employee2)
	passwordEmployee3, err := bcrypt.GenerateFromPassword([]byte("123426"), 14)
	Employee3 := Employee{
		FirstName:     "sukda",
		LastName:	"mama",
		Email:    "sakda@gmail.com",
		Password: string(passwordEmployee3),
	}
	db.Raw("SELECT * FROM employees WHERE email = ?", "thanaponkhanoon1123@gmail.com").Scan(&Employee3)


	//  Pcode
	id1235461 := Pcode{
		Name:			"id1234561",
	}
	db.Model(&Pcode{}).Create(&id1235461)

	id1235462 := Pcode{
		Name:			"id1234562",
	}
	db.Model(&Pcode{}).Create(&id1235462)

	id1235463 := Pcode{
		Name:			"id1234563",
	}
	db.Model(&Pcode{}).Create(&id1235463)

	// Detail
	fishnet := Detail{
		Name: "fishnet",
	}
	db.Model(&Detail{}).Create(&fishnet)

	net := Detail{
		Name: "net",
	}
	db.Model(&Detail{}).Create(&net)

	stew := Detail{
		Name: "stew",
	}
	db.Model(&Detail{}).Create(&stew)

	// Product
	db.Model(&Product{}).Create(&Product{
		Name: "pronpan",
		Price: 999,
		Amount: 999,
		Pcode: id1235461,
		Detail: fishnet,
		Time: time.Now(),
	})
	
	db.Model(&Product{}).Create(&Product{
		Name: "Thongchai",
		Price: 9999,
		Amount: 9999,
		Pcode: id1235462,
		Detail: net,
		Time: time.Now(),
	})

	db.Model(&Product{}).Create(&Product{
		Name: "Sutthiphong",
		Price: 99999,
		Amount: 99999,
		Pcode: id1235463,
		Detail: stew,
		Time: time.Now(),
	})

	// Ccode
	id12354671 := Ccode{
		Code: "id12345671",
	}
	db.Model(&Ccode{}).Create(&id12354671)

	id12354672 := Ccode{
		Code: "id12345672",
	}
	db.Model(&Ccode{}).Create(&id12354672)

	id12354673 := Ccode{
		Code: "id12345673",
	}
	db.Model(&Ccode{}).Create(&id12354673)

	// Cusdetail
	Moueang := Cusdetail{
		Cdname: "Moueang",
	}
	db.Model(&Cusdetail{}).Create(&Moueang)

	Kranuan := Cusdetail{
		Cdname: "Kranuan",
	}
	db.Model(&Cusdetail{}).Create(&Kranuan)

	PhuWiang := Cusdetail{
		Cdname: "PhuWiang",
	}
	db.Model(&Cusdetail{}).Create(&PhuWiang)

	// Customer
	db.Model(&Customer{}).Create(&Customer{
		Name:	"Thana",
		Ccode: 	id12354671,
		Cusdetail:	Moueang,
	})

	db.Model(&Customer{}).Create(&Customer{
		Name:	"Phongpad",
		Ccode: 	id12354672,
		Cusdetail:	Kranuan,
	})

	db.Model(&Customer{}).Create(&Customer{
		Name:	"Malee",
		Ccode: 	id12354673,
		Cusdetail:	PhuWiang,
	})
}