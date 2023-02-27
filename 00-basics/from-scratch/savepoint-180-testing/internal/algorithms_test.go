package algorithms

import (
	"testing"
)

func TestFizzbuzz_Calculate(t *testing.T) {
	type fields struct {
		Min int
		Max int
	}
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Fizz",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 3,
			},
			want:    "Fizz",
			wantErr: false,
		},
		{
			name: "Buzz",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 5,
			},
			want:    "Buzz",
			wantErr: false,
		},
		{
			name: "FizzBuzz",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 15,
			},
			want:    "FizzBuzz",
			wantErr: false,
		},
		{
			name: "Number",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 7,
			},
			want:    "7",
			wantErr: false,
		},
		{
			name: "Number",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 7,
			},
			want:    "7",
			wantErr: false,
		},
		{
			name: "Number",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 7,
			},
			want:    "7",
			wantErr: false,
		},
		{
			name: "Number",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 7,
			},
			want:    "7",
			wantErr: false,
		},
		{
			name: "Number",
			fields: fields{
				Min: 1,
				Max: 20,
			},
			args: args{
				number: 21,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Fizzbuzz{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			got, err := f.Calculate(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fizzbuzz.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fizzbuzz.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJamBoo_Calculate(t *testing.T) {
	type fields struct {
	}
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Jam",
			fields: fields{},
			args: args{
				number: 3,
			},
			want:    "Jam",
			wantErr: false,
		},
		{
			name:   "Boo",
			fields: fields{},
			args: args{
				number: 5,
			},
			want:    "Boo",
			wantErr: false,
		},
		{
			name:   "JamBoo",
			fields: fields{},
			args: args{
				number: 15,
			},
			want:    "JamBoo",
			wantErr: false,
		},
		{
			name:   "Number",
			fields: fields{},
			args: args{
				number: 7,
			},
			want:    "7",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Jamboo{}
			got, err := f.Calculate(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("Jamboo.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Jamboo.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
