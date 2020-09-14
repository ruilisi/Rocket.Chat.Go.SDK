// +build !release

package conf

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	// any approach to require this configuration into your program.
	str, _ := os.Getwd()
	str = strings.Replace(str, "/test/module-test", "", 1)
	str = strings.Replace(str, "test/goroutine", "", 1)
	str = strings.Replace(str, "test/rm", "", 1)
	url := str + "/application.yml"
	data, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(data))
}
