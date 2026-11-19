package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_TwoMissingNumbers_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=func() [2]int {a,b:=TwoMissingNumbers(nil);return [2]int{a,b}}();want:=[2]int{1,2};if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=func() [2]int {a,b:=TwoMissingNumbers([]int{1,3,6,2});return [2]int{a,b}}();want:=[2]int{4,5};if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=func() [2]int {a,b:=TwoMissingNumbers([]int{2,4,5,6});return [2]int{a,b}}();want:=[2]int{1,3};if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=func() [2]int {a,b:=TwoMissingNumbers([]int{2,3});return [2]int{a,b}}();want:=[2]int{1,4};if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
