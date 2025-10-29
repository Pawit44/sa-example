package users

import (
   "errors"
   "net/http"
   "time"
   "github.com/gin-gonic/gin"
   "golang.org/x/crypto/bcrypt"
   "gorm.io/gorm"
   "example.com/sa-example/config"
   "example.com/sa-example/entity"
   "example.com/sa-example/services"
)

type (
   Authen struct {
       Email    string `json:"email"`
       Password string `json:"password"`
   }
   
   signUp struct {
       FirstName string    `json:"first_name"`
       LastName  string    `json:"last_name"`
       Email     string    `json:"email"`
       Age       uint8     `json:"age"`
       Password  string    `json:"password"`
       BirthDay  time.Time `json:"birthday"`
       GenderID  uint      `json:"gender_id"`
       Address  string     `json:"address"`
   }
)


func SignUp(c *gin.Context) {
   var payload signUp
   // Bind JSON payload to the struct
   // รับข้อมูล JSON จากฝั่ง client (เช่น POST หรือ PUT) แล้ว แปลง (bind) JSON นั้นให้เข้ากับ struct ที่ชื่อว่า payload แล้วเอาใส่เข้าไปใน payload ทั้งหมด 
   if err := c.ShouldBindJSON(&payload); err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }

   db := config.DB()
   var userCheck entity.Users

   // Check if the user with the provided email already exists
   result := db.Where("email = ?", payload.Email).First(&userCheck)// “ตรวจสอบว่าอีเมลที่ client ส่งเข้ามามีอยู่ในฐานข้อมูลหรือไม่” และถ้ามี → จะดึงข้อมูลมาเก็บไว้ในตัวแปร userCheck แต่ถ้าไม่มีจะเป็นเหมือนการ select จะไม่ถูกเก็บลงใน userCheck
   if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
       // If there's a database error other than "record not found" หากมีข้อผิดพลาดฐานข้อมูลอื่นนอกเหนือจาก "ไม่พบระเบียน"
       // กรณี error จริง (เช่น DB ล่ม)
       c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
       return
   }

   if userCheck.ID != 0 { // ถ้า ID != 0 แปลว่า มีข้อมูล user อยู่ใน DB แล้ว (เพราะถ้าไม่มีข้อมูล userCheck จะมีค่า default ของ ID เป็น 0)
       // If the user with the provided email already exists
       c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
       return
   }

   // Hash the user's password
   hashedPassword, _ := config.HashPassword(payload.Password)
   // Create a new user
   user := entity.Users{
       FirstName: payload.FirstName,
       LastName:  payload.LastName,
       Email:     payload.Email,
       Age:       payload.Age,
       Password:  hashedPassword,
       BirthDay:  payload.BirthDay,
       GenderID:  payload.GenderID,
       Address:  payload.Address,
   }
    // Save the user to the database
    // db.Create(&user) → พยายาม INSERT ข้อมูลจาก struct user เข้าไปในตาราง → ถ้า user ตรงกับ model Users ก็จะ insert ลงตารางนั้นเลย
    //.Error → GORM จะคืนค่า struct *gorm.DB ที่มี field Error → ถ้าเกิด error ตอน insert เช่น primary key ซ้ำ หรือ constraint violation มันจะไม่เป็น nil
   if err := db.Create(&user).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }
   c.JSON(http.StatusCreated, gin.H{"message": "Sign-up successful"})
}

func SignIn(c *gin.Context) {
   var payload Authen
   var user entity.Users
   if err := c.ShouldBindJSON(&payload); err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }
   // ค้นหา user ด้วย Username ที่ผู้ใช้กรอกเข้ามา
   if err := config.DB().Raw("SELECT * FROM users WHERE email = ?", payload.Email).Scan(&user).Error; err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
       return
   }

   // ตรวจสอบรหัสผ่าน
   err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
       return
   }
   jwtWrapper := services.JwtWrapper{
       SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
       Issuer:          "AuthService",
       ExpirationHours: 24,
   }

   signedToken, err := jwtWrapper.GenerateToken(user.Email)

   if err != nil {
       c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
       return
   }
   c.JSON(http.StatusOK, gin.H{"token_type": "Bearer", "token": signedToken, "id": user.ID})
}