package main

import (
	"fmt"
	"runtime"
)

type Inner struct {
	Func   string `json:"func"`
	Period string `json:"period"`
}

type Outer struct {
	WorkspaceId string `json:"workspaceId"`
	Inner
}

// func main() {
// 	router := gin.Default()
// 	router.GET("/ping", func(c *gin.Context) {
// 		o := Outer{
// 			WorkspaceId: "123",
// 			Inner: Inner{
// 				Func:   "sum",
// 				Period: "last_month",
// 			},
// 		}
// 		c.JSON(200, o)
// 	})
//
// 	router.POST("/test", func(c *gin.Context) {
// 		var o Outer
// 		if err := c.ShouldBindBodyWithJSON(&o); err != nil {
// 			fmt.Println("error", err)
// 			c.AbortWithError(400, err)
// 		}
// 		val, _ := json.Marshal(o)
// 		fmt.Println("outer: \n", string(val))
// 		c.IndentedJSON(200, "success")
// 	})
// 	router.Run() // listens on 0.0.0.0:8080 by default
// }

func main() {
	fmt.Println("Logical CPUs: ", runtime.GOMAXPROCS(0))
}
