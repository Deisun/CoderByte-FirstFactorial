# CoderByte-FirstFactorial
## Challenge
Using the Go language, have the function FirstFactorial(num) take the num parameter being passed and return the factorial of it (e.g. if num = 4, return (4 * 3 * 2 * 1)). For the test cases, the range will be between 1 and 18 and the input will always be an integer. 

## Examples
Input:4
Output:24


Input:8
Output:40320

## Testing
```
	tests := []testpair{
		{0,1},
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{15,1307674368000},
	}


=== RUN   TestFirstFactorial
--- PASS: TestFirstFactorial (0.00s)
PASS
```
