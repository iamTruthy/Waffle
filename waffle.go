package waffle 

import (
  "github.com/iamTruthy/Coffee"
)

func Order () string {
	return "Waffles and Honey"
}

func Bill () string {
	return "Your bill is $50.99c"
}


func Meal () string {
 return coffee.WithLatte("You Ordered Waffles and Honey with a Cup of Latte")
}