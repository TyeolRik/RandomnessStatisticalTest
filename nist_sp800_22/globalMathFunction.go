package nist_sp800_22

import (
	"math"
)

func DecisionRule(_P_value float64, level float64) bool {
	if level < 0 || level > 1 {
		panic("input level is wrong. this value should be between 0 < level < 1")
	}
	if _P_value < level {
		return false // Non-random
	} else {
		return true // random
	}
}

func DecisionRule2(_P_value1 float64, _P_value2 float64, level float64) bool {
	if level < 0 || level > 1 {
		panic("input level is wrong. this value should be between 0 < level < 1")
	}
	if _P_value1 < level || _P_value2 < level {
		return false // Non-random
	} else {
		return true // random
	}
}

/**
* According to, NIST SP800-22 Page 99, Gamma Function and Imcomplete Gamma Function are described
* Fully Implemented from Cephes C
* Mirror : https://github.com/jeremybarnes/cephes
 */
func igamc(a float64, x float64) float64 {
	if x <= 0 || a <= 0 {
		return 1.0
	}
	if x < 1.0 || x < a {
		return 1.0 - igam(a, x)
	}

	tempLgam, _ := math.Lgamma(a)
	var ax float64 = a*math.Log(x) - x - tempLgam
	if ax < -MAXLOG {
		panic("igamc UNDERFLOW")
	}
	ax = math.Exp(ax)

	/* continued fraction */
	y := 1.0 - a
	z := x + y + 1.0
	c := 0.0
	pkm2 := 1.0
	qkm2 := x
	pkm1 := x + 1.0
	qkm1 := z * x
	ans := pkm1 / qkm1

	// Do-while phase
	var yc, pk, qk, r, t float64
	for {
		c += 1.0
		y += 1.0
		z += 2.0
		yc = y * c
		pk = pkm1*z - pkm2*yc
		qk = qkm1*z - qkm2*yc
		if qk != 0 {
			r = pk / qk
			t = math.Abs((ans - r) / r)
			ans = r
		} else {
			t = 1.0
		}
		pkm2 = pkm1
		pkm1 = pk
		qkm2 = qkm1
		qkm1 = qk
		if math.Abs(pk) > big {
			pkm2 *= biginv
			pkm1 *= biginv
			qkm2 *= biginv
			qkm1 *= biginv
		}
		if t <= MACHEP {
			break
		}
	}

	return ans * ax
}

func igam(a float64, x float64) float64 {
	var ans, ax, c, r float64
	if x <= 0 || a <= 0 {
		return 0.0
	}
	if (x > 1.0) && (x > a) {
		return 1.0 - igamc(a, x)
	}

	/* Compute  x**a * exp(-x) / gamma(a)  */
	tempLgam, _ := math.Lgamma(a)
	ax = a*math.Log(x) - x - tempLgam
	if ax < -MAXLOG {
		panic("igam UNDERFLOW")
	}
	ax = math.Exp(ax)

	/* power series */
	r = a
	c = 1.0
	ans = 1.0

	for {
		r += 1.0
		c *= x / r
		ans += c
		if c/ans <= MACHEP {
			break
		}
	}
	return ans * ax / a
}

func Max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

// According to NIST SP800-22 Revision 1a. Page. 123
// F.1 Rank Computation of Binary Matrices
func RankComputationOfBinaryMatrices(matrix [][]uint8) uint64 {

	// Forward Application of Elementary Row Operations
	// Declare Variables
	var row, col int
	var m int = len(matrix)

	// Step 1. Set i = 1
	i := 0

	// Step 2. If element a(i,i) = 0 (i.e., the element on the diagonal ≠ 1),
	// then swap all elements in the ith row with all elements in the next row that contains a one in the i-th column.
	// (i.e., this row is the kth row, where i < k <= m)
	// If no row contains a “1” in this position, go to step 4.
Forward_STEP2:
	if matrix[i][i] == 0 {
		var tempIndex int
		var isContained bool = false
		for tempIndex = i; tempIndex < m; tempIndex++ {
			if matrix[tempIndex][i] == 1 {
				matrix[i], matrix[tempIndex] = matrix[tempIndex], matrix[i]
				isContained = true
				break
			}
		}
		if !isContained {
			goto Forward_STEP4
		}
	}

	// Step 3. If element a(i,i) = 1, then if any subsequent row contains a “1” in the i-th column,
	// replace each element in that row with the exclusive-OR of that element and the corresponding element in the i-th row.

	// Step 3-a.
	row = i + 1
	// Step 3-b.
Forward_STEP_3B:
	col = i
	// Step 3-c.
	if matrix[row][col] == 0 {
		goto Forward_STEP_3G
	}
	// Step 3-d.
Forward_STEP_3D:
	matrix[row][col] = matrix[row][col] ^ matrix[i][col]
	// Step 3-e.
	if col == (m - 1) {
		goto Forward_STEP_3G
	}
	// Step 3-f.
	col = col + 1
	goto Forward_STEP_3D
	// Step 3-g.
Forward_STEP_3G:
	if row == (m - 1) {
		goto Forward_STEP4
	}
	// Step 3-h.
	row = row + 1
	goto Forward_STEP_3B

	// Step 4.
Forward_STEP4:
	if i < m-2 {
		i = i + 1
		goto Forward_STEP2
	}

	// Step 5. Forward row operations completed.

	// The Subsequent Backward Row Operations
	// Step 1. Set i = m	// But, matrix index [0, m-1].
	i = m - 1

	// Step 2. If element a(i, i) = 0,
Backward_STEP_2:
	if matrix[i][i] == 0 {
		// swap all elements in the i-th row with all elements in the next row that contains a one in the i-th column
		var tempIndex int
		var isContained bool = false
		for tempIndex = i; tempIndex >= 0; tempIndex-- {
			if matrix[tempIndex][i] == 1 {
				matrix[i], matrix[tempIndex] = matrix[tempIndex], matrix[i]
				isContained = true
				break
			}
		}
		if !isContained {
			goto Backward_STEP_4
		}
	}

	// If element a(i, i) = 1,
	// Step 3-a
	row = i - 1

	// Step 3-b
Backward_STEP_3B:
	col = i

	// Step 3-c
	if matrix[row][col] == 0 {
		goto Backward_STEP_3G
	}

	// Step 3-d
Backward_STEP_3D:
	matrix[row][col] = matrix[row][col] ^ matrix[i][col]

	// Step 3-e
	if col == 1 {
		goto Backward_STEP_3G
	}

	// Step 3-f
	col = col - 1
	goto Backward_STEP_3D

	// Step 3-g
Backward_STEP_3G:
	if row == 1 {
		goto Backward_STEP_4
	}

	// Step 3-h.
	row = row - 1
	goto Backward_STEP_3B

	// Step 4.
Backward_STEP_4:
	if i > 2 {
		i = i - 1
		goto Backward_STEP_2
	}

	// Step 5. Backward row operation complete.

	// The rank of the matrix = the number of non-zero rows.
	var rank uint64 = 0
	for _, row := range matrix {
		for _, eachValue := range row {
			if eachValue == 1 {
				rank++
				break
			}
		}
	}

	return rank
}

// https://stats.stackexchange.com/a/187909
// Wolfram Alpha : Online Computation
// https://www.wolframalpha.com/input/?i=1+%2F+sqrt%282+*+pi%29+integral_%28-inf%29%5Ez+e%5E%28-u%5E2+%2F+2%29+du
func CumulativeDistribution(z float64) float64 {
	return 0.5 * (math.Erf(z/math.Sqrt2) + 1.0)
}
