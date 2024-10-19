package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"math"
	"math/rand"
	"slices"
	"strconv"
)

type Generator interface {
	Generate() float64
	GetDistribution(float64) float64
}

type FirstGenerator struct {
	lambda float64
}

func NewFirstGenerator(lambda float64) *FirstGenerator {
	return &FirstGenerator{lambda: lambda}
}

func (f *FirstGenerator) Generate() float64 {
	return (-1 / f.lambda) * math.Log(rand.Float64())
}

func (f *FirstGenerator) GetDistribution(x float64) float64 {
	return 1 - math.Exp(-f.lambda*x)
}

type SecondGenerator struct {
	alpha float64
	sigma float64
}

func NewSecondGenerator(alpha, sigma float64) *SecondGenerator {
	return &SecondGenerator{alpha: alpha, sigma: sigma}
}

func (s *SecondGenerator) Generate() float64 {
	u := 0.0
	for i := 1; i < 13; i++ {
		u += rand.Float64()
	}

	u -= 6
	return s.alpha + s.sigma*u
}

func (s *SecondGenerator) GetDistribution(x float64) float64 {
	return (1 + math.Erf((x-s.alpha)/(s.sigma*math.Sqrt2))) / 2
}

type ThirdGenerator struct {
	a float64
	c float64
	z float64
}

func NewThirdGenerator(a, c, z float64) *ThirdGenerator {
	return &ThirdGenerator{a: a, c: c, z: z}
}

func (t *ThirdGenerator) Generate() float64 {
	t.z = math.Mod(t.a*t.z, t.c)
	return t.z / t.c
}

func (t *ThirdGenerator) GetDistribution(x float64) float64 {
	if x < 0 {
		return 0
	} else if x > 1 {
		return 1
	} else {
		return x
	}
}

//--------------------

func main() {
	generators := []Generator{
		NewFirstGenerator(10),
		NewFirstGenerator(20),
		NewFirstGenerator(30),
		NewSecondGenerator(0, 3),
		NewSecondGenerator(22, 5),
		NewSecondGenerator(44, 7),
		NewThirdGenerator(math.Pow(5, 13), math.Pow(2, 31), 1),
		NewThirdGenerator(math.Pow(2, 31), math.Pow(7, 13), 1),
		NewThirdGenerator(math.Pow(2, 4), math.Pow(5, 13), 1),
	}

	values := make([][]float64, len(generators))

	forPlotValues := make([]plotter.Values, len(generators))

	for i := 0; i < 10000; i++ {
		for j, generator := range generators {
			v := generator.Generate()
			values[j] = append(values[j], v)
			forPlotValues[j] = append(forPlotValues[j], v)
		}
	}

	for i, gen := range generators {
		avg := getAvg(values[i])
		variance := getVar(values[i])
		xi := getXi(values[i], gen)

		fmt.Printf("Xi^2 for generator %d: %f\n", i+1, xi)
		fmt.Printf("Avg for generator %d: %f\n", i+1, avg)
		fmt.Printf("Variance for generator %d: %f\n", i+1, variance)
		fmt.Println("-------------------------------------------------")
	}

	for i := range generators {
		makePlot(forPlotValues[i], i)
	}

}

func getVar(values []float64) interface{} {
	avg := getAvg(values)
	sum := 0.0
	for _, v := range values {
		sum += math.Pow(v-avg, 2)
	}
	return sum / float64(len(values))
}

func getAvg(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func getXi(m []float64, gen Generator) float64 {
	intervalCount := 20
	intervals, intervalSize := getIntervals(m, intervalCount)

	minValue := slices.Min(m)

	xi := 0.0
	for i, fo := range intervals {
		left := minValue + float64(i)*intervalSize
		right := minValue + float64(i+1)*intervalSize

		fe := float64(len(m)) * (gen.GetDistribution(right) - gen.GetDistribution(left))
		xi += math.Pow(float64(fo)-fe, 2) / fe
	}
	return xi
}

func getIntervals(m []float64, count int) ([]int, float64) {
	minV := m[0]
	maxV := m[0]
	for _, v := range m {
		if v < minV {
			minV = v
		}
		if v > maxV {
			maxV = v
		}
	}

	intervalSize := (maxV - minV) / float64(count)
	intervals := make([]int, count)

	for _, v := range m {
		if v == maxV {
			intervals[count-1]++
			continue
		}
		intervals[int((v-minV)/intervalSize)]++
	}

	return intervals, intervalSize
}

func makePlot(values plotter.Values, i int) {
	p := plot.New()
	p.Title.Text = "histogram plot"

	hist, err := plotter.NewHist(values, 20)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	filename := "hist" + strconv.Itoa(i+1) + ".png"

	if err := p.Save(3*vg.Inch, 3*vg.Inch, filename); err != nil {
		panic(err)
	}
}
