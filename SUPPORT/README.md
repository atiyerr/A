下载gin框架：go get -u github.com/gin-gonic/gin
提取配置文件用到的viper：go get github.com/spf13/viper
下载gorm框架：go get -u gorm.io/gorm
下载mysql驱动：go get -u github.com/go-sql-driver/mysql
下载redis驱动：go get -u github.com/go-redis/redis/v8
下载jwt：go get github.com/golang-jwt/jwt/v5
下载bcrypt：go get golang.org/x/crypto/bcrypt
管理依赖包：go mod tidy
config文件夹：设置配置信息
router文件夹：设置路由信息
controllers文件夹：控制代码
models文件夹：设置数据库映射模型
global文件夹：设置全局变量
utils文件夹：设置工具类