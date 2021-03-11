// Code generated by goa v3.2.6, DO NOT EDIT.
//
// playback HTTP server types
//
// Command:
// $ goa gen github.com/lbryio/lbrytv/apps/watchman/design -o apps/watchman

package server

import (
	"unicode/utf8"

	playback "github.com/lbryio/lbrytv/apps/watchman/gen/playback"
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "playback" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// LBRY URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// Playback event stream position, ms
	Pos *uint64 `form:"pos,omitempty" json:"pos,omitempty" xml:"pos,omitempty"`
	// Playback event duration, ms
	Dur *uint32 `form:"dur,omitempty" json:"dur,omitempty" xml:"dur,omitempty"`
	// Buffering events count
	Bfc *uint32 `form:"bfc,omitempty" json:"bfc,omitempty" xml:"bfc,omitempty"`
	// Buffering events total duration, ms
	Bfd *uint64 `form:"bfd,omitempty" json:"bfd,omitempty" xml:"bfd,omitempty"`
	// Video format
	Fmt *string `form:"fmt,omitempty" json:"fmt,omitempty" xml:"fmt,omitempty"`
	// Player server name
	Pid *string `form:"pid,omitempty" json:"pid,omitempty" xml:"pid,omitempty"`
	// Client download rate, bits/s
	Crt *uint64 `form:"crt,omitempty" json:"crt,omitempty" xml:"crt,omitempty"`
	// Client area
	Car *string `form:"car,omitempty" json:"car,omitempty" xml:"car,omitempty"`
	// Unique client ID
	Cid *string `form:"cid,omitempty" json:"cid,omitempty" xml:"cid,omitempty"`
	// Client device
	Cdv *string `form:"cdv,omitempty" json:"cdv,omitempty" xml:"cdv,omitempty"`
}

// NewAddPlayback builds a playback service add endpoint payload.
func NewAddPlayback(body *AddRequestBody) *playback.Playback {
	v := &playback.Playback{
		URL: *body.URL,
		Pos: *body.Pos,
		Dur: *body.Dur,
		Bfc: *body.Bfc,
		Bfd: *body.Bfd,
		Fmt: *body.Fmt,
		Pid: *body.Pid,
		Crt: body.Crt,
		Car: body.Car,
		Cid: *body.Cid,
		Cdv: *body.Cdv,
	}

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	if body.Pos == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pos", "body"))
	}
	if body.Dur == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("dur", "body"))
	}
	if body.Bfc == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bfc", "body"))
	}
	if body.Bfd == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("bfd", "body"))
	}
	if body.Fmt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fmt", "body"))
	}
	if body.Pid == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pid", "body"))
	}
	if body.Cid == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cid", "body"))
	}
	if body.Cdv == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cdv", "body"))
	}
	if body.URL != nil {
		if utf8.RuneCountInString(*body.URL) > 512 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.url", *body.URL, utf8.RuneCountInString(*body.URL), 512, false))
		}
	}
	if body.Dur != nil {
		if *body.Dur < 1000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.dur", *body.Dur, 1000, true))
		}
	}
	if body.Dur != nil {
		if *body.Dur > 3.6e+06 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.dur", *body.Dur, 3.6e+06, false))
		}
	}
	if body.Fmt != nil {
		if !(*body.Fmt == "def" || *body.Fmt == "hls") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.fmt", *body.Fmt, []interface{}{"def", "hls"}))
		}
	}
	if body.Pid != nil {
		if utf8.RuneCountInString(*body.Pid) > 32 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.pid", *body.Pid, utf8.RuneCountInString(*body.Pid), 32, false))
		}
	}
	if body.Car != nil {
		if utf8.RuneCountInString(*body.Car) > 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.car", *body.Car, utf8.RuneCountInString(*body.Car), 3, false))
		}
	}
	if body.Cid != nil {
		if utf8.RuneCountInString(*body.Cid) > 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.cid", *body.Cid, utf8.RuneCountInString(*body.Cid), 64, false))
		}
	}
	if body.Cdv != nil {
		if !(*body.Cdv == "ios" || *body.Cdv == "and" || *body.Cdv == "web") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.cdv", *body.Cdv, []interface{}{"ios", "and", "web"}))
		}
	}
	return
}
