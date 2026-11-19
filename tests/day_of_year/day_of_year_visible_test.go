package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_DayOfYear_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=DayOfYear(1,3,2024);want:=61;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=DayOfYear(1,3,2025);want:=60;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=DayOfYear(29,2,2025);want:=-1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=DayOfYear(29,2,2024);want:=60;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
