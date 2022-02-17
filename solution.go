package square

import "math"

// Определить пользовательский тип int для хранения номера сторон и обновить подпись CalcSquare, заменив #yourTypeNameHere#

// Определяем константы для представления сторон 0, 3 и 4. В тесте используются мнемосхемы: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// это как:
// ВычислитьКвадрат(10.0, СтороныТреугольника)
// ВычислитьКвадрат(10.0, СтороныКвадрат)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum float64) float64 {
	switch sidesNum {
	case 3:
		return (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4

	case 4:
		return math.Pow(sideLen, 2)

	case 0:
		return math.Pi * math.Pow(sideLen, 2)

	default:
		return 0

	}
}
