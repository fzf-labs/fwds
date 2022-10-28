package bcrypt

import (
	"fwds/internal/conf"
	"github.com/speps/go-hashids"
)

type hash struct {
	secret string
	length int
}

func New(secret string, length int) *hash {
	return &hash{
		secret: secret,
		length: length,
	}
}

func NewDefault() *hash {
	return New(conf.Conf.HashIds.Secret, conf.Conf.HashIds.Length)
}

func (h *hash) HashidsEncode(params []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	hashID, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	hashStr, err := hashID.Encode(params)
	if err != nil {
		return "", err
	}

	return hashStr, nil
}

func (h *hash) HashidsDecode(hash string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	hashID, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	ids, err := hashID.DecodeWithError(hash)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
