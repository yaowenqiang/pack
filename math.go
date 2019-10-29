package pack
import (
	"math"
)

//go:generate ls -l

//go generate


type Integrator interface {
	Integrate(min, max float64, coefs ...float64) float64
}


type PolyIntegrator struct {}

type XYFunc func(x float64) float64

func (pi *PolyIntegrator) Integrate(min, max float64, coefs ...float64) float64 {
	newCoefs := coefs[:]

	for i := range newCoefs {
		exp := float64(len(newCoefs) -i)

		newCoefs[i] = newCoefs[i] / exp

	}
	newCoefs = append(newCoefs, 0.)

	return calcPoly(newCoefs...)(max) - calcPoly(newCoefs...)(min)
}

func calcPoly(coefs ...float64) XYFunc {
	return func(x float64) float64 {
		result := 0.
		for i , coef := range coefs {
			result += coef * math.Pow(x, float64(len(coefs) -i -1))
		}
		return result
	}
}

type RiemannIntegrator struct {}

func (ri *RiemannIntegrator) IntegrateRiemann(min, max float64, f XYFunc) float64 {
	result := 0.

	numDivisions := 1000

	step := (max - min) / float64(numDivisions)

	for i := 0; i < numDivisions; i ++ {
		xNow := min  + step * float64(i)
		result += f(xNow)
	}

	return result * step
}

func (ri *RiemannIntegrator) Integrate(min, max float64, coefs ...float64) float64 {
    return ri.IntegrateRiemann(min, max, calcPoly(coefs...))
}

func Add(nums... int) int {
	//return 0
	var result int
	if len(nums) == 0 {
		println("no argumenhts provided")
		return 0
	}
	for _, i := range nums {
		result += i
	}
	return result
	//return 0

}

func Subtract(initial int, nums ...int) int {
	for _, i := range nums {
		initial -= i
	}
	return initial
}
//Polynomial Integrals  多项式积分
