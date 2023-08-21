package main
import(
	"fmt"
	"os"
	"io"
)
func main(){
	fmt.Println(os.Args)
	filename:=os.Args[1]
	

	//open file ->file pointer
	f,err:=os.Open(filename)
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	//read file content
	io.Copy(os.Stdout, f)


}