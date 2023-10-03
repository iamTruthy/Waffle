package waffle

import (
	"github.com/iamTruthy/Coffee"
)

func Order() string {
	return "Waffles with Honey"
}

func Meal() string {
	return coffee.WithLatte(Order() + " and a cup of Latte")
}

func Bill() string {
	return "Your bill is $50.99c"
}
