package test_demo

import (
	"reflect"
	"testing"
)

//测试

func TestSplit(t *testing.T) {
	//got := Split("我打你", "打")
	//want := []string{"我", "你"}
	tests := []struct {
		name string
		s    string
		sep  string
		want []string
	}{
		{name: "test1", s: "hjdgaoh", sep: "j", want: []string{"h", "dgaoh"}},
		{name: "test2", s: "阿斯噶额我给", sep: "额", want: []string{"阿斯噶", "我给"}},
		{name: "test3", s: "12345432", sep: "1", want: []string{"", "2345432"}},
		{name: "test4", s: "hjdgaoh", sep: "h", want: []string{"", "jdgao", ""}},
		{name: "test5", s: "a:b:c:d", sep: ":", want: []string{"a", "b", "c", "d"}},
	}
	for _, tt := range tests {

		got := Split(tt.s, tt.sep)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("name:%v,want:%v, got:%v", tt.name, tt.want, got)
		}
	}

}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("上海自来水来自海上", "海")
	}
}
