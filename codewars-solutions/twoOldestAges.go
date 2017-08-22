//Problem link: https://www.codewars.com/kata/511f11d355fe575d2c000001
//You'll need an account to follow the link successfully

package kata

func TwoOldestAges(ages []int) [2]int {
  var oldestAge, secondOldestAge int
  for _, v := range ages {
    if v > oldestAge {
      secondOldestAge = oldestAge
      oldestAge = v
    } else if v > secondOldestAge {
      secondOldestAge = v
    }
  }
  return [2]int{secondOldestAge, oldestAge}
}