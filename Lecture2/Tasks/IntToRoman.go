package Tasks

func IntToRoman(num int) string {

	ans := ""
	k := 0
	str := []string{"I", "V", "X", "L", "C", "D", "M"}
	for num > 0 {
		d := num % 10
		v := ""
		if d < 4 {
			for j := 0; j < d; j++ {
				v += str[k]
			}
		} else if d == 4 {
			v += str[k] + str[k+1]
		} else if d < 9 {
			v += str[k+1]
			for j := 0; j < d-5; j++ {
				v += str[k]
			}
		} else if d == 9 {
			v += str[k] + str[k+2]
		}
		ans = v + ans
		k += 2
		num /= 10
	}

	return ans
}
