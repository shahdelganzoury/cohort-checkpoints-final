package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_CashRegister_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=CashRegister(0);want:=0;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=CashRegister(1);want:=1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=CashRegister(5);want:=1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=CashRegister(10);want:=1;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
