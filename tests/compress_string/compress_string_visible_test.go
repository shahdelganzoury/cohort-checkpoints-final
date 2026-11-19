package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_CompressString_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=CompressString("");want:="";if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=CompressString("a");want:="a1";if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=CompressString("aaabbcccc");want:="a3b2c4";if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=CompressString("aabbaa");want:="a2b2a2";if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
