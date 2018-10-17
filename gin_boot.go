package ginboot

import (
	"github.com/letcgo/gin"
	"gopkg.in/yaml.v2"
	"os"
	"fmt"
)

type App struct {
	Name string
	Ver string
	Host string
	Port int
	//ReadTimeoutMS  int
	//writeTimeoutMS int
}
type Config struct {
	App App
}

func init(){
	f,err1 := os.Open("config/prod.boot.yml")
	conf := Config{}
	err2 := yaml.NewDecoder(f).Decode(&conf)
	fmt.Println(conf, err1, err2)
	//init config
	//yaml.Unmarshal()
}

func New()*gin.Engine{
	return gin.New()
}
