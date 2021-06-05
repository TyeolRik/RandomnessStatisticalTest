# Randomness Statistical Test

I haven't finished implementing NIST SP800-22. **There is no "2.8 The Overlapping Template Matching Test"**. Because there is some misprint and errors in official documents. I need time to cross-check. :(

## Introduction

This project implements "Randomness Statistical Tests" in Go-Language (1.16.4).

I have been curious **what Random is**. Even though I read some documents and papers, I can't tell the definition of Random. However, in NIST SP800-22, they define random as coin flip. Looked at in computer perspective, there is just two state, 0 or 1. It is same with coin, which has just two state, head and tail.

So, NIST SP800-22 tests focus on statistics. Because random bits and coin flips are same fundamentally, random bits means that there is 50% chance to appear.

## How to use

#### **`main.go`**
```go
func main() {
    Prepare_CONSTANT_E_asEpsilon()          // Input is natural number E in binary.
    testArray := GetEpsilon()[0:1000000]    // Epsilon means test bits (in terms of NIST SP800-22)
    Examine_NIST_SP800_22(testArray, 0.01)  // Examine 15 tests.
}
```

If you want to test in more detail, you can adjust some options (like Block size) in function, ```Examine_NIST_SP800_22()```.

## Result example

```
$ go run main.go
+----+------------------------------------------------+----------------+------------+---------------+
|  # |                    TEST NAME                   | SUB TEST COUNT | CONCLUSION |    P-VALUE    |
+----+------------------------------------------------+----------------+------------+---------------+
|  1 |          The Frequency (Monobit) Test          |        -       |   Random   | 0.95374862853 |
+----+------------------------------------------------+----------------+------------+---------------+
|  2 |          Frequency Test within a Block         |        -       |   Random   | 0.17667504647 |
+----+------------------------------------------------+----------------+------------+---------------+
|  3 |                  The Runs Test                 |        -       |   Random   | 0.56191688503 |
+----+------------------------------------------------+----------------+------------+---------------+
|  4 |  Tests for the Longest-Run-of-Ones in a Block  |        -       |   Random   | 0.71894532990 |
+----+------------------------------------------------+----------------+------------+---------------+
|  5 |           The Binary Matrix Rank Test          |        -       |   Random   | 0.63739676560 |
+----+------------------------------------------------+----------------+------------+---------------+
|  6 | The Discrete Fourier Transform (Spectral) Test |        -       |   Random   | 0.84718670507 |
+----+------------------------------------------------+----------------+------------+---------------+
|  7 |   The Non-overlapping Template Matching Test   | 245 / 250 PASS |   Random   |       -       |
|  - |                                                |        1       |   Random   | 0.53654959420 |
|  - |                                                |        2       |   Random   | 0.23246291065 |
|  - |                                                |       ...      |            |               |
|  - |                                                |       249      |   Random   | 0.42150925672 |
|  - |                                                |       250      |   Random   | 0.05370018454 |
+----+------------------------------------------------+----------------+------------+---------------+
|  8 |      Maurer's "Universal Statistical" Test     |        -       |   Random   | 0.99963445481 |
+----+------------------------------------------------+----------------+------------+---------------+
|  9 |             Linear Complexity Test             |        -       |   Random   | 0.77480980929 |
+----+------------------------------------------------+----------------+------------+---------------+
| 10 |                   Serial Test                  |   2 / 2 PASS   |   Random   |       -       |
|  - |                   Serial Test                  |        1       |   Random   | 0.84376437495 |
|  - |                   Serial Test                  |        2       |   Random   | 0.56191461785 |
+----+------------------------------------------------+----------------+------------+---------------+
| 11 |            Approximate Entropy Test            |        -       |   Random   | 0.36168763190 |
+----+------------------------------------------------+----------------+------------+---------------+
| 12 |          Cumulative Sums (Cusum) Test          |   2 / 2 PASS   |   Random   |       -       |
|  - |          Cumulative Sums (Cusum) Test          |        1       |   Random   | 0.72426530997 |
|  - |          Cumulative Sums (Cusum) Test          |        2       |   Random   | 0.66988646417 |
+----+------------------------------------------------+----------------+------------+---------------+
| 13 |             Random Excursions Test             |   8 / 8 PASS   |   Random   |       -       |
|  - |                                                |        1       |   Random   | 0.26136281011 |
|  - |                                                |        2       |   Random   | 0.80619660748 |
|  - |                                                |       ...      |            |               |
|  - |                                                |        7       |   Random   | 0.54766817696 |
|  - |                                                |        8       |   Random   | 0.24544453536 |
+----+------------------------------------------------+----------------+------------+---------------+
| 14 |         Random Excursions Variant Test         |  18 / 18 PASS  |   Random   |       -       |
|  - |                                                |        1       |   Random   | 0.95775095848 |
|  - |                                                |        2       |   Random   | 0.90125654188 |
|  - |                                                |       ...      |            |               |
|  - |                                                |       17       |   Random   | 0.38717084099 |
|  - |                                                |       18       |   Random   | 0.34209300630 |
+----+------------------------------------------------+----------------+------------+---------------+
```

## External Libraries

I am not sure whether writing like this is proper or not. If this notation is wrong, please let me know how to fix this.

- Cephes Math Library (igamc, igam) : [https://www.netlib.org/cephes/](https://www.netlib.org/cephes/)

- mjibson/go-dsp : [https://github.com/mjibson/go-dsp](https://github.com/mjibson/go-dsp)

- jedib0t/go-pretty : [https://github.com/jedib0t/go-pretty](https://github.com/jedib0t/go-pretty)
