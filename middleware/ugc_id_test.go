package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUgcIdPathStore_GetUgcId(t *testing.T) {
	ugcId, err := NewUgcIdPathGetting().GetUgcId(&gin.Context{})
	assert.NoError(t, err)
	assert.Equal(t, ugcId, "123465xxx202401311817")

	err = NewUgcIdContextSetting().SetUgcId(&gin.Context{}, "47691")
	assert.NoError(t, err)

	ugcId, err = NewUgcIdPathGetting().GetUgcId(&gin.Context{})
	assert.NoError(t, err)
	assert.Equal(t, ugcId, "47691")
}
