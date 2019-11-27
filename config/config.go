package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Conf interface {
	NewConfig (filePath string) //初始化一个配置文件
	GetString (name string, key string) string //获取字符串
	GetInt (name string, key string) int //获取数字
	GetAll (name string) map[string]interface{} //获取单个配置项的所有内容

}
type ConfData struct {
	filePath string
	list map[string]map[string]interface{}
}

func NewConfig(filePath string) *ConfData  {
	c := new(ConfData)
	c.filePath = filePath
	c.list = make(map[string]map[string]interface{})
	c.read()
	return c
}

func (c *ConfData) GetString (name string, key string) string  {
	if _, ok := c.list[name]; !ok {
		return "";
	}
	return c.list[name][key].(string)
}
func (c *ConfData) GetInt (name string, key string) int  {
	if _, ok := c.list[name]; !ok {
		return 0;
	}
	return c.list[name][key].(int)
}

func (c *ConfData) GetAll(name string) map[string]interface{}  {

	if _, ok := c.list[name]; !ok {
		return make(map[string]interface{});
	}
	return c.list[name]
}

func (c *ConfData) read() {
	if c.filePath == "" {
		panic("not file path!")
	}
	file, err := os.Open(c.filePath)
	if err != nil {
		panic("Err file path")
	}
	defer file.Close()
	var header string
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Read configuration file error: %s", err.Error())
			}
		}
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue //annotation
		}
		if line[0] == '[' {
			//header
			line = strings.TrimSpace(line[1 : len(line)-1])
			c.list[line] = make(map[string]interface{})
			header = line
			continue
		}
		//body
		if _, ok := c.list[header]; !ok {
			log.Println("Missing configuration items %s", line)
			continue
		}
		exStr := strings.Split(line, "=")
		c.list[header][exStr[0]] = exStr[1]
	}
}