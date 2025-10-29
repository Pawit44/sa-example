package config

import (
   "fmt"
   "time"
   "example.com/sa-example/entity"
   "gorm.io/driver/sqlite"
   "gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
   return db
}

func ConnectionDB() {
   database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
    if err != nil {
        panic(fmt.Sprintf("failed to connect database: %v", err))
    }
   fmt.Println("connected database")
   db = database
}

func SetupDatabase() {
   db.AutoMigrate(
       &entity.Users{},
       &entity.Genders{},
   )
   
   // สร้างตัวแปร struct Genders จำนวน 2 ตัว แทนเพศชายและหญิง
   GenderMale := entity.Genders{Gender: "Male"}
   GenderFemale := entity.Genders{Gender: "Female"}

    // เช็คในตาราง genders ว่ามีแถวไหนที่ Gender = "Male" แล้วหรือยัง ถ้ายัง ไม่มี → จะ สร้าง แถวใหม่ (Insert) ถ้ามีอยู่แล้ว → จะ ไม่สร้างซ้ำ (เหมือน SELECT)
   db.FirstOrCreate(&GenderMale, &entity.Genders{Gender: "Male"})
    // เช็คในตาราง genders ว่ามีแถวไหนที่ Gender = "Female" แล้วหรือยัง ถ้ายัง ไม่มี → จะ สร้าง แถวใหม่ (Insert) ถ้ามีอยู่แล้ว → จะ ไม่สร้างซ้ำ (เหมือน SELECT)
   db.FirstOrCreate(&GenderFemale, &entity.Genders{Gender: "Female"})

   // สร้าง "บัญชีผู้ใช้เริ่มต้น"
   hashedPassword, _ := HashPassword("123456")
   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12") // แปลง string "1988-11-12" ให้กลายเป็นค่า time.Time ที่ Go ใช้ได้ รูปแบบ 2006-01-02 คือ format พิเศษของ Go (ยึดปี 2006 เป็นแม่แบบ)
   User := &entity.Users{
       FirstName: "Software",
       LastName:  "Analysis",
       Email:     "sa@gmail.com",
       Age:       80,
       Password:  hashedPassword,
       BirthDay:  BirthDay,
       GenderID:  1,
       Address: "F11",
   }
   db.FirstOrCreate(User, &entity.Users{ //ถ้ามีผู้ใช้ที่ Email = "sa@gmail.com" อยู่แล้ว → จะไม่ insert ซ้ำ ถ้ายังไม่มี → จะ insert ผู้ใช้นี้ลงฐานข้อมูล
       Email: "sa@gmail.com", // email คือ 
   })
}