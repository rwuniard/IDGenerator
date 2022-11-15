package main

import (
	"idgenerator/controllers"
	"idgenerator/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
}

func main() {
	web := gin.Default()
	web.GET("/generateId", controllers.GenId)
	web.Run()
	// var id int64
	// start := time.Now()
	// fmt.Println("Start Time:", start)
	// for i := 0; i < 10000; i++ {
	// 	id = GenerateID(1, 1)
	// }
	// end := time.Now()
	// fmt.Println("End Time:", end)
	// fmt.Println("Duration:", end.Sub(start))
	// fmt.Println("id:", id)

}
