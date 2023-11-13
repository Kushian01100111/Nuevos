package kata

func PositiveSum(numbers []int) (sum int) {
  for i := 0; i < len(numbers); i++{
    if numbers[i] > 0{
      sum += 0 + numbers[i] 
    }else{
      sum += 0
    }
  }
  return
}