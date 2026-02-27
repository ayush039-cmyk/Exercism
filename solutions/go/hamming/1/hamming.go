package hamming
import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b){
        return 0,errors.New("error found")
    }
	total := 0
    for i:=0 ; i<=len(a)-1 ;i++{
        if a[i] != b[i]{
            total += 1
        }
    }
    return total,nil
}
