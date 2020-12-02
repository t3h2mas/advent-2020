package passwords

import (
	"reflect"
	"testing"
)

func Test_corporatePasswordFrom(t *testing.T) {
	passwordEntry := "1-3 b: cdefg"

	expected := &CorporatePassword{
		password: "cdefg",
		policy: &PasswordPolicy{
			min:  1,
			max:  3,
			char: 'b',
		},
	}

	got := corporatePasswordFrom(passwordEntry)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("corporatePolicyFrom got() = %+v, want %+v", got, expected)
	}
}

func TestCorporatePassword_IsValid(t *testing.T) {
	type fields struct {
		password string
		policy   *PasswordPolicy
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid corporate password",
			fields: fields{
				password: "abcde",
				policy: &PasswordPolicy{
					min:  1,
					max:  3,
					char: 'a',
				},
			},
			want: true,
		},
		{
			name: "invalid corporate password",
			fields: fields{
				password: "cdefg",
				policy: &PasswordPolicy{
					min:  1,
					max:  3,
					char: 'b',
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &CorporatePassword{
				password: tt.fields.password,
				policy:   tt.fields.policy,
			}
			if got := cp.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
