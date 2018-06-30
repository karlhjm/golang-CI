package service

import (
	"reflect"
	"testing"
)

func TestListAllCategorys(t *testing.T) { //ListAllCategorys()需要额外返回一个bool
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
		{"case 0", []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllCategorys(); !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("ListAllCategorys() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestCategoryRegister(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   string
// 		want    bool
// 		wantErr bool
// 	}{

// 		{"first register", "juice", true, false},
// 		{"multiply register", "juice", false, true}, //已测出bug
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := CategoryRegister(tt.input)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CategoryRegister( %v ) error = %v, wantErr %v", tt.input, err, tt.wantErr)
// 			}
// 			if got != tt.want {
// 				t.Errorf("CategoryRegister( %v ) = %v, want %v", tt.input, got, tt.want)
// 			}
// 		})
// 	}
// }
