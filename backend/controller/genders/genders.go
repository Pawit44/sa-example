package genders

import (
   "net/http"
   "example.com/sa-example/config"
   "example.com/sa-example/entity"
   "github.com/gin-gonic/gin"
)

// ฟังก์ชันนี้เป็น handler สำหรับ route เช่น GET /genders โดยรับ context c จาก Gin
func GetAll(c *gin.Context) {
   db := config.DB()
   var genders []entity.Genders // ประกาศ slice ของ Genders struct เพื่อเก็บข้อมูลทั้งหมดจาก DB
   db.Find(&genders) // ใช้ GORM ดึงข้อมูลทั้งหมดจากตาราง genders (select all) แล้วใส่ใน genders slice
   // GORM จะทำการ SELECT * FROM genders แล้ว map ผลลัพธ์แต่ละแถวใส่ลงใน slice genders ทีละตัว

   c.JSON(http.StatusOK, &genders) // ส่ง response กลับไปเป็น JSON พร้อม HTTP Status 200 OK
}


// ตัวอย่าง JSON ที่จะได้ ไปเก็บไว้ใน slice genders
// [
//   {
//     "ID": 1,
//     "Gender": "Male"
//   },
//   {
//     "ID": 2,
//     "Gender": "Female"
//   }
// ]