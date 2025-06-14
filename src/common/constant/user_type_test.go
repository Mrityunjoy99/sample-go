package constant

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserType(t *testing.T) {
	tests := []struct {
		name     string
		userType string
		want     UserType
		wantErr  bool
	}{
		{
			name:     "valid_admin",
			userType: "ADMIN",
			want:     UserTypeAdmin,
			wantErr:  false,
		},
		{
			name:     "valid_manager",
			userType: "MANAGER",
			want:     UserTypeManager,
			wantErr:  false,
		},
		{
			name:     "valid_user",
			userType: "USER",
			want:     UserTypeUser,
			wantErr:  false,
		},
		{
			name:     "invalid_user_type",
			userType: "INVALID",
			want:     -1,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserType(tt.userType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
