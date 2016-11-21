package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"ethos/efmt"
	"log"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}


	data    := MyType { 1, 2, 3, 4 }

	data.Write(fd)
	data.WriteVar(path +"foobar")
	efmt.Println("Hello! Harsha")
	efmt.Println("myOutput:",data.x1,data.y1,data.x2,data.y2)
	data.x1 = data.x1 + 10
	data.y1 = data.y1 + 10
 	data.x2 = data.x2 + 10
	data.y2 = data.y2 + 10

	efmt.Println("myOutput:",data.x1,data.y1,data.x2,data.y2)
}

