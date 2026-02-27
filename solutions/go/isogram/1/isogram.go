package isogram
import "strings"

func IsIsogram(word string) bool {
    word = strings.ToLower(word)
	for i:=0;i<len(word);i++{
        for j:=i+1 ; j<len(word);j++{
			if word[j] == ' ' || word[j] == '-' {
				continue
			}
            if word[i]==word[j]{
                if word[i] != '-'{
                    return false
                } else{
                    continue
                }
            } else{
                continue
            }
        }
    }
    return true
}
