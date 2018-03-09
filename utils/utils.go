package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gopkg.in/d4l3k/messagediff.v1"
	"mycmdb/models"
	"net"
	"net/smtp"
	"reflect"
	"strings"
	"sync"
)

var AvailIpLock *sync.Mutex

type MyOrmer struct {
	orm.Ormer
}

func MyNewOrm() *MyOrmer {
	o := orm.NewOrm()
	return &MyOrmer{o}
}

type SessionInterface interface {
	GetSession(interface{}) interface{}
}

func (this *MyOrmer) Update(c SessionInterface, x interface{}, s ...string) (int64, error) {
	o := orm.NewOrm()
	name, value, _ := GetPK(x)
	var diff string

	switch new := x.(type) {
	case *models.Idc:
		old := &models.Idc{Id: new.Id}
		o.Read(old)
		new.CreateTime = old.CreateTime
		diff = DiffObj(old, new)
	default:
		panic("Unknown model")
	}

	log := models.OperateLog{User: c.GetSession("name").(string),
		Action:     c.GetSession("action").(string),
		Url:        c.GetSession("url").(string),
		Model:      reflect.TypeOf(x).Elem().Name(),
		PrimaryKey: fmt.Sprintf("%v=%v", name, value),
		Detail:     diff}
	beego.Error(log)
	o.Insert(&log)

	return o.Update(x, s...)
}

func GetPK(x interface{}) (string, interface{}, error) {
	v := reflect.ValueOf(x).Elem()

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		// fmt.Println(i, v.Type().Field(i).Name, values[i], reflect.TypeOf(values[i]), v.Type().Field(i).Tag.Get("orm"))
		if v.Type().Field(i).Name == "Id" {
			return v.Type().Field(i).Name, value, nil
		} else if strings.Contains(v.Type().Field(i).Tag.Get("orm"), "pk") {
			return v.Type().Field(i).Name, value, nil
		}
	}

	return "", nil, errors.New("Not found primary key")
}

func DiffObj(origin, modify interface{}) string {
	diff, _ := messagediff.PrettyDiff(origin, modify)
	return diff
}

type ParseDataTableParamsInterface interface {
	GetInt(key string, def ...int) (int, error)
	GetString(key string, def ...string) string
}

func ParseDataTableParams(this ParseDataTableParamsInterface) (int, int, int, int, string) {
	secho, err := this.GetInt("sEcho")
	if err != nil {
		secho = 0
	}
	start, err := this.GetInt("iDisplayStart")
	if err != nil {
		start = 0
	}
	length, err := this.GetInt("iDisplayLength")
	if err != nil {
		length = 10
	}
	sort_th, err := this.GetInt("iSortCol_0")
	if err != nil {
		sort_th = 0
	}
	sort_type := this.GetString("sSortDir_0")
	if sort_type == "desc" {
		sort_type = "-"
	} else {
		sort_type = ""
	}

	return secho, start, length, sort_th, sort_type
}

func GetIPFromCIDR(cidr string) ([]string, string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	beego.Error(ip, ipnet, err)
	if err != nil {
		return nil, "", err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	beego.Error(ips)

	if len(ips) <= 2 {
		return []string{ip.String()}, ipnet.String(), nil
	}

	// remove network address and broadcast address
	return ips[1 : len(ips)-1], ipnet.String(), nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func SendMail(to, subject, body, mailtype string) error {
	user, password, host := "winway1988@163.com", "123456", "smtp.163.net:25"

	auth := smtp.PlainAuth("", user, password, strings.Split(host, ":")[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

//aes的加密字符串
var key_text = "12d7w948tahkvl1jzmknm.ahkjkljl;k"

func MyEncode(plaintext []byte) ([]byte, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		beego.Error("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		return []byte{}, err
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func MyDecode(ciphertext []byte) ([]byte, error) {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		beego.Error("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		return []byte{}, err
	}

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintext := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintext, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintext)

	return plaintext, nil
}

func GetAvailIp(idc int, ipType int, num int) ([]string, error) {
	ips := []string{}

	AvailIpLock.Lock()
	defer AvailIpLock.Unlock()

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Ip)).
		Filter("IpType", ipType).
		Filter("Status__ne", 1).
		RelatedSel("Idc").
		Filter("Idc__Id", idc).
		Offset(0).
		Limit(num)

	var maps []*models.Ip
	if n, err := qs.All(&maps, "Ip"); err != nil {
		beego.Error(err)
		return ips, err
	} else if int(n) < num {
		beego.Error("Not Enough IP")
		return ips, errors.New("Not Enough IP")
	}

	o.Begin()
	for _, i := range maps {
		i.Status = 1
		if _, err := o.Update(i, "Status"); err != nil {
			beego.Error(err)
			o.Rollback()
			return ips, err
		}
	}
	o.Commit()

	for _, v := range maps {
		ips = append(ips, v.Ip)
	}

	return ips, nil
}

func init() {
	AvailIpLock = new(sync.Mutex)
}
