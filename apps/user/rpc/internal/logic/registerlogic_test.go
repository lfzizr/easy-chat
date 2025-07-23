package logic

import (
	"context"
	"reflect"
	"testing"

	"gitee.com/lfzizr/easy-chat/apps/user/rpc/internal/svc"
	"gitee.com/lfzizr/easy-chat/apps/user/rpc/user"
)

func TestNewRegisterLogic(t *testing.T) {
	type args struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	tests := []struct {
		name string
		args args
		want *RegisterLogic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegisterLogic(tt.args.ctx, tt.args.svcCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRegisterLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterLogic_Register(t *testing.T) {
	
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name    string
		args    args
		wantPrint bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name:"1",args:args{in:&user.RegisterReq{
			Phone:"13800000001",
			Nickname:"1",
			Password:"1",
			Avatar:"1",
			Sex:1,
		}},wantPrint:true,wantErr:false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterLogic.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name,got)
			}
		})
	}
}
