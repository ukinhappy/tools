package command

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
)
import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"hash"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// oa represents the time command
var oaCmd = &cobra.Command{
	Use:   "oa [OPTIONS]",
	Short: "oa command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			var flowresult XmlItems
			auth := NewGoogleAuth("", SecretOption(args[0]))
			code, _ := auth.GetCode()
			flowresult.Iterm = append(flowresult.Iterm,
				Item{
					Title:    code,
					Subtitle: "加密后的结果:",
					Arg:      code,
					Valid:    "yes"})

			b, _ := xml.Marshal(flowresult)
			fmt.Println(string(b))
		}
	},
}

func init() {
	rootCmd.AddCommand(oaCmd)
}


var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLen = len(codes)
)

type (
	GoogleAuth struct {
		user, issuer, secret string
	}
	Option func(o *GoogleAuth)
)

func SecretOption(s string) Option {
	return func(opts *GoogleAuth) {
		opts.secret = s
	}
}

func IssuerOption(issuer string) Option {
	return func(opts *GoogleAuth) {
		opts.issuer = issuer
	}
}

func getCode(key []byte, value []byte) uint32 {
	var (
		hmacSha1         hash.Hash
		bytes, hashParts []byte
		offset           uint8
		number, pwd      uint32
	)
	hmacSha1 = hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	bytes = hmacSha1.Sum(nil)
	offset = bytes[len(bytes)-1] & 0x0F
	hashParts = bytes[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number = toUint32(hashParts)
	pwd = number % 1000000
	return pwd
}

func toBytes(value int64) []byte {
	var (
		result []byte
		mask   int64
		shifts [8]uint16
	)
	mask = int64(0xFF)
	shifts = [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

func (g *GoogleAuth) GetCode() (code string, err error) {
	var (
		key      []byte
		secret   string
		codeUI32 uint32
	)
	secret = strings.ToUpper(strings.Replace(g.secret, " ", "", -1))
	key, err = base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", nil
	}
	codeUI32 = getCode(key, toBytes(time.Now().Unix()/30))
	code = fmt.Sprintf("%0*d", 6, codeUI32)
	return
}

// ProvisionURI 生成url
func (g *GoogleAuth) ProvisionURI() (codeUrl string) {
	auth := "totp/"
	q := make(url.Values)
	q.Add("secret", g.secret)
	if g.issuer != "" {
		q.Add("issuer", g.issuer)
		auth += g.issuer + ":"
	}
	return "otpauth://" + auth + g.user + "?" + q.Encode()
}

// Authenticate 生成验证code正确
func (g *GoogleAuth) Authenticate(code string) (ok bool, err error) {
	var (
		nowCode string
	)
	nowCode, err = g.GetCode()
	return nowCode == code, err
}

// GenerateKey 生成密钥
func (g *GoogleAuth) GenerateKey() string {
	data := make([]byte, 32)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 32; i++ {
		idx := rand.Intn(codeLen)
		data[i] = codes[idx]
	}
	g.secret = string(data)
	return string(data)
}

func NewGoogleAuth(user string, options ...Option) *GoogleAuth {
	g := &GoogleAuth{
		user: user,
	}
	for _, op := range options {
		op(g)
	}
	return g
}
