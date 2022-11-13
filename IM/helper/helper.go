package helper

import (
	"IM/define"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

type UserClaims struct {
	//Identity primitive.ObjectID `json:"identity"`
	Identity string `json:"identity"`

	Email     string `json:"email"`
	jwt.StandardClaims
}
//发送验证码
func MailSendCode(mail,code string)error {

	m := gomail.NewMessage()
	//发送人1836447715@qq.com
	m.SetHeader("From", "1836447715@qq.com")
	//接收人
	//m.SetHeader("To", "3207512353@qq.com")
	m.SetHeader("To", mail)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "那些年的风花雪月__By gongZhaowei")
	//内容
	m.SetBody("text/html", "<h1>"+code+"</h1>")
	//附件
	//m.Attach("./myIpPic.png")
	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "1836447715@qq.com", define.Mialpassword)

	//d := gomail.NewDialer("smtp.qq.com", 587, "1836447715@qq.com", code)
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		return err
	}
	//fmt.Printf("send mail success\n")
	return nil
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// GenerateToken
// 生成 token
func GenerateToken(identity, email string) (string, error) {
	//objectID,err:=primitive.ObjectIDFromHex(identity)
	UserClaim := &UserClaims{
		Identity:       identity,
		Email:email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}


//var r *rand.Rand
//
//func init() {
//	r = rand.New(rand.NewSource(time.Now().Unix()))
//}
//
//// RandString 生成随机字符串
//func RandString(len int) string {
//	bytes := make([]byte, len)
//	for i := 0; i < len; i++ {
//		b := r.Intn(26) + 65
//		bytes[i] = byte(b)
//	}
//	return string(bytes)
//}

func RandName() string{
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	len:=5
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return (string(bytes))
}
func RandInt()int32{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	 return r.Int31()
}
func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix()) //随机种子
	return rand.Intn(max - min) + min
}
