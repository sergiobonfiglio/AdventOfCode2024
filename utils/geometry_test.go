package utils

import "testing"

func TestLine_IsOnLine(t *testing.T) {
	type fields struct {
		a *Cell
		b *Cell
	}
	type args struct {
		c []*Cell
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "horizontal true",
			fields: fields{
				a: NewCell(0, 0),
				b: NewCell(0, 2),
			},
			args: args{
				c: []*Cell{
					NewCell(0, 0),
					NewCell(0, 1),
					NewCell(0, 2),
					NewCell(0, 20),
					NewCell(0, -20),
				},
			},
			want: true,
		},
		{
			name: "horizontal false",
			fields: fields{
				a: NewCell(0, 0),
				b: NewCell(0, 2),
			},
			args: args{
				c: []*Cell{
					NewCell(1, 0),
					NewCell(2, 0),
					NewCell(-14, 2),
				},
			},
			want: false,
		},

		{
			name: "diag true",
			fields: fields{
				a: NewCell(0, 0),
				b: NewCell(2, 2),
			},
			args: args{
				c: []*Cell{
					NewCell(1, 1),
					NewCell(2, 2),
					NewCell(5, 5),
				},
			},
			want: true,
		},
		{
			name: "diag false",
			fields: fields{
				a: NewCell(0, 0),
				b: NewCell(2, 2),
			},
			args: args{
				c: []*Cell{
					NewCell(1, 2),
					NewCell(2, 3),
					NewCell(5, 4),
				},
			},
			want: false,
		},

		{
			name: "diag 2-4 true",
			fields: fields{
				a: NewCell(2, 4),
				b: NewCell(0, 0),
			},
			args: args{
				c: []*Cell{
					NewCell(4, 8),
					NewCell(8, 16),
					NewCell(1, 2),
					NewCell(3, 6),
					NewCell(5, 10),
				},
			},
			want: true,
		},
		{
			name: "diag 2-4 false",
			fields: fields{
				a: NewCell(2, 4),
				b: NewCell(0, 0),
			},
			args: args{
				c: []*Cell{
					NewCell(4, 9),
					NewCell(4, 7),
					NewCell(3, 8),
					NewCell(5, 8),
					NewCell(5, 9),
					NewCell(3, 7),
					NewCell(3, 9),
					NewCell(5, 7),
				},
			},
			want: false,
		},

		{
			name: "diag 1-3 true",
			fields: fields{
				a: NewCell(0, 0),
				b: NewCell(1, 3),
			},
			args: args{
				c: []*Cell{
					NewCell(2, 6),
					NewCell(3, 9),
					NewCell(4, 12),
				},
			},
			want: true,
		},
		{
			name: "diag 1-3 false",
			fields: fields{
				a: NewCell(0, 0),
				b: NewCell(1, 3),
			},
			args: args{
				c: []*Cell{
					NewCell(3, 8),
					NewCell(3, 7),
					NewCell(2, 9),
					NewCell(2, 9),
				},
			},
			want: false,
		},

		{
			name: "diag 1-3 false",
			fields: fields{
				a: NewCell(2, 5),
				b: NewCell(1, 8),
			},
			args: args{
				c: []*Cell{
					NewCell(0, 11),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLine(tt.fields.a, tt.fields.b)
			for _, c := range tt.args.c {
				if got := l.IsOnLine(c); got != tt.want {
					t.Errorf("IsOnLine(%v) = %v, want %v", c, got, tt.want)
				}
			}
		})
	}
}
