package nist_sp800_22

import (
	"math"
)

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
