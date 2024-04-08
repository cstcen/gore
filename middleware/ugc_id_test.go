package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUgcIdPathStore_GetUgcId(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "https://api-sdev0.xk5.com/zelos/v1.0/game/224/ugc/123465xxx202401311817/products", nil)
	assert.NoError(t, err)

	ugcId, err := NewUgcIdPathStore().GetUgcId(request)
	assert.NoError(t, err)
	assert.Equal(t, ugcId, "123465xxx202401311817")

	err = NewUgcIdPathStore().SetUgcId(request, "47691")
	assert.NoError(t, err)

	ugcId, err = NewUgcIdPathStore().GetUgcId(request)
	assert.NoError(t, err)
	assert.Equal(t, ugcId, "47691")
}
