package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_PasswordStrength_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=PasswordStrength("");want:=0;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=PasswordStrength("abc");want:=0;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=PasswordStrength("Go2026!");want:=3;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=PasswordStrength("GoLang2026!");want:=4;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
