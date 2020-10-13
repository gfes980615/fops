package helper

import (
	"io/ioutil"
	"testing"
)

func TestCheckFileExist(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "file is not exist",
			args: args{
				"check",
			},
			wantErr: true,
		},
		{
			name: "file is exist ",
			args: args{
				"check.go",
			},
			wantErr: false,
		},
		{
			name: "file is exist in folder",
			args: args{
				"test/test_txt.txt",
			},
			wantErr: false,
		},
		{
			name: "get route is only folder",
			args: args{
				"/test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckFileExist(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("CheckFileExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckFileIsBinary(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file format:png",
			args: args{
				file: "test/test_png.png",
			},
			want: true,
		},
		{
			name: "file format:jpg",
			args: args{
				file: "test/test_jpg.jpg",
			},
			want: true,
		},
		{
			name: "file format:none",
			args: args{
				file: "test/fops",
			},
			want: true,
		},
		{
			name: "file format:txt",
			args: args{
				file: "test/test_txt.txt",
			},
			want: false,
		},
		{
			name: "file format:csv",
			args: args{
				file: "test/test_csv.csv",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := ioutil.ReadFile(tt.args.file)
			if got := CheckFileIsBinary(b); got != tt.want {
				t.Errorf("CheckFileIsBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distinguishFileOrFolder(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "recognize file case 1",
			args: args{
				f: "test_txt.txt",
			},
			want: File,
		},
		{
			name: "recognize file case 2",
			args: args{
				f: "/test/test_txt.txt",
			},
			want: File,
		},
		{
			name: "recognize folder",
			args: args{
				f: "/test",
			},
			want: Folder,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinguishFileOrFolder(tt.args.f); got != tt.want {
				t.Errorf("distinguishFileOrFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
