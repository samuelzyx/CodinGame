package main

import ("fmt"
        "os"
        "bufio"
        "strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var w int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&w)
    
    var h int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&h)

    var duck1X1, duck1Y1, duck1X2, duck1Y2 int
    var duck2X1, duck2Y1, duck2X2, duck2Y2 int

    //-----------------TURN 1--------------------------
    fmt.Fprintln(os.Stderr, "Field view turn 1")
    for y := 0; y < h; y++ {
        scanner.Scan()
        row := scanner.Text()
        _ = row // to avoid unused error
        chars := []rune(row)
        for x := 0; x < len(chars); x++ {
            char := string(chars[x])
            //fmt.Fprintln(os.Stderr, char)
            //if char is equals to id=2 save value in duck2X1, duck2Y1
            if(char=="2"){
                // character to ASCII
                duck2X1 = x
                duck2Y1 = y
            }
            if(char=="1"){
                // character to ASCII
                duck1X1 = x
                duck1Y1 = y
            }
        }
        fmt.Fprintln(os.Stderr, row)    
    }
    //-----------------TURN 2--------------------------
    fmt.Fprintln(os.Stderr, "Field view turn 2")
    for y := 0; y < h; y++ {
        scanner.Scan()
        row := scanner.Text()
        _ = row // to avoid unused error
        chars := []rune(row)
        for x := 0; x < len(chars); x++ {
            char := string(chars[x])
            if(char=="2"){
                duck2X2 = x
                duck2Y2 = y
            }
            if(char=="1"){
                // character to ASCII
                duck1X2 = x
                duck1Y2 = y
            }
        }
        fmt.Fprintln(os.Stderr, row)
    }
    log := "Found PATO 1 | TURN 1 | ID=1 in X:"+strconv.FormatInt(int64(duck1X1), 10)+" Y="+strconv.FormatInt(int64(duck1Y1), 10)
    fmt.Fprintln(os.Stderr, log)
    //log = "Found PATO 2 | TURN 1 | ID=2 in X:"+strconv.FormatInt(int64(duck2X1), 10)+" Y="+strconv.FormatInt(int64(duck2Y1), 10)
    //fmt.Fprintln(os.Stderr, log)

    log = "Found PATO 1 | TURN 2 | ID=1 in X:"+strconv.FormatInt(int64(duck1X2), 10)+" Y="+strconv.FormatInt(int64(duck1Y2), 10)
    fmt.Fprintln(os.Stderr, log)
    //log = "Found PATO 2 | TURN 2 | ID=2 in X:"+strconv.FormatInt(int64(duck2X2), 10)+" Y="+strconv.FormatInt(int64(duck2Y2), 10)
    //fmt.Fprintln(os.Stderr, log)

    fmt.Fprintln(os.Stderr, "PREDICT PATO 1 | TURN 3...")
    duck1X3 := duck1X2 - duck1X1 //2-3=-1
    duck1X3 = duck1X2 + duck1X3 //2-1=1
    duck1Y3 := duck1Y2 - duck1Y1
    duck1Y3 = duck1Y2+duck1Y3
    log = "PATO 1 | TURN 3 | ID=1 in X:"+strconv.FormatInt(int64(duck1X3), 10)+" Y="+strconv.FormatInt(int64(duck1Y3), 10)
    fmt.Fprintln(os.Stderr, log)

    fmt.Fprintln(os.Stderr, "PREDICT PATO 1 | TURN 4...")
    duck1X4 := duck1X3 - duck1X2
    duck1X4 = duck1X3 + duck1X4
    duck1Y4 := duck1Y3 - duck1Y2
    duck1Y4 = duck1Y3+duck1Y4
    
    log = "PATO 1 | TURN 4 | ID=1 in X:"+strconv.FormatInt(int64(duck1X4), 10)+" Y="+strconv.FormatInt(int64(duck1Y4), 10)
    fmt.Fprintln(os.Stderr, log)

    //fmt.Fprintln(os.Stderr, "PREDICT PATO 2 | TURN 3...")
    duck2X3 := duck2X2 - duck2X1
    duck2X3 = duck2X2+duck2X3
    duck2Y3 := duck2Y2 - duck2Y1
    duck2Y3 = duck2Y2+duck2Y3
    /*log = "PATO 2 | TURN 3 | ID=2 in X:"+strconv.FormatInt(int64(duck2X3), 10)+" Y="+strconv.FormatInt(int64(duck2Y3), 10)
    fmt.Fprintln(os.Stderr, log)*/

    //If duck 1 
    log = "2 "+strconv.FormatInt(int64(duck2X3), 10)+" "+strconv.FormatInt(int64(duck2Y3), 10)
    fmt.Println(log)


    log = "1 "+strconv.FormatInt(int64(duck1X4), 10)+" "+strconv.FormatInt(int64(duck1Y4), 10)
    fmt.Println(log)
}