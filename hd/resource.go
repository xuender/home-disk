package hd
import(
	"time"
)
type Resource struct {
  Name string // 文件名
	CreateAt time.Time // 创建时间
	Mime string // 类型
}
