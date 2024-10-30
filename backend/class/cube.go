package class

import (
	"math"
	"math/rand"
)

func hashIndex(x, y, z int) int {
	return x + y*5 + z*25
}

// Cube represents a cube with a given side length
type Cube struct {
	currentState []int
	sideLength   int
	blockCount   int
	currentScore int
}

// NewCube creates a new Cube instance
func NewCube(sideLength int) *Cube {
	temp := int(math.Pow(float64(sideLength), 3))
	return &Cube{sideLength: sideLength, blockCount: temp, currentState: make([]int, temp), currentScore: 0}
}

func (c *Cube) GetCurrentState() []int {
	return c.currentState
}

func (c *Cube) GetBlockCount() int {
	return c.blockCount
}

func (c *Cube) GetSideLength() int {
	return c.sideLength
}

func (c *Cube) GetCurrentScore() int {
	return c.currentScore
}

func (c *Cube) CopyCube() *Cube {
	newCube := Cube{blockCount: c.blockCount, currentState: make([]int, c.blockCount), currentScore: c.currentScore, sideLength: c.sideLength}
	copy(newCube.currentState, c.currentState)
	return &newCube
}

func (c *Cube) SetCurrentState(currentState []int) {
	c.currentState = currentState

	c.currentScore = c.CalculateCurrentScore()
}

func (c *Cube) SetRandomStartState() {
	c.currentState = make([]int, c.blockCount)
	for i := 0; i < c.blockCount; i++ {
		c.currentState[i] = i + 1
	}

	// Shuffle the slice
	rand.Shuffle(c.blockCount, func(i, j int) {
		c.currentState[i], c.currentState[j] = c.currentState[j], c.currentState[i]
	})

	c.currentScore = c.CalculateCurrentScore()
}

func (c *Cube) SwitchState(i, j int) {
	c.currentState[i], c.currentState[j] = c.currentState[j], c.currentState[i]
	c.currentScore = c.CalculateCurrentScore()
}

func (c *Cube) GetRandomSuccessor() *Cube {
	i := rand.Intn(c.blockCount)
	j := rand.Intn(c.blockCount)

	for i == j {
		j = rand.Intn(c.blockCount)
	}

	newCube := Cube{blockCount: c.blockCount, currentState: make([]int, c.blockCount), currentScore: c.currentScore, sideLength: c.sideLength}
	copy(newCube.currentState, c.currentState)

	newCube.SwitchState(i, j)

	return &newCube
}

// Fungsi untuk menghitung skor magic cube
func (c *Cube) CalculateCurrentScore() int {
	n := c.GetSideLength()
	cube := c.GetCurrentState()

	magicNum := 315
	totalScore := 0

	// Perhitungan untuk setiap baris
	for z := 0; z < n; z++ {
		for y := 0; y < n; y++ {
			sum := 0
			for x := 0; x < n; x++ {
				sum += cube[hashIndex(x, y, z)]
			}
			totalScore += int(math.Abs(float64(sum - magicNum)))
		}
	}

	// Perhitungan untuk setiap kolom
	for z := 0; z < n; z++ {
		for x := 0; x < n; x++ {
			sum := 0
			for y := 0; y < n; y++ {
				sum += cube[hashIndex(x, y, z)]
			}
			totalScore += int(math.Abs(float64(sum - magicNum)))
		}
	}

	// Perhitungan untuk setiap tiang
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			sum := 0
			for z := 0; z < n; z++ {
				sum += cube[hashIndex(x, y, z)]
			}
			totalScore += int(math.Abs(float64(sum - magicNum)))
		}
	}

	// Perhitungan untuk seluruh diagonal ruang
	for i := 0; i < n; i++ {
		sum1 := 0
		sum2 := 0
		sum3 := 0
		sum4 := 0
		for j := 0; j < n; j++ {
			sum1 += cube[hashIndex(j, j, j)]
			sum2 += cube[hashIndex(j, j, n-j-1)]
			sum3 += cube[hashIndex(j, n-j-1, j)]
			sum4 += cube[hashIndex(n-j-1, j, j)]
		}
		totalScore += int(math.Abs(float64(sum1 - magicNum)))
		totalScore += int(math.Abs(float64(sum2 - magicNum)))
		totalScore += int(math.Abs(float64(sum3 - magicNum)))
		totalScore += int(math.Abs(float64(sum4 - magicNum)))
	}

	// Perhitungan untuk seluruh diagonal bidang pada kubus
	for z := 0; z < n; z++ {
		sum1 := 0
		sum2 := 0
		sum3 := 0
		sum4 := 0
		for i := 0; i < n; i++ {
			sum1 += cube[hashIndex(i, i, z)]
			sum2 += cube[hashIndex(i, n-i-1, z)]
			sum3 += cube[hashIndex(z, i, i)]
			sum4 += cube[hashIndex(z, i, n-i-1)]
		}
		totalScore += int(math.Abs(float64(sum1 - magicNum)))
		totalScore += int(math.Abs(float64(sum2 - magicNum)))
		totalScore += int(math.Abs(float64(sum3 - magicNum)))
		totalScore += int(math.Abs(float64(sum4 - magicNum)))
	}

	return -totalScore
}

func (c *Cube) GetBestSuccessor() *Cube {
	tempCube := c.CopyCube()
	bestCube := c.CopyCube()

	bestScore := c.currentScore
	for i := 0; i < c.blockCount-1; i++ {
		for j := i + 1; j < c.blockCount; j++ {
			newCube := tempCube.CopyCube()
			newCube.SwitchState(i, j)

			if newCube.GetCurrentScore() == 0 {
				return newCube
			}

			if newCube.GetCurrentScore() < bestScore {
				bestCube = newCube
				bestScore = newCube.GetCurrentScore()
			}
		}
	}

	return bestCube
}

// func main() {
// 	cube := NewCube(5)
// 	cube.SetRandomStartState()
// 	fmt.Println("Initial score: ", cube.GetCurrentScore())

// 	cu := []int{
// 		// z=1
// 		25, 16, 80, 104, 90,
// 		115, 98, 4, 1, 97,
// 		42, 111, 85, 2, 75,
// 		66, 72, 27, 102, 48,
// 		67, 18, 119, 106, 5,

// 		// z=2
// 		91, 77, 71, 6, 70,
// 		52, 64, 117, 69, 13,
// 		30, 118, 21, 123, 23,
// 		26, 39, 92, 44, 114,
// 		116, 17, 14, 73, 95,

// 		// z=3
// 		47, 61, 45, 76, 86,
// 		107, 43, 38, 33, 94,
// 		89, 68, 63, 58, 37,
// 		32, 93, 88, 83, 19,
// 		40, 50, 81, 65, 79,

// 		// z=4
// 		31, 53, 112, 109, 10,
// 		12, 82, 34, 87, 100,
// 		103, 3, 105, 8, 96,
// 		113, 57, 9, 62, 74,
// 		56, 120, 55, 49, 35,

// 		// z=5
// 		121, 108, 7, 20, 59,
// 		29, 28, 122, 125, 11,
// 		51, 15, 41, 124, 84,
// 		78, 54, 99, 24, 60,
// 		36, 110, 46, 22, 101,
// 	}

// 	cube.SetCurrentState(cu)

// 	fmt.Println("Initial score: ", cube.GetCurrentScore())
// }
