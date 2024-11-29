package mnemonic

import "testing"

func TestGenerateMnemonic(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "asdd",
			want:    "sdada",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateMnemonic()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateMnemonic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateMnemonic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateMnemonic(t *testing.T) {
	type args struct {
		mnemonic string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateMnemonic(tt.args.mnemonic); got != tt.want {
				t.Errorf("ValidateMnemonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
