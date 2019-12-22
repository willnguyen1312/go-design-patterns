package patterns

// ReactTranscode struct
type ReactTranscode struct{}

// Write method for ReactTranscode struct
func (r *ReactTranscode) Write(s string) string {
	return "From ReactTranscode: " + s
}

// VueTranscode struct
type VueTranscode struct{}

// Convert method for VueTranscode struct
func (v *VueTranscode) Convert(s string) string {
	return "From VueTranscode: " + s
}

// AngularTranscode struct
type AngularTranscode struct{}

// Interprete method for AngularTranscode struct
func (v *AngularTranscode) Interprete(s string) string {
	return "From AngularTranscode: " + s
}

// Transcoder interface
type Transcoder interface {
	Transcode(string) string
}

// TranscoderAdapter struct
type TranscoderAdapter struct {
	fileType string
}

// Transcode method on TranscoderAdapter struct
func (trans *TranscoderAdapter) Transcode(s string) string {
	if trans.fileType == "react" {
		return (&ReactTranscode{}).Write(s)
	}

	if trans.fileType == "vue" {
		return (&VueTranscode{}).Convert(s)
	}

	if trans.fileType == "angular" {
		return (&AngularTranscode{}).Interprete(s)
	}

	return "Unrecognize filType " + s
}
