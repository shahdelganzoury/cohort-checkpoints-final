package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_FirstCompleteRow_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=FirstCompleteRow([][]int{{4,1,7},{3,9,2},{8,5,6}}, []int{9,4,1,3,7,2});want:=5;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=FirstCompleteRow([][]int{{1}}, []int{1});want:=1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=FirstCompleteRow([][]int{{1}}, []int{2});want:=-1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=FirstCompleteRow([][]int{{1,2},{3,4}}, []int{1,3,2});want:=3;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
