package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
  //"log"
  //"fmt"
)

func main() {
  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  router.GET("/:n", func(c *gin.Context) {
    n, _ := strconv.ParseUint(c.Param("n"), 10, 64)
    //log.Println("Err: ", err)
    res := fact(n)
    //message := fmt.Sprintf("Factorial of %d is %d", n, res)
    //log.Println("M: ",message)
    c.String(http.StatusOK, string(res))
  })

  router.Run(":8091")
}

func fact(n uint64) uint64 {
  if (n == 1) {
    return 1
  }
  return n*fact(n-1)
}
