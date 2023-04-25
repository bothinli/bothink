package main

import (
	"fmt"
	"github.com/bothinli/bothink-core/models"
)

/**
 * @Author : bothinli
 * @Description:
 */

func main() {
	var item models.Item
	item = models.Job{Name: "testJob"}
	fmt.Println(item.GetName())
}
