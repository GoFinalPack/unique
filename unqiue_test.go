package unique

import "testing"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 11:14
 * @Desc:
 */

func TestUnique_Time64(t *testing.T) {
	u := &Unique{}
	t.Log(u.Time64())
}

func TestUnique_Time32(t *testing.T) {
	u := &Unique{}
	t.Log(u.Time32())
}

func TestGetUniqueCodeByTime(t *testing.T) {
	u := &Unique{}
	t.Log(u.GetUniqueCodeByTime())
}

func TestUnique_GetUniqueCodeBySnowflake(t *testing.T) {
	u := &Unique{}
	t.Log(u.GetUniqueCodeBySnowflake().String())
}
