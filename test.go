package main

func main() {
  e := echo.New()
  
  e.GET( path: "/hello", hello)
  
  e.GET( path: "/ping", ping)
  
  e.Logger.Fatal(e.Start( address: ":3000"))
}

func hello( c echo.Context) error {
  return c.String(http.StatusOK, s: "Hello, World!!!!!")
  
}

func ping( c echo.Context) error {
  return c.String(http.StatusOK, s: "Pong")
}
