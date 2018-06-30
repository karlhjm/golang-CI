package service

import (
	"reflect"
	"testing"

	"github.com/moandy/canyonsysu/entity"
)

func TestCommentRegister(t *testing.T) {
	type args struct {
		order_id      int
		rating_star   int
		rate_at       string
		username      string
		tags          string
		buddha_src    string
		client_text   string
		merchant_text string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	//{"case 0", args{1, 5, "20180630", "huangjm", "好吃", "touxiang.jpg", "你家店不错", "谢谢惠顾"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommentRegister(tt.args.order_id, tt.args.rating_star, tt.args.rate_at, tt.args.username, tt.args.tags, tt.args.buddha_src, tt.args.client_text, tt.args.merchant_text)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CommentRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllComments(t *testing.T) {
	tests := []struct {
		name string
		want []entity.Comment
	}{
	// TODO: Add test cases.
	//{"case 0", []entity.Comment{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllComments(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommentCountByTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	//{"case 0", args{"好吃"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCommentCountByTag(tt.args.tag); got != tt.want {
				t.Errorf("GetCommentCountByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAllTags(t *testing.T) { //ListAllTags()需要额外返回bool
	tests := []struct {
		name string
		want []entity.Tags
	}{
	// TODO: Add test cases.
	//{"case 0", []entity.Tags{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListCommentsByCount(t *testing.T) {
	type args struct {
		begin  int
		offset int
	}
	tests := []struct {
		name string
		args args
		want []entity.Comment
	}{
	// TODO: Add test cases.
	//{"case 0", args{0, 5}, []entity.Comment{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListCommentsByCount(tt.args.begin, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCommentsByCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
