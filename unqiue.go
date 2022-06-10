package unique

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 11:11
 * @Desc:
 */

type Unique struct {
}

// Time64 返回当前时间的64位时间戳
func (u *Unique) Time64() int64 {
	return time.Now().UnixNano() / 1e6
}

// Time32 返回当前时间的32位时间戳
func (u *Unique) Time32() int32 {
	return int32(time.Now().Unix())
}

// getFormatTime 获取格式化时间
func (u *Unique) getFormatTime(time time.Time) string {
	return time.Format("20060102")
}

// GetUniqueCodeByTime 日期20191025时间戳1571987125435+3位随机数
func (u *Unique) GetUniqueCodeByTime() string {
	date := u.getFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d03%d", date, u.Time64(), r)
	return code
}

// GetUniqueCodeBySnowflake 返回雪花算法的唯一ID
func (u *Unique) GetUniqueCodeBySnowflake() snowflake.ID {
	node, _ := snowflake.NewNode(1)
	id := node.Generate()
	return id
}
