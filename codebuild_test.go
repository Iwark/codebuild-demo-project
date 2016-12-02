package codebuild

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// コードのテスト
func TestSum(t *testing.T) {
	assert := assert.New(t)
	sum := Sum(1, 2)
	assert.Equal(3, sum)
}

// 環境変数のテスト
func TestEnv(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("codebuild", os.Getenv("MY_ENV"))
}
