package arrays3

func pow(x, n int) int{
	if n == 0{
		return 1
	}

	if n < 0{
		x = 1/x
		n = -n
	}

	return x * pow(x, n-1)
}