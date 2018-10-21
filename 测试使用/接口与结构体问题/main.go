package main

/*接口与结构体的问题*/
import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Cofox struct {
	name string
}

func (c *Cofox) String() string {
	//return "Joel " + c.name
	return fmt.Sprintf("full name is Joel %v", c.name) //
}

func main() {
	var S Stringer
	c := Cofox{"Smith"} //可以使用c.string()
	S = &c              //可以使用S.string(),将c的地址给S

	fmt.Println(S.String()) //joel smitch
	fmt.Println(c.String()) //joel smitch
	fmt.Println(c.name)     //smitch
}
