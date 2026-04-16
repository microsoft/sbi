//     MIT License
//
//     Copyright (c) Microsoft Corporation.
//
//     Permission is hereby granted, free of charge, to any person obtaining a copy
//     of this software and associated documentation files (the "Software"), to deal
//     in the Software without restriction, including without limitation the rights
//     to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//     copies of the Software, and to permit persons to whom the Software is
//     furnished to do so, subject to the following conditions:
//
//     The above copyright notice and this permission notice shall be included in all
//     copies or substantial portions of the Software.
//
//     THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//     IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//     FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//     AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//     LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//     OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//     SOFTWARE

package report

import "testing"

func FuzzHumanSize(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(1024))
	f.Add(int64(1048576))
	f.Add(int64(1073741824))
	f.Add(int64(-1))
	f.Add(int64(9999999999999))

	f.Fuzz(func(t *testing.T, numBytes int64) {
		result := HumanSize(numBytes)
		if result == "" {
			t.Error("HumanSize returned empty string")
		}
	})
}

func FuzzFormatDigest(f *testing.F) {
	f.Add("sha256:abcdef1234567890abcdef1234567890")
	f.Add("")
	f.Add("sha256:short")
	f.Add("not-a-digest")
	f.Add("sha256:")

	f.Fuzz(func(t *testing.T, digest string) {
		// FormatDigest must not panic on any input
		_ = FormatDigest(digest)
	})
}
