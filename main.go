package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/iotaledger/iota.go/guards/validators"
	"github.com/iotaledger/iota.go/v2/bech32"
)

const recoveryPubkey = `-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBFn7N9UBEADAel8VNth0zbhVobJ21zhPiqMCQuZdc3zc8ojc5ZBGWTMmrcMK
vK1Jt+ganZCLdusRLtWTCkvlpwhtUI24cSB07auGhyJnpG7uA4CaEcSrXdLhZkKi
CRUTXKUtxNLNzPcF7E5bxt5pyK3zlMMZT/y2XenfV4q9/WktnCNnPwV14/0AxxI1
4WIXU+0SGajQSnf8DmDkpJOYad7g5Cd1NWEGcnsY2vnbsdf2qN23wkLaA6RFS3Jj
PCLdTspQyxkoX3fhAYYOqH2mt8PhEtnrlXc1adGaCH7megX+qRPqUM9TZQIu1l2A
qa4oOvuqEsbyhwG0Fdx3MtX1zyCPJGbl6FOWLUBqsf85ROdxjWaFuPoq0iM3rU5J
cb7B5GpmhCWFcN3axBv5JsHvmBgEK3MSapGZC3yF4wp8iF1GF8nmCMWPUQruxsm+
dCsAoCR3RAxbsM8krr1sU5gVvwim9euXkd9gzK7mGXfdbQ1p1tdwSidHoJoZ7Zr6
ihsoXMP2FH046hHuIz+Cxu3U8tKPmgLyh/wIE1izTIKsMVavwu+sRbwYEiLrjC/x
W7VRaj7uvjFgche5EnSB2UwtkQDyP94BvDILOb4YU0VmBJv26mKHn4PS9v9CDfZ3
WsCssYE5GaPOyQG+4fMNwNAs1mluv/UWRTR7J1RBUAHtPH8DqHeDUhKlIwARAQAB
tEZJT1RBIEZvdW5kYXRpb24gUmVjb3ZlcnkgVG9vbCAobm8taHVtYW5zLWludm9s
dmVkKSA8cmVjb3ZlcnlAaW90YS5vcmc+iQJUBBMBCAA+AhsNBQsJCAcCBhUICQoL
AgQWAgMBAh4BAheAFiEECmyt5VeVj8atvnAQ8bCzdIXm/xMFAmCP3WYFCQpXDJEA
CgkQ8bCzdIXm/xNB3Q//SH63wyKWMlW93gKxY50ues1ea0O7ZtKKpDqUCvH0gim6
mB5tm3P81r35PWot4I04+Sd99MtfLIfuVOpDSm7FONRgmj2m7Dh4GX8AnvnZZVl0
U6bQCfjRlEYK7Wsyf+zOrRkYMR36Ps5u6MVSJJvKmRKvXo4x1EKXITLJSseum1il
cx5dfm+jSXANyYggFzpFb4r3bIRCR/wnh3H/J31nKtF48ij6PPhhH0pPS7MPJ25Z
j1KnzcLkS3KkVFevdekM8tSlvI9Zt66DoVtwHmNRH5/Ysiy8sPKtYioRRY+fq5nP
FID+RFjU2qe0amT1VL1N4QP8/aHKnqd84HovtPIa4VPfV8GvVpr4JXV6oDu40Ylm
cmlCWKAUi0HhB+L2O+TRbhvhJgKdP3+T7pcr9yuRzGUdCaLMAB+uLaEa4fpcmbJx
ZjFThs8QSuVXLiFQ1O3L9wDBqLw2LjxLwRq8k/bK//fT/aMu365kpRc3tJxG+53U
QYEB6Mep9LKJsUh4YrMV7s258JUlceslXAzOSkOE1wBv0PGfJtizLcoMAuKmdUod
vOS1ECMwA44JlA7ABr3MliYiJQHYGck9DAko+lQYZYz5KPuf3qp9OsgiExpgp2U6
tP48EO+zWRLns8fGBz4nIC76Ffs9iQNp9EfZnMMzBgfZXPLrb6MgWUuX3nIkvci5
Ag0EWkYQwwEQAKeRpk4zpWzB2d6ps9TWN15jFZG2LjjI2zz415Hf5DlPhLcPY7eb
NSFMD4PrQmIJTS6GQIewQCwS3X0E+NXF7zcjx16utZdpyv6fC+E6EzLqsjP8Xgbb
cFNEST4/+bd1SDaFmCmpq9MqGgrw/I86bIaaDNLTHPcyqm4Zjdfss1vwbnQbkJFI
PIbhCJea+2Htwi1Q8a/5Rlrm3uqxTjp5ov2W/CtF0JLWPQdXRNZYWN3+WeYXZC1x
Im59O+YuPjitMBrRGxd5XTfQTWmE58FejCIR2nESfSryJadrkw4pbc0UlhydMs2r
t1HKcqg/AnTXZJenh7e5YW/pKWBCUOb1xVf8gtgbMGy8lnSXdRRX16pu6PP70mOj
xE1v2fsfwr/bXeIJtMkZjj/c0qECPcrR0UsR/pbzjJqSD/OqRcfgc6ckDpqsJoFq
IuKaTdSu3eEYuua+KJtzmpB1XXQoAhUkn9Txv8pPI1DKsNykBtLy15w3QoaC8dmy
9qfJ20klgUoUapyNHinHty6VZ+tmnXYRwp5fWWnlhJbOgCLxn5bD8iBePwz8fzp3
1W0n3Y9W+zI5JGdPTtlAAI8p91IEcMukMzDbIvR6PWZ2415Pj4rjAuGq8yjn/max
pMvBFtKk8DCr5ZJ+eJXb5H0ZL3BtAK0WB7cMjuCiaflGnGOT61Rh3RAtABEBAAGJ
BEQEGAECAA8FAlpGEMMCGwIFCQHhM4ACKQkQ8bCzdIXm/xPBXSAEGQECAAYFAlpG
EMMACgkQpMnMZm1Y/FRsiRAAlnw7kM4HiUyDO4bCQFA1DjUrlJAF7wn97VBo5WOQ
ZBROYUSw8G7hoWzwONLrbKnE6hc18Yind6K46GOO8a9xxrtYmfXOxeGU43pSk07e
EXL2VniMAMWVLPzU8FMbnfCfxI78qriZl9b7lsctDellO1kVmKqcax7E259sq3uY
M8CZoCqWPGN3uIT+SgPzuG9WPoefBRqWxpKMfZD/r40riwIvgC7grqDIuvghhFDA
F2+LKjv3W/aB4cb3HRbYPzBcuNXaZFmqw7f5erAyCfgRmF2Ybq894313Kz6khFgV
FUrob0WUwmYbhRldEEymRZynq90hbk+Bk/3k6JAhFGAuWti4G7ZclrGBi/oSFcmM
Q5DkWUMfv8c0bzax6gqRWxfV62g5EpkTkXTtFJB3UE9cdmGuMfwFQ+sYL3Vhhzjz
mPMRvfcXv9e0qAyZbfXcrK0TywYifr/K9XNaLmmKX5/Tqh0dxEGS4SDByF5XZiM7
ESSRgDk9fAV+WqTitZysrbYKBojRU7GduaeeTkfsySzv6ekKL3JjS5vRiu8cTHhe
2aG2Z1h5RqtAkPk1yeA2KQlTQsHer8yUAJOdvWUP4BzmPmLSY1ZGgYxa4IpSaNVo
TdgGcCzl+BiWHp/fuvX5cry0UNGgBJr/D5WaT4IxVC/WJAKQ/ol8brx4XCv6wSWp
VFPGlg//WPHhQmrcveL+xVI/HVeJDeTFsUYZV9Jn/KIcEt8KeApROrL9fgLDvM00
dH481tsc3RS4vpMRC6xF0HFRbB59WCgLL1CiCGNXwQNfU/d8IqjHamKawrfX1sdh
fney2ydo1Xw6Ed5sOrxCuoUPE2Gcml7EUefBTLZZi7mbeNgxrl6jl2WNPY5ViX3x
ZcrxOgEStbJS+xc2REHLaEFGO+llQOzDO67j0lOYkXD0FMV/Sdz5qvA/KTG9DyFx
H7T8EKYHtWpfAA48PvdcQsHBwCT0S9mHQM7f162ySjSXuDK/L4m2KWjpEHen6mTe
PnYZ+HcCd2tYxNDMPPXau+o9QYuP1g2t+6hZMtx+eNiWQfBjeSKkVVI8RDAdFT1x
iNRv8OnP0hWzL8qjh1JQwTceyzY2qpa4c+yAZN8Ni1j1jTyJXwmIVdsnUIRyBnJJ
cgv79NkkZIZDrLiChvgovu9WIKhG3i7/AnUdEuVLVvfs+mQDAkA5OglP4VPvCosw
ttIFo3ODQuzGKPX9mSdVWiVw5raBgfIv+agg8R1oVIQSL3KHtsMYjiRTZsp3i5bY
y4fP9qL58OAhKiZGpYvHaG482pE07VqpHtJugK0QKS6giSTASTBmVOSZgt3obz7J
jCUA95iCK3iwlNEuYRsWFFt7tHYoufIbsW++c5X2TWCYsYRtyjA=
=Fuz2
-----END PGP PUBLIC KEY BLOCK-----
`

func main() {
	reader := bufio.NewReader(os.Stdin)
	address, payout, seed, err := input(reader)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		message := fmt.Sprintf("Address: %s\nPayout: %s\nSeed: %s\n", address, payout, seed)
		if armor, err := helper.EncryptMessageArmored(recoveryPubkey, message); err != nil {
			fmt.Printf("ERROR: failed to encrypt message: %s\n", err)
		} else {
			fmt.Printf("\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##\n##")
			fmt.Println("\n\n=================================CUT HERE=====================================\n")
			fmt.Println("Send __only__ the following text to our community managers (via Discord or Email)")
			fmt.Printf("\n\n\n%s\n\n", armor)
			fmt.Println("\n\n==============================================================================\n")
		}
	}
	fmt.Println("Press ENTER to exit...")
	reader.ReadString('\n')
}

func input(reader *bufio.Reader) (string, string, string, error) {

	fmt.Print("Enter Claimed address including checksum: ")
	address, err := reader.ReadString('\n')
	address = strings.Replace(address, "\n", "", -1)
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Enter Bech32 Payout address (including `iota` prefix): ")
	payout, err := reader.ReadString('\n')
	payout = strings.Replace(payout, "\n", "", -1)
	if err != nil {
		return "", "", "", err
	}

	fmt.Print("Enter Seed for address: ")
	seed, err := reader.ReadString('\n')
	seed = strings.Replace(seed, "\n", "", -1)
	if err != nil {
		return "", "", "", err
	}

	if err := validators.ValidateAddresses(true, address)(); err != nil {
		err := fmt.Errorf("The entered Claimed address is invalid:\n%w", err)
		return "", "", "", err
	}

	if _, _, err := bech32.Decode(payout); err != nil {
		err := fmt.Errorf("The entered Bech32 Payout address is invalid:\n%w", err)
		return "", "", "", err
	}

	if err := validators.ValidateSeed(seed)(); err != nil {
		err := fmt.Errorf("The entered Seed is invalid:\n%w", err)
		return "", "", "", err
	}

	return address, payout, seed, nil

}
