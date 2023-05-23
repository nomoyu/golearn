package test

import (
	"fmt"
	"github.com/nomoyu/golearn/utils"
)

/*测试函数入口*/
func TestJwt() {
	tokenn, _ := utils.GenerateToken(1, "zs")
	fmt.Println(tokenn)

	_, _ = utils.ParseToken(tokenn)
	fmt.Println(utils.IsTokenValid(tokenn))
	fmt.Println(utils.IsTokenValid(tokenn + "111"))

}
