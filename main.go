package main
func main(){
	app := App{}
	app.Start(DbUser,DbPassword,DBName)
	app.Run(":8080")
}