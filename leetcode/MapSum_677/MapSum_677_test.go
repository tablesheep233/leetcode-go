package MapSum_677

import (
	"fmt"
	"testing"
)

func Test_MapSum_677(t *testing.T) {
	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	ms.Insert("apple", 2)
	fmt.Println(ms.Sum("ap"))

}
