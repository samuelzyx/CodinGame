package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // W: width of the building.
    // H: height of the building.
    var W, H int //Min
    fmt.Scan(&W, &H)
    
    // N: maximum number of turns before game over.
    var N int
    fmt.Scan(&N)
    
    var X0, Y0 int 
    fmt.Scan(&X0, &Y0)

    fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Fprintln(os.Stderr, "Edificio Ancho: ",W)
    fmt.Fprintln(os.Stderr, "Edificio Alto: ",H)
    batman := "(" + strconv.FormatInt(int64(X0), 10) + "," + strconv.FormatInt(int64(Y0), 10) + ")"
    fmt.Fprintln(os.Stderr, "Posicion Batman: ",batman)
    
    var bombDir string
    var guessX = -1
    var guessY = -1 
    var minX, maxX, minY, maxY int
    minX=0
    maxX=W
    minY=0
    maxY=H

    // bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
    fmt.Fprintln(os.Stderr, "")
    fmt.Fprintln(os.Stderr, "Siguiente escaneo de la bomba")
    fmt.Scan(&bombDir) //End loop
    if(strings.Contains(bombDir, "D")){
        fmt.Fprintln(os.Stderr, "---------------bomba mas abajo")
    }
    if(strings.Contains(bombDir, "U")){ 
        fmt.Fprintln(os.Stderr, "---------------bomba mas arriba")
    }

    for {//Debug

        valX := "Valores X: Min="+strconv.FormatInt(int64(minX), 10)+" Max="+strconv.FormatInt(int64(maxX), 10)
        fmt.Fprintln(os.Stderr, valX)
        valY := "Valores Y: Min="+strconv.FormatInt(int64(minY), 10)+" Max="+strconv.FormatInt(int64(maxY), 10)
        fmt.Fprintln(os.Stderr, valY)

        batman := "(" + strconv.FormatInt(int64(X0), 10) + "," + strconv.FormatInt(int64(Y0), 10) + ")"
        fmt.Fprintln(os.Stderr, "Posicion Batman: ",batman)
        
        if(strings.Contains(bombDir, "R")||strings.Contains(bombDir, "L")){
            guessX = (maxX+minX)/2 
            operacion := "(" + strconv.FormatInt(int64(maxX), 10) + "+" + strconv.FormatInt(int64(minX), 10) + ")/2="+strconv.FormatInt(int64(guessX), 10)
            fmt.Fprintln(os.Stderr, "Operacion en X: ",operacion)
            fmt.Fprintln(os.Stderr, "Adivinando en X: ",guessX)
            X0=guessX
        }

        if(strings.Contains(bombDir, "D")||strings.Contains(bombDir, "U")){
            guessY = (maxY+minY)/2
            operacion := "(" + strconv.FormatInt(int64(maxY), 10) + "+" + strconv.FormatInt(int64(minY), 10) + ")/2="+strconv.FormatInt(int64(guessY), 10)
            fmt.Fprintln(os.Stderr, "Operacion en Y: ",operacion)
            fmt.Fprintln(os.Stderr, "Adivinando en Y: ",guessY)
            Y0=guessY
        }

        // the location of the next window Batman should jump to.
        fmt.Println(X0,Y0)

        // bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        fmt.Fprintln(os.Stderr, "")
        fmt.Fprintln(os.Stderr, "Siguiente escaneo de la bomba")
        fmt.Scan(&bombDir) //End loop
        if(strings.Contains(bombDir, "U")){ 
            fmt.Fprintln(os.Stderr, "---------------bomba mas arriba")
            maxY = guessY-1
        }
        if(strings.Contains(bombDir, "D")){
            fmt.Fprintln(os.Stderr, "---------------bomba mas abajo")
            minY = guessY+1
        }
        
        if(strings.Contains(bombDir, "L")){
            fmt.Fprintln(os.Stderr, "---------------bomba mas izq")
            maxX = guessX-1
        }
        if(strings.Contains(bombDir, "R")){ 
            fmt.Fprintln(os.Stderr, "---------------bomba mas der")
            minX = guessX+1
        }
    }
}