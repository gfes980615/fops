package cmd

import "testing"

func Test_checksumFunc(t *testing.T) {
	type args struct {
		method string
		file   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "md5",
			args: args{
				method: "md5",
				file:   "checksum.go",
			},
			wantErr: false,
		},
		{
			name: "sha1",
			args: args{
				method: "sha1",
				file:   "checksum.go",
			},
			wantErr: false,
		},
		{
			name: "sha256",
			args: args{
				method: "sha256",
				file:   "checksum.go",
			},
			wantErr: false,
		},
		{
			name: "other case",
			args: args{
				method: "sha512",
				file:   "checksum.go",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checksumFunc(tt.args.method, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("checksumFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Logf("%s method is not provided yet", tt.args.method)
			} else {
				t.Logf("file checksum %s is: %s", tt.args.method, got)
			}
		})
	}
}
