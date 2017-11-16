var intValues = map[rune]int{
     'I':1, 
     'V':5, 
     'X':10, 
     'L':50, 
     'C':100,
     'D':500,
     'M':1000 }

func romanToInt(s string) int {
    total := 0
    decimal := 0

    for i,_:= range s {
        i = len(s) - 1 - i
        s1 := intValues[rune(s[i])]
        if s1 < decimal {
            total -=s1
        } else {
            total +=s1
        }
        decimal = intValues[rune(s[i])]
    }
    return total
}
