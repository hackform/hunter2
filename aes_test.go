package hunter2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAESCipher_Encrypt(t *testing.T) {
	t.Parallel()

	assert := require.New(t)

	config, err := NewAESConfig()
	assert.NoError(err)
	assert.NotNil(config)
	key := config.String()
	aead, err := AESCipherFromParams(key)
	assert.NoError(err)
	assert.NotNil(aead)

	for _, tc := range []struct {
		Plaintext string
	}{
		{
			Plaintext: "hello, world",
		},
		{
			Plaintext: "Lorem ipsum",
		},
	} {
		tc := tc
		t.Run(tc.Plaintext, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)

			ciphertext, err := aead.Encrypt(tc.Plaintext)
			assert.NoError(err)
			plaintext, err := aead.Decrypt(ciphertext)
			assert.NoError(err)
			assert.Equal(tc.Plaintext, plaintext)
		})
	}
}
