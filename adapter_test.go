package patterns

import "testing"

func TestTranscoderAdapter(t *testing.T) {

	var tests = []struct {
		name     string
		fileType string
		in       string
		out      string
	}{
		{"React Transcoder", "react", "Hello", "From ReactTranscode: Hello"},
		{"Vue Transcoder", "vue", "Moi", "From VueTranscode: Moi"},
		{"Angular Transcoder", "angular", "Hola", "From AngularTranscode: Hola"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			transcoder := &TranscoderAdapter{fileType: test.fileType}
			got := transcoder.Transcode(test.in)
			if got != test.out {
				t.Errorf("got %v, want %v", got, test.out)
			}
		})
	}
}
