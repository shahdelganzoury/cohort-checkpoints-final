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
func hiddenRef(grid [][]int,called []int) int {if len(grid)==0||len(grid[0])==0{return -1};seen:=make(map[int]bool,len(called));for i,v:=range called{seen[v]=true;for _,row:=range grid{ok:=true;for _,n:=range row{if !seen[n]{ok=false;break}};if ok{return i+1}}};return -1}
func hiddenRandomString(rng *rand.Rand, chars string, n int) string {var b strings.Builder;for i:=0;i<n;i++{b.WriteByte(chars[rng.Intn(len(chars))])};return b.String()}
func hiddenRandomDigitPair(rng *rand.Rand)(string,string){n:=rng.Intn(80);return hiddenRandomString(rng,"0123456789",n),hiddenRandomString(rng,"0123456789",n)}
func hiddenRandomMissing(rng *rand.Rand)[]int{n:=rng.Intn(120)+2;a:=rng.Intn(n)+1;b:=rng.Intn(n-1)+1;if b>=a{b++};nums:=make([]int,0,n-2);for i:=1;i<=n;i++{if i!=a&&i!=b{nums=append(nums,i)}};rng.Shuffle(len(nums),func(i,j int){nums[i],nums[j]=nums[j],nums[i]});return nums}
func hiddenRandomStops(rng *rand.Rand)(int,[][2]int){capacity:=rng.Intn(60);onboard:=0;stops:=make([][2]int, rng.Intn(65));for i:=range stops{out:=rng.Intn(onboard+2);in:=rng.Intn(capacity+3);if rng.Intn(7)==0{out=-1};if rng.Intn(7)==0{in=-1};stops[i]=[2]int{in,out};if in<0||out<0||out>onboard||in>capacity-(onboard-out){onboard=0}else{onboard=onboard-out+in}};return capacity,stops}
func hiddenRandomMeetings(rng *rand.Rand)[][2]int {n:=rng.Intn(35);meetings:=make([][2]int,n);for i:=range meetings{s:=rng.Intn(1440);meetings[i]=[2]int{s,s+1+rng.Intn(1440-s)}};return meetings}
func hiddenRandomGrid(rng *rand.Rand)([][]int,[]int){rows,cols:=rng.Intn(4)+1,rng.Intn(5)+1;pool:=rng.Perm(300);grid:=make([][]int,rows);for r:=0;r<rows;r++{grid[r]=make([]int,cols);for c:=0;c<cols;c++{grid[r][c]=pool[r*cols+c]+1}};calls:=append([]int{},pool[:rows*cols]...);for i:=range calls{calls[i]++};for i:=0;i<rng.Intn(5);i++{calls=append(calls,500+i)};rng.Shuffle(len(calls),func(i,j int){calls[i],calls[j]=calls[j],calls[i]});calls=calls[:rng.Intn(len(calls)+1)];return grid,calls}

func TestHidden_FirstCompleteRow_CornerCases(t *testing.T){
t.Run("corner_1",func(t *testing.T){want:=hiddenRef([][]int{{4,1,7},{3,9,2},{8,5,6}}, []int{9,4,1,3,7,2});got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{4,1,7},{3,9,2},{8,5,6}}, []int{9,4,1,3,7,2})});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_2",func(t *testing.T){want:=hiddenRef([][]int{{1}}, []int{1});got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{1}}, []int{1})});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_3",func(t *testing.T){want:=hiddenRef([][]int{{1}}, []int{2});got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{1}}, []int{2})});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_4",func(t *testing.T){want:=hiddenRef([][]int{{1,2},{3,4}}, []int{1,3,2});got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{1,2},{3,4}}, []int{1,3,2})});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_5",func(t *testing.T){want:=hiddenRef([][]int{{1,2},{3,4}}, nil);got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{1,2},{3,4}}, nil)});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_6",func(t *testing.T){want:=hiddenRef([][]int{{1},{2}}, []int{2});got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{1},{2}}, []int{2})});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
t.Run("corner_7",func(t *testing.T){want:=hiddenRef([][]int{{1,2,3}}, []int{99,2,1,3});got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow([][]int{{1,2,3}}, []int{99,2,1,3})});if !reflect.DeepEqual(got,want){t.Fatalf("corner case mismatch: got %v want %v",got,want)}})
}
func TestHidden_FirstCompleteRow_SeededProperty(t *testing.T){rng:=rand.New(rand.NewSource(1009));for trial:=0;trial<500;trial++{grid, called := hiddenRandomGrid(rng);want:=hiddenRef(grid, called);got:=hiddenCallWithTimeout(t,func() int {return FirstCompleteRow(grid, called)});if !reflect.DeepEqual(got,want){t.Fatalf("trial %d: got %v want %v",trial,got,want)}}}
