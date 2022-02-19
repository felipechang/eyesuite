package suitetalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Settings interface {
	GetVersion() string
	GetWsUrl() string
	GetRealm() string
	GetConsumerKey() string
	GetTokenID() string
	GetNonce() string
	GetTimestamp() string
	GetSignature(nonce string, timeStamp string) string
}

type settings struct {
	*storage.Config
}

func NewSettings(config *storage.Config) Settings {
	t := &settings{config}
	return t
}

func (s *settings) GetVersion() string {
	return s.Version
}

func (s *settings) GetWsUrl() string {
	return s.WsUrl
}

func (s *settings) GetRealm() string {
	return s.Realm
}

func (s *settings) GetConsumerKey() string {
	return s.ConsumerKey
}

func (s *settings) GetTokenID() string {
	return s.TokenID
}

// GetNonce generate 32 character nonce
func (s *settings) GetNonce() string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

// GetTimestamp generate a unix timestamp
func (s *settings) GetTimestamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

// GetSignature calculate signature from base and signing Key
func (s *settings) GetSignature(nonce string, timeStamp string) string {
	base := strings.Join([]string{
		s.Realm,
		s.ConsumerKey,
		s.TokenID,
		nonce,
		timeStamp,
	}, "&")
	key := strings.Join([]string{
		s.ConsumerSecret,
		s.TokenSecret,
	}, "&")
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(base))
	signature := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}
