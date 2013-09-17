package main

import (
    "log"
    "errors"
)

var wheels []*Wheel


var alphabet = map[string]int{
    "2" : 0,
    "T" : 1,
    "3" : 2,
    "O" : 3, 
    "4" : 4,
    "H" : 5,
    "N" : 6,
    "M" : 7,
    "5" : 8,
    "L" : 9,
    "R" : 10,
    "G" : 11,
    "I" : 12,
    "P" : 13,
    "C" : 14,
    "V" : 15,
    "E" : 16,
    "Z" : 17,
    "D" : 18,
    "B" : 19,
    "S" : 20,
    "Y" : 21,
    "F" : 22,
    "X" : 23,
    "A" : 24,
    "W" : 25,
    "J" : 26,
    "6" : 27,
    "U" : 28,
    "Q" : 29,
    "K" : 30,
    "7" : 31,
}



type Wheel struct{
    Items []int
    CurrentIndex int  //Starts at zero
    MaxSize int
}


func NewWheel(max_size int) (w *Wheel){
    w = new(Wheel)
    w.MaxSize = max_size

    w.Items = []int{1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0}

    w.CurrentIndex = 0
    return w

}

func (w *Wheel) CurrentBit() (bit int){
    bit = w.Items[w.CurrentIndex]
    w.CurrentIndex = (w.CurrentIndex + 1) % w.MaxSize
    return 
}




//EncryptCharacter takes a single character and encrypts it with all ten wheels in Wheels
func EncryptCharacter(char string) error {

    c, ok := alphabet[char]
    if !ok {
        return errors.New("error: character not in alphabet")
    }
    return nil

    var i uint8
    for i = 0; i < 5; i++ {
        c = (c ^ wheels[i].CurrentBit()) << (4-i) //
    }
    
    return nil
}



func main(){
    //TODO initialize Wheels

    wheels = make([]*Wheel, 10)

    for i := 0; i <10; i++ {
        wheels[i] = NewWheel(47)
    }


    for j := 0; j < 10; j++ {
        log.Print(wheels[0].CurrentBit())
    }

    log.Print(1 ^ 1)
}
