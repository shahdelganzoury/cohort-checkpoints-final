package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_BusCapacity_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=BusCapacity(20, [][2]int{{10,0},{8,3},{2,5},{4,1}});want:=15;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=BusCapacity(10, [][2]int{{10,0}});want:=10;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=BusCapacity(10, [][2]int{{11,0}});want:=-1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=BusCapacity(10, [][2]int{{0,1}});want:=-1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
