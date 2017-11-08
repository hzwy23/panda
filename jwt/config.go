package jwt

// jwt配置类
// TODO
// 继续完善功能需求，增加更多可选的配置项
type jwtConfig struct {
	key      []byte
	duration int64
	ipValid  bool
	owner    string
}

// 创建配置实例对象
func NewJwtConfig(key []byte) *jwtConfig {
	return &jwtConfig{
		key:      key,
		duration: 3600,
		ipValid:  false,
		owner:    "hzwy23",
	}
}

// 默认配置
var defaultConfig = &jwtConfig{
	key:      []byte("hzwy23@163.com-jwt"),
	duration: 3600,
	ipValid:  false,
	owner:    "hzwy23",
}
