package badges_test

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/connorholt/explain_interfaces/badges"
)

func TestGitStar_IsEnabled(t *testing.T) {
	type fields struct {
		commitCount int
		prepare     func(reviewHandlerMock *Mockhandler)
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				commitCount: 200,
				prepare: func(reviewHandlerMock *Mockhandler) {
					reviewHandlerMock.EXPECT().IsEnabled().Return(true)
				},
			},
			want: true,
		},
		{
			name: "failed: less commitCount",
			fields: fields{
				commitCount: 1,
				prepare: func(reviewHandlerMock *Mockhandler) {
					reviewHandlerMock.EXPECT().IsEnabled().Return(true)
				},
			},
			want: false,
		},
		{
			name: "failed: reviewHandler is not ok",
			fields: fields{
				commitCount: 1000,
				prepare: func(reviewHandlerMock *Mockhandler) {
					reviewHandlerMock.EXPECT().IsEnabled().Return(false)
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			reviewHandlerMock := NewMockhandler(ctrl)

			tt.fields.prepare(reviewHandlerMock)

			b := badges.NewGitStar(tt.fields.commitCount, reviewHandlerMock)
			if got := b.IsEnabled(); got != tt.want {
				t.Errorf("IsEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
