//go:build hidden
// +build hidden

package main
import (
 "fmt"
 "math/rand"
 "reflect"
 "strconv"
 "strings"
 "testing"
 "time"
)
var (_ = fmt.Sprint; _ = rand.New; _ = reflect.DeepEqual; _ = strconv.Itoa; _ = strings.Builder{}; _ = time.Second)
func hiddenCallWithTimeout[T any](t *testing.T, fn func() T) T {
 t.Helper()
 ch := make(chan T,1)
 panicCh := make(chan any,1)
 go func(){ defer func(){if v:=recover();v!=nil {panicCh<-v}}(); ch<-fn() }()
 select {case got:=<-ch:return got;case p:=<-panicCh:t.Fatalf("student function panicked: %v",p);var zero T;return zero;case <-time.After(800*time.Millisecond):t.Fatal("student function timed out");var zero T;return zero}
}
func hiddenRef(s string) string {if s==""{return ""};var b strings.Builder;start:=0;for i:=1;i<=len(s);i++{if i==len(s)||s[i]!=s[start]{b.WriteByte(s[start]);b.WriteString(strconv.Itoa(i-start));start=i}};return b.String()}
func hiddenRandomString(rng *rand.Rand, chars string, n int) string {var b strings.Builder;for i:=0;i<n;i++{b.WriteByte(chars[rng.Intn(len(chars))])};return b.String()}
func hiddenRandomDigitPair(rng *rand.Rand)(string,string){n:=rng.Intn(80);return hiddenRandomString(rng,"0123456789",n),hiddenRandomString(rng,"0123456789",n)}
func hiddenRandomMissing(rng *rand.Rand)[]int{n:=rng.Intn(120)+2;a:=rng.Intn(n)+1;b:=rng.Intn(n-1)+1;if b>=a{b++};nums:=make([]int,0,n-2);for i:=1;i<=n;i++{if i!=a&&i!=b{nums=append(nums,i)}};rng.Shuffle(len(nums),func(i,j int){nums[i],nums[j]=nums[j],nums[i]});return nums}
func hiddenRandomStops(rng *rand.Rand)(int,[][2]int){capacity:=rng.Intn(60);onboard:=0;stops:=make([][2]int, rng.Intn(65));for i:=range stops{out:=rng.Intn(onboard+2);in:=rng.Intn(capacity+3);if rng.Intn(7)==0{out=-1};if rng.Intn(7)==0{in=-1};stops[i]=[2]int{in,out};if in<0||out<0||out>onboard||in>capacity-(onboard-out){onboard=0}else{onboard=onboard-out+in}};return capacity,stops}
func hiddenRandomMeetings(rng *rand.Rand)[][2]int {n:=rng.Intn(35);meetings:=make([][2]int,n);for i:=range meetings{s:=rng.Intn(1440);meetings[i]=[2]int{s,s+1+rng.Intn(1440-s)}};return meetings}
func hiddenRandomGrid(rng *rand.Rand)([][]int,[]int){rows,cols:=rng.Intn(4)+1,rng.Intn(5)+1;pool:=rng.Perm(300);grid:=make([][]int,rows);for r:=0;r<rows;r++{grid[r]=make([]int,cols);for c:=0;c<cols;c++{grid[r][c]=pool[r*cols+c]+1}};calls:=append([]int{},pool[:rows*cols]...);for i:=range calls{calls[i]++};for i:=0;i<rng.Intn(5);i++{calls=append(calls,500+i)};rng.Shuffle(len(calls),func(i,j int){calls[i],calls[j]=calls[j],calls[i]});calls=calls[:rng.Intn(len(calls)+1)];return grid,calls}

func TestHidden_CompressString_CornerCases(t *testing.T){
t.Run("corner_1",func(t *testing.T){want:=hiddenRef("");got:=hiddenCallWithTimeout(t,func() string {return CompressString("")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_2",func(t *testing.T){want:=hiddenRef("a");got:=hiddenCallWithTimeout(t,func() string {return CompressString("a")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_3",func(t *testing.T){want:=hiddenRef("aaabbcccc");got:=hiddenCallWithTimeout(t,func() string {return CompressString("aaabbcccc")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_4",func(t *testing.T){want:=hiddenRef("aabbaa");got:=hiddenCallWithTimeout(t,func() string {return CompressString("aabbaa")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_5",func(t *testing.T){want:=hiddenRef("abcd");got:=hiddenCallWithTimeout(t,func() string {return CompressString("abcd")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_6",func(t *testing.T){want:=hiddenRef("aaaaaaaaaaaa");got:=hiddenCallWithTimeout(t,func() string {return CompressString("aaaaaaaaaaaa")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_7",func(t *testing.T){want:=hiddenRef("ababa");got:=hiddenCallWithTimeout(t,func() string {return CompressString("ababa")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_8",func(t *testing.T){want:=hiddenRef("zzzzzzzzzzzzzzzzz");got:=hiddenCallWithTimeout(t,func() string {return CompressString("zzzzzzzzzzzzzzzzz")});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
}
func TestHidden_CompressString_SeededProperty(t *testing.T){rng:=rand.New(rand.NewSource(1005));for trial:=0;trial<500;trial++{s := hiddenRandomString(rng,"aabbccddef",rng.Intn(101));want:=hiddenRef(s);got:=hiddenCallWithTimeout(t,func() string {return CompressString(s)});if !reflect.DeepEqual(got,want){t.Fatalf("trial %d: got %v want %v",trial,got,want)}}}
