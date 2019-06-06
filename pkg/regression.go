package main

type Point struct {
	X float64
	Y float64
}

type RecordTimeCalc func(float64) float64

// list of world records as (distance (m), time (s))
var points = []Point{
	{
		X: 100,
		Y: 9.58,
	},
	{
		X: 200,
		Y: 19.19,
	},
	{
		X: 400,
		Y: 43.03,
	},
	{
		X: 800,
		Y: 100.91,
	},
	{
		X: 1000,
		Y: 131.96,
	},
	{
		X: 1500,
		Y: 306,
	},
	{
		X: 2000,
		Y: 284.79,
	},
	{
		X: 3000,
		Y: 440.67,
	},
	{
		X: 5000,
		Y: 757.67,
	},
	{
		X: 10000,
		Y: 1577.53,
	},
	{
		X: 21097,
		Y: 3503,
	},
}

func linearRegressionLSE() RecordTimeCalc {

	q := len(points)

	if q == 0 {
		return func(dist float64) float64 {
			return 0.0
		}
	}

	p := float64(q)

	sum_x, sum_y, sum_xx, sum_xy := 0.0, 0.0, 0.0, 0.0

	for _, p := range points {
		sum_x += p.X
		sum_y += p.Y
		sum_xx += p.X * p.X
		sum_xy += p.X * p.Y
	}

	m := (p*sum_xy - sum_x*sum_y) / (p*sum_xx - sum_x*sum_x)
	b := (sum_y / p) - (m * sum_x / p)

	return func(dist float64) float64 {
		return dist*m + b
	}
}

func main() {}
