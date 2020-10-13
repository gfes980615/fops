package cmd

import "testing"

func Test_fileLineCount(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: `\n count`,
			args: args{
				file: "test/for_test.txt",
			},
			want:    7,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fileLineCount(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileLineCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fileLineCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readFileContent(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "read file",
			args: args{
				file: "test/for_test.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFileContent(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("file content: \n%s", got)
		})
	}
}
