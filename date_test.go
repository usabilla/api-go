/*
Copyright (c) 2015 Usabilla

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish, dis-
tribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the fol-
lowing conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABIL-
ITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT
SHALL THE AUTHOR BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
*/

package gobilla

import (
	"testing"
	"time"

	"github.com/usabilla/gobilla/internal"
)

var (
	date = time.Date(2015, time.February, 10, 23, 0, 0, 0, time.UTC)
)

func Test_GetRFC1123GMT(t *testing.T) {
	spec := internal.Spec(t)
	rfcDate := getRFC1123GMT(date)
	expected := "Tue, 10 Feb 2015 23:00:00 GMT"
	spec.Expect(rfcDate).ToEqual(expected)
}

func Test_GetShortDate(t *testing.T) {
	spec := internal.Spec(t)
	shortDate := getShortDate(date)
	expected := "20150210"
	spec.Expect(shortDate).ToEqual(expected)
}

func Test_GetShortDateTime(t *testing.T) {
	spec := internal.Spec(t)
	shortDateTime := getShortDateTime(date)
	expected := "20150210T230000Z"
	spec.Expect(shortDateTime).ToEqual(expected)
}
