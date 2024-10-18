package cores

import (
	"fmt"
	"testing"
)

func TestJwtClaims_ToJwtClaimsAccessData(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiJ0aGlzIGlzIGp3dCBpZGVudGl0eSIsInNpZCI6InRoaXMgaXMgc2Vzc2lvbiBpZCIsImV4cCI6NDc2OTI0MDkzMSwiaXNzIjoidGhpcyBpcyBpc3N1ZXIiLCJhdWQiOiJ0aGlzIGlzIGF1ZGllbmNlIiwic3ViIjoiMTIzNDU2Nzg5MCIsInVzZXIiOiJKb2huIERvZSIsImVtYWlsIjoiam9obmRvZUBleGFtcGxlLmNvbSIsInJvbGUiOiJVc2VyIiwiaWF0IjoxNzI5MjQwOTMxfQ.F9bf-5vQDuPvssys73JGtOgVPPngd2CUx_6L6v6QnEY"
	jwtToken := Unwrap(ParseJwtToken(token, "your-256-bit-secret"))
	jwtClaims := Unwrap(GetJwtClaimsFromJwtToken(jwtToken))
	jwtClaimsAccessData := CvtJwtClaimsToJwtClaimsAccessData(jwtClaims)

	//fmt.Println(JsonPreviewReflection(jwtToken))
	//fmt.Println(JsonPreviewReflection(jwtClaims))
	fmt.Println(JsonPreviewReflection(ShikaObjectConversionPreview(jwtClaimsAccessData)))
}
