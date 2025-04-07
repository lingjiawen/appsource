/*
* @desc:雪花ID生成
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/6/2 14:52
 */

package snowIDGen

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sony/sonyflake"
	"mangosmithy/internal/app/common/service"
	"math/big"
)

var machineID uint16 = 1

const base36Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeBase36(id uint64) string {
	var result []byte
	num := new(big.Int).SetUint64(id)
	base := big.NewInt(36)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result = append([]byte{base36Chars[mod.Int64()]}, result...)
	}

	return string(result)
}

func init() {
	service.RegisterSnowID(New())
}

func New() service.ISnowID {
	return &sSnowID{
		sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: gtime.NewFromStr("2010-05-01").Time,
			MachineID: GetMachineId,
		}),
	}
}

type sSnowID struct {
	*sonyflake.Sonyflake
}

func (s *sSnowID) GenID() (uint64, error) {
	return s.NextID()
}

func (s *sSnowID) GenBase36ID() (string, error) {
	id, err := s.NextID()
	if err != nil {
		return "", err
	}
	return EncodeBase36(id), nil
}

func GetMachineId() (uint16, error) {
	return machineID, nil
}
