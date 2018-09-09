package blog

import (
	"testing"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas porttitor congue massa. Fusce posuere, magna sed pulvinar ultricies, purus lectus malesuada libero, sit amet commodo magna eros quis urna. Nunc viverra imperdiet enim. Fusce est. Vivamus a tellus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Proin pharetra nonummy pede. Mauris et orci."

func TestLinterText(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Perfect",
			args{
				lorem,
			},
			lorem,
		},
		{
			"Multiple spaces",
			args{
				" Lorem  ipsum      dolor    ",
			},
			" Lorem ipsum dolor ",
		},
		{
			"Phrase.",
			args{
				"Lorem ipsum dolor.",
			},
			"Lorem ipsum dolor.",
		},
		{
			"Phrase .",
			args{
				"Lorem ipsum dolor . sit amet",
			},
			"Lorem ipsum dolor. sit amet",
		},
		{
			"Phrase . Phase",
			args{
				"Lorem ipsum dolor . sit amet",
			},
			"Lorem ipsum dolor. sit amet",
		},
		{
			"Multiple space Phrase.",
			args{
				"Lorem ipsum   dolor   .  sit amet",
			},
			"Lorem ipsum dolor. sit amet",
		},
		{
			"Multiple space Phrase,",
			args{
				"Lorem ipsum   dolor   ,  sit amet",
			},
			"Lorem ipsum dolor, sit amet",
		},
		{
			"No space after comma",
			args{
				"Lorem ipsum dolor,sit amet",
			},
			"Lorem ipsum dolor, sit amet",
		},
		{
			" space before and after comma",
			args{
				"Lorem ipsum dolor , sit amet",
			},
			"Lorem ipsum dolor, sit amet",
		},
		{
			"No space after dot",
			args{
				"Lorem ipsum dolor.Sit amet",
			},
			"Lorem ipsum dolor. Sit amet",
		},
		{
			"Phrase:Phrase",
			args{
				"Lorem ipsum dolor:sit amet",
			},
			"Lorem ipsum dolor\u00A0: sit amet",
		},
		{
			"Phrase : Phrase",
			args{
				"Lorem ipsum dolor : sit amet",
			},
			"Lorem ipsum dolor\u00A0: sit amet",
		},

		{
			"Phrase :",
			args{
				"Lorem ipsum dolor:",
			},
			"Lorem ipsum dolor\u00A0:",
		},
		{
			"Phrase ?",
			args{
				"Lorem ipsum dolor?",
			},
			"Lorem ipsum dolor\u00A0?",
		},
		{
			"Phrase..",
			args{
				"Lorem ipsum dolor..",
			},
			"Lorem ipsum dolor.",
		},
		{
			"Phrase...",
			args{
				"Lorem ipsum dolor...",
			},
			"Lorem ipsum dolor…",
		},
		{
			"Phrase ? Phrase.",
			args{
				"Lorem ipsum dolor?Sit amet, consectetur adipiscing elit.",
			},
			"Lorem ipsum dolor\u00A0? Sit amet, consectetur adipiscing elit.",
		},
		{
			"Phrase ? Phrase!",
			args{
				"Lorem ipsum dolor?Sit amet, consectetur adipiscing elit!",
			},
			"Lorem ipsum dolor\u00A0? Sit amet, consectetur adipiscing elit\u00A0!",
		},
		{
			"Phrase ? Phrase !",
			args{
				"Lorem ipsum dolor?Sit amet, consectetur adipiscing elit !",
			},
			"Lorem ipsum dolor\u00A0? Sit amet, consectetur adipiscing elit\u00A0!",
		},
		{
			"Phrase ? Phrase !    ",
			args{
				"Lorem ipsum dolor?Sit amet, consectetur adipiscing elit !    ",
			},
			"Lorem ipsum dolor\u00A0? Sit amet, consectetur adipiscing elit\u00A0!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextLinter(tt.args.t); got != tt.want {
				t.Errorf("LinterText() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestLinterAllCap(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Phrase ? Phrase !    ",
			args{
				"Lorem ipsum dolor?Sit amet, consectetur adipiscing elit !    ",
			},
			"LOREM IPSUM DOLOR ? SIT AMET, CONSECTETUR ADIPISCING ELIT !",
		},
		{
			"Phrase à Phrase !    ",
			args{
				"Phrase à Phrase !    ",
			},
			"PHRASE À PHRASE !",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllCapLinter(tt.args.t); got != tt.want {
				t.Errorf("LinterAllCap() = %v, want %v", got, tt.want)
			}
		})
	}
}
