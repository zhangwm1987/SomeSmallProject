package main
import (
        "os"
        "runtime"
        "strconv"
        "fmt"
        "github.com/go-vgo/robotgo"
)

var ostype = runtime.GOOS

func main() {

    x, y := robotgo.GetMousePos()
    
    var f *os.File    
    var FileName string    
    
    arg_num := len(os.Args)
    if arg_num == 1 {
        fmt.Printf("Please enter the argument: 1, 2, 3, 4\n")
        fmt.Printf("and run again!!!!!!!!!!!!!!!!!\n")
        return
    }
    
    if os.Args[1] == "1" {
        FileName = "first_pos"
    } else if os.Args[1] == "2" {
        FileName = "seconde_pos"
    } else if os.Args[1] == "3" {
        FileName = "third_pos"
    } else {
        FileName = "fourth_pos"
    }
    
    var Path string = "C:\\Go\\myfile\\"

    fmt.Printf("The position is (%d, %d)\n", x, y)
    if ostype == "windows" {
        f, _ = os.OpenFile(Path + FileName, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0)
        f.WriteString(strconv.Itoa(x) + "\r\n")
        f.WriteString(strconv.Itoa(y) + "\r\n")
    } else {
        f, _ = os.OpenFile("/root/" + FileName, os.O_RDWR | os.O_CREATE | os.O_TRUNC, 0)
        f.WriteString(strconv.Itoa(x) + "\n")
        f.WriteString(strconv.Itoa(y) + "\n")
    
    }
    f.Close()
}
