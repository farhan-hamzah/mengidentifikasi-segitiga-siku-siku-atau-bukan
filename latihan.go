package main
import ("fmt"
        "math"
)
func segitiga(a float64, b float64, c float64){
    var hasil float64
    hasil = (a*a) + (b*b)
    hasil = math.Sqrt(float64(hasil))
    if a <= 0 || b <= 0 || c <= 0{
        fmt.Print("bukan segitiga")
    }else if hasil == c {
        fmt.Print("Segitiga")
    }else{
        fmt.Print("Bukan Segitiga")
    }
}
func main(){
    var sisi1, sisi2, sisi3 float64
    fmt.Scan(&sisi1, &sisi2, &sisi3)
    segitiga(sisi1, sisi2, sisi3)
}