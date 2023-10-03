package waffle 

import (
  "github.com/iamTruthy/Coffee"
)

func Order () string {
	return "Waffles with Honey"
}

func Bill () string {
	return "Your bill is $50.99c"
}


func Meal () string {
 return coffee.WithLatte("You Ordered " + Order() + " and a cup of Latte")
}