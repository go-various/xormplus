package utils

//A 65 Z 90
//a 97 z 122
func TransCamelToUnderline(n string)string  {
	if n == ""{
		return n
	}
	var out []uint8
	for i := range n {
		if n[i] >= 'A' && n[i]<='Z'{
			if i > 0{
				out = append(out, '_')
			}
			out = append(out, n[i]+32)
		}else {
			out = append(out, n[i])
		}
	}
	return string(out)
}
