package main

import (
	"fmt"
	"os"
)

func main(){
	file,err:= os.OpenFile("./aaa.txt",os.O_CREATE|os.O_WRONLY,0777)
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	file.Write([]byte("jjjjj\n"))
	file.Write([]byte("mmmm\n"))
	file.Write([]byte("zzzzzz\n"))

	file.Close()
}
