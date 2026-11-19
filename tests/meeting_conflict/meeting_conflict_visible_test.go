package main
import ("reflect";"testing")
var _ = reflect.DeepEqual
func TestLevel_MeetingConflict_Visible(t *testing.T){
t.Run("example_1",func(t *testing.T){got:=MeetingConflict(nil);want:=false;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_2",func(t *testing.T){got:=MeetingConflict([][2]int{{540,600},{570,630},{660,720}});want:=true;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_3",func(t *testing.T){got:=MeetingConflict([][2]int{{540,600},{600,660}});want:=false;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
t.Run("example_4",func(t *testing.T){got:=MeetingConflict([][2]int{{0,1440},{1,2}});want:=true;if !reflect.DeepEqual(got,want){t.Fatalf("got %v, want %v",got,want)}})
}
