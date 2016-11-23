// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: IPlayerService.go
// DO NOT EDIT!

package steam

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *OGResponse) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *OGResponse) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"game_count":`)
	fflib.FormatBits2(buf, uint64(mj.GameCount), 10, mj.GameCount < 0)
	buf.WriteString(`,"games":`)
	if mj.Games != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Games {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_OGResponsebase = iota
	ffj_t_OGResponseno_such_key

	ffj_t_OGResponse_GameCount

	ffj_t_OGResponse_Games
)

var ffj_key_OGResponse_GameCount = []byte("game_count")

var ffj_key_OGResponse_Games = []byte("games")

func (uj *OGResponse) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *OGResponse) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_OGResponsebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_OGResponseno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'g':

					if bytes.Equal(ffj_key_OGResponse_GameCount, kn) {
						currentKey = ffj_t_OGResponse_GameCount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_OGResponse_Games, kn) {
						currentKey = ffj_t_OGResponse_Games
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_OGResponse_Games, kn) {
					currentKey = ffj_t_OGResponse_Games
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_OGResponse_GameCount, kn) {
					currentKey = ffj_t_OGResponse_GameCount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_OGResponseno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_OGResponse_GameCount:
					goto handle_GameCount

				case ffj_t_OGResponse_Games:
					goto handle_Games

				case ffj_t_OGResponseno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_GameCount:

	/* handler: uj.GameCount type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.GameCount = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Games:

	/* handler: uj.Games type=[]steam.UserGame kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Games = nil
		} else {

			uj.Games = make([]UserGame, 0)

			wantVal := true

			for {

				var tmp_uj__Games UserGame

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Games type=steam.UserGame kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = tmp_uj__Games.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.Games = append(uj.Games, tmp_uj__Games)
				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *OwnedGameResponse) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *OwnedGameResponse) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"response":`)

	{

		err = mj.OGResponse.MarshalJSONBuf(buf)
		if err != nil {
			return err
		}

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_OwnedGameResponsebase = iota
	ffj_t_OwnedGameResponseno_such_key

	ffj_t_OwnedGameResponse_OGResponse
)

var ffj_key_OwnedGameResponse_OGResponse = []byte("response")

func (uj *OwnedGameResponse) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *OwnedGameResponse) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_OwnedGameResponsebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_OwnedGameResponseno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'r':

					if bytes.Equal(ffj_key_OwnedGameResponse_OGResponse, kn) {
						currentKey = ffj_t_OwnedGameResponse_OGResponse
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_OwnedGameResponse_OGResponse, kn) {
					currentKey = ffj_t_OwnedGameResponse_OGResponse
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_OwnedGameResponseno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_OwnedGameResponse_OGResponse:
					goto handle_OGResponse

				case ffj_t_OwnedGameResponseno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_OGResponse:

	/* handler: uj.OGResponse type=steam.OGResponse kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			state = fflib.FFParse_after_value
			goto mainparse
		}

		err = uj.OGResponse.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
		if err != nil {
			return err
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *UserGame) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *UserGame) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "appid":`)
	fflib.FormatBits2(buf, uint64(mj.AppID), 10, mj.AppID < 0)
	buf.WriteByte(',')
	if len(mj.Name) != 0 {
		buf.WriteString(`"name":`)
		fflib.WriteJsonString(buf, string(mj.Name))
		buf.WriteByte(',')
	}
	buf.WriteString(`"playtime_2weeks":`)
	fflib.FormatBits2(buf, uint64(mj.Playtime2Weeks), 10, mj.Playtime2Weeks < 0)
	buf.WriteString(`,"playtime_forever":`)
	fflib.FormatBits2(buf, uint64(mj.PlaytimeForever), 10, mj.PlaytimeForever < 0)
	buf.WriteByte(',')
	if len(mj.ImgIconUrl) != 0 {
		buf.WriteString(`"img_icon_url":`)
		fflib.WriteJsonString(buf, string(mj.ImgIconUrl))
		buf.WriteByte(',')
	}
	if len(mj.ImgLogoUrl) != 0 {
		buf.WriteString(`"img_logo_url":`)
		fflib.WriteJsonString(buf, string(mj.ImgLogoUrl))
		buf.WriteByte(',')
	}
	if mj.VisibleStats != false {
		if mj.VisibleStats {
			buf.WriteString(`"has_community_visible_stats":true`)
		} else {
			buf.WriteString(`"has_community_visible_stats":false`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_UserGamebase = iota
	ffj_t_UserGameno_such_key

	ffj_t_UserGame_AppID

	ffj_t_UserGame_Name

	ffj_t_UserGame_Playtime2Weeks

	ffj_t_UserGame_PlaytimeForever

	ffj_t_UserGame_ImgIconUrl

	ffj_t_UserGame_ImgLogoUrl

	ffj_t_UserGame_VisibleStats
)

var ffj_key_UserGame_AppID = []byte("appid")

var ffj_key_UserGame_Name = []byte("name")

var ffj_key_UserGame_Playtime2Weeks = []byte("playtime_2weeks")

var ffj_key_UserGame_PlaytimeForever = []byte("playtime_forever")

var ffj_key_UserGame_ImgIconUrl = []byte("img_icon_url")

var ffj_key_UserGame_ImgLogoUrl = []byte("img_logo_url")

var ffj_key_UserGame_VisibleStats = []byte("has_community_visible_stats")

func (uj *UserGame) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *UserGame) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_UserGamebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_UserGameno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_UserGame_AppID, kn) {
						currentKey = ffj_t_UserGame_AppID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_UserGame_VisibleStats, kn) {
						currentKey = ffj_t_UserGame_VisibleStats
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_UserGame_ImgIconUrl, kn) {
						currentKey = ffj_t_UserGame_ImgIconUrl
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_UserGame_ImgLogoUrl, kn) {
						currentKey = ffj_t_UserGame_ImgLogoUrl
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffj_key_UserGame_Name, kn) {
						currentKey = ffj_t_UserGame_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_UserGame_Playtime2Weeks, kn) {
						currentKey = ffj_t_UserGame_Playtime2Weeks
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_UserGame_PlaytimeForever, kn) {
						currentKey = ffj_t_UserGame_PlaytimeForever
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_UserGame_VisibleStats, kn) {
					currentKey = ffj_t_UserGame_VisibleStats
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_UserGame_ImgLogoUrl, kn) {
					currentKey = ffj_t_UserGame_ImgLogoUrl
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_UserGame_ImgIconUrl, kn) {
					currentKey = ffj_t_UserGame_ImgIconUrl
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_UserGame_PlaytimeForever, kn) {
					currentKey = ffj_t_UserGame_PlaytimeForever
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_UserGame_Playtime2Weeks, kn) {
					currentKey = ffj_t_UserGame_Playtime2Weeks
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_UserGame_Name, kn) {
					currentKey = ffj_t_UserGame_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_UserGame_AppID, kn) {
					currentKey = ffj_t_UserGame_AppID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_UserGameno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_UserGame_AppID:
					goto handle_AppID

				case ffj_t_UserGame_Name:
					goto handle_Name

				case ffj_t_UserGame_Playtime2Weeks:
					goto handle_Playtime2Weeks

				case ffj_t_UserGame_PlaytimeForever:
					goto handle_PlaytimeForever

				case ffj_t_UserGame_ImgIconUrl:
					goto handle_ImgIconUrl

				case ffj_t_UserGame_ImgLogoUrl:
					goto handle_ImgLogoUrl

				case ffj_t_UserGame_VisibleStats:
					goto handle_VisibleStats

				case ffj_t_UserGameno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_AppID:

	/* handler: uj.AppID type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.AppID = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: uj.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Playtime2Weeks:

	/* handler: uj.Playtime2Weeks type=int32 kind=int32 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int32", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Playtime2Weeks = int32(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PlaytimeForever:

	/* handler: uj.PlaytimeForever type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.PlaytimeForever = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ImgIconUrl:

	/* handler: uj.ImgIconUrl type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ImgIconUrl = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ImgLogoUrl:

	/* handler: uj.ImgLogoUrl type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ImgLogoUrl = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VisibleStats:

	/* handler: uj.VisibleStats type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.VisibleStats = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.VisibleStats = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
