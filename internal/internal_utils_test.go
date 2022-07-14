package internal

import "testing"

func Test_selectorToGoName(t *testing.T) {
	type args struct {
		sel string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{sel: "menuWillOpen:"},
			want: "MenuWillOpen",
		},
		{
			args: args{sel: "menu:updateItem:atIndex:shouldCancel:"},
			want: "Menu_UpdateItem_AtIndex_ShouldCancel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectorToGoName(tt.args.sel); got != tt.want {
				t.Errorf("selectorToGoName() = %v, want %v", got, tt.want)
			}
		})
	}
}
