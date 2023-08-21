package main
import(
	"fmt"
	"net/http"
	"os"
	"io"
)
type logWriter struct{}

func main(){
	resp,err:= http.Get("http://google.com")
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// bs:=[]byte{}
	// bs:=make([]byte,99999)
	// resp.Body.Read(bs)
	// fmt.Println("bs:--------------\n",string(bs))
	// io.Copy(os.Stdout, resp.Body)

	
	lw:=logWriter{}
	io.Copy(lw, resp.Body)
}

func(l logWriter) Write(bs []byte)(int,error){
	fmt.Println(string(bs))
	fmt.Println("wrote this many bytes:",len(bs))
	return len(bs),nil
}
