package jwt

// jwt配置类
// TODO
// 继续完善功能需求，增加更多可选的配置项
type jwtConfig struct {
	key     []byte
	ipValid bool
}

// 创建配置实例对象
func NewJwtConfig(key []byte) *jwtConfig {
	return &jwtConfig{
		key:     key,
		ipValid: false,
	}
}

// 默认配置
var defaultConfig = &jwtConfig{
	key:     []byte("hzwy23@163.com-jwt"),
	ipValid: false,
}
