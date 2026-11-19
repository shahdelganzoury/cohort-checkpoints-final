package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_DigitalRoot_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=DigitalRoot(0);want:=0;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=DigitalRoot(9);want:=9;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=DigitalRoot(10);want:=1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=DigitalRoot(99);want:=9;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
