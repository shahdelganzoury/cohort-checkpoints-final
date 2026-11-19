package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_DigitPositionMatch_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=DigitPositionMatch("58342", "57382");want:=3;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=DigitPositionMatch("", "");want:=0;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=DigitPositionMatch("000", "000");want:=3;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=DigitPositionMatch("123", "456");want:=0;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
