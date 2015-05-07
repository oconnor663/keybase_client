package libkb

import (
	"testing"
)

// See Issue #40: https://github.com/keybase/client/issues/40
func TestPGPGetPrimaryUID(t *testing.T) {
	armored := `
-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: GPGTools - https://gpgtools.org

mI0EVUubbgEEAKvIPrcQksiZW/dIX3sfZCSxfZqS8nojN/307rZ7323oAPTG6W7Y
EHfdgpKxUfMrdiNiWtujbybT8XUFhZwuPZF06BMIVmXESBr7rHLZ1Up41HmjvTtR
n1ejyy/tMl9zjydTIxjXsML9mPe3k3KVlrfgrbvadeM6PdYw+oE8tLZRABEBAAG0
IVByaW1hcnkgUHJpbWFyeSA8cHJpbWFyeUB1aWQuY29tPoi8BBMBCgAmAhsDBwsJ
CAcDAgEGFQgCCQoLBBYCAwECHgECF4AFAlVLm6YCGQEACgkQ5+9u0rO5NRk/4wP/
WBY3/C1keVxQMdma/F6l+yeP8SQvzqTLWZKRu1tPyi2DWYu8HnVw5/Upazj/UC8u
4Q7OzBsU0dn+5TUrwSfuTATY7OOSgWo/Nf++rqTU2z4h8kmV/qelGTrBKvshAEMk
/9ejJhPcErXocHHkAGCB+V4lSC70lBsAYoriEg2jmvW0J1NlY29uZGFyeSBTZWNv
bmRhcnkgPHNlY29uZGFyeUB1aWQuY29tPoi5BBMBCgAjBQJVS5uBAhsDBwsJCAcD
AgEGFQgCCQoLBBYCAwECHgECF4AACgkQ5+9u0rO5NRmZzwP/UnO2OBcaT3doVDbz
o246Ur3M/USz9XVhBMabCpd4iANXswi/9iJyFW0rMn3DXAUO3Qj8yh/077FfMqUj
srcmbC4qtscd3X1V7hQ4rOsiEghUu2ZG8XsFxBh01NpDYxrTXQfmJeegUMAirgrA
r8WBLqMoZo8TarCDCw/6ygTxV6y0JFRlcnRpYXJ5IFRlcnRpYXJ5IDx0ZXJ0aWFy
eUB1aWQuY29tPoi5BBMBCgAjBQJVS5uVAhsDBwsJCAcDAgEGFQgCCQoLBBYCAwEC
HgECF4AACgkQ5+9u0rO5NRlZegP8Dl2IGu3WwF4w2Qxj/WzhTeiaLQoTzEsF0IPy
+f+IBO/uC/5bw/b3uNAhGK1wK3hy77II+py2x7/EvJAz2w1Ua9IMD6YuZBQaSzwV
tRhoHOnptdxLNDZPvZVuPl8G6p4yKLqelGymtdtpkObz9w8f+KTcieEUVeM7HbI8
/dwF7da4jQRVS5tuAQQAy0uPwJBoQeXz/uv5ifdXJG39cYFTaONQxa6U3/Wtx3oV
ibdqyygmFliHJJJyx0wjMohVbDT2sl8bHr0gpZxnkDF7+Gmcwcn/ohfKf0h7hrmL
W4nGNdiQiD6QZZz7ebEBuqzUmkAYg33lsYZ8MLI6wLK3sMmzA468di06YUa9eksA
EQEAAYifBBgBCgAJBQJVS5tuAhsMAAoJEOfvbtKzuTUZMFMD/28lUOflsdcLvJyL
L22yGWeoUlKMQmdbreSUo3Ibcdkpsy9ZHDHxWRBN5s/cxHtX6nna5IeHSNIdDruX
HKfiyXs8709e067vsE5FCTMvZCq4vt/lkEJ59xn58QBfEILMwQDNLqVGyA54MPwh
+NwHX7807fKQwvqyVixJoZ3yS6Js
=RH/W
-----END PGP PUBLIC KEY BLOCK-----
`
	expected := "Primary Primary <primary@uid.com>"
	for i := 0; i < 100; i++ {
		key, err := ReadOneKeyFromString(armored)
		if err != nil {
			t.Error(err)
		}
		primary := key.GetPrimaryUID()
		if primary != expected {
			t.Errorf("Expected '%s' as a primary UID; got '%s'", expected, primary)
		}
	}
	return
}
