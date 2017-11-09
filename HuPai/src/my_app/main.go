package main

import (
        "fmt"
        "bufio"
        "os"
        "strings"
        "strconv"
        "time"
        "github.com/go-vgo/robotgo"
)


func main() {
    var Path string = "C:\\Go\\myfile\\"
    /*
     * Get pos
     */
    f, _ := os.OpenFile(Path + "first_pos", os.O_RDWR, 0)
    buff := bufio.NewReader(f)
    line1, _ := buff.ReadString('\n')
    line1 = strings.TrimSpace(line1)
    line2, _ := buff.ReadString('\n')
    line2 = strings.TrimSpace(line2)
    x1, _ := strconv.Atoi(line1)
    y1, _ := strconv.Atoi(line2)
    f.Close()
    fmt.Printf("%d, %d\n", x1, y1)
    
    /*
     * Get pos
     */
    f, _ = os.OpenFile(Path + "seconde_pos", os.O_RDWR, 0)
    buff  = bufio.NewReader(f)
    line1, _ = buff.ReadString('\n')
    line1 = strings.TrimSpace(line1)
    line2, _ = buff.ReadString('\n')
    line2 = strings.TrimSpace(line2)
    x2, _ := strconv.Atoi(line1)
    y2, _ := strconv.Atoi(line2)
    f.Close()
    fmt.Printf("%d, %d\n", x2, y2)

    /*
     * Get pos
     */
    f, _ = os.OpenFile(Path + "third_pos", os.O_RDWR, 0)
    buff  = bufio.NewReader(f)
    line1, _ = buff.ReadString('\n')
    line1 = strings.TrimSpace(line1)
    line2, _ = buff.ReadString('\n')
    line2 = strings.TrimSpace(line2)
    x3, _ := strconv.Atoi(line1)
    y3, _ := strconv.Atoi(line2)
    f.Close()
    fmt.Printf("%d, %d\n", x3, y3)

    /*
     * Get pos
     */
    f, _ = os.OpenFile(Path + "fourth_pos", os.O_RDWR, 0)
    buff  = bufio.NewReader(f)
    line1, _ = buff.ReadString('\n')
    line1 = strings.TrimSpace(line1)
    line2, _ = buff.ReadString('\n')
    line2 = strings.TrimSpace(line2)
    x4, _ := strconv.Atoi(line1)
    y4, _ := strconv.Atoi(line2)
    f.Close()
    fmt.Printf("%d, %d\n", x4, y4)

    
    local, _ := time.LoadLocation("Local")
    /*
     * Get time
     */
    f, _ = os.OpenFile(Path + "time", os.O_RDWR, 0)
    buff  = bufio.NewReader(f)
    line1, _ = buff.ReadString('\n')
    line1 = strings.TrimSpace(line1)
    time1 := line1 + os.Args[2]
    time2 := line1 + os.Args[3]
    t1, _ := time.ParseInLocation("2006-01-02 15:04:05 ", time1, local)
    t2, _ := time.ParseInLocation("2006-01-02 15:04:05 ", time2, local)
    
    /*
     * Type the price bonus
     */
    //FIXME need, compare the time
    fmt.Printf("Waiting the time to add bonus!!!!!!!!!!!!!!!!: \n")
    fmt.Println(t1)
    fmt.Printf("Waiting !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")

    for {
        if time.Now().Unix() >= t1.Unix() {
            fmt.Printf("Price bonus %s\n", os.Args[1])
            robotgo.MoveMouse(x1, y1)
            robotgo.MouseClick()
            robotgo.MouseClick()
            robotgo.TypeString("backspace")
            robotgo.TypeString(os.Args[1])
            
            fmt.Printf("click pos2 %d %d\n", x2, y2)
            robotgo.MoveClick(x2, y2)
            time.Sleep(200 * time.Millisecond)
            robotgo.MoveClick(x2, y2)
            time.Sleep(200 * time.Millisecond)
            robotgo.MoveClick(x2, y2)
            time.Sleep(200 * time.Millisecond)
            fmt.Printf("click pos3 %d %d\n", x3, y3)
            robotgo.MoveClick(x3, y3)
            break
        }
        time.Sleep(500 * time.Millisecond)
        fmt.Println(time.Now())
    }

    /*
     * Now, we at t2, send out the price
     */
    fmt.Printf("Now we will wait to send out price!!!!!!!!!!!!!!!!!!!!!!!\n")
    for {
        if time.Now().Unix() >= t2.Unix() {
            robotgo.MoveClick(x4, y4)
            break
        }
        time.Sleep(100 * time.Millisecond)
        fmt.Println(time.Now())
    }
    fmt.Printf("Successfully send out the price!!!\n")
    fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
    fmt.Printf("!!!!!!!!!!!!!!!!!!!!!GOD BLESS ME!!!!!!!!!!!!!!!!!!!!!!!!\n")
    fmt.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")

}
