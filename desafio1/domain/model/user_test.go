package model

import (
	"github.com/codeedu/imersao/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

package model_test

import (
"testing"

uuid "github.com/satori/go.uuid"

"github.com/codeedu/imersao/codepix-go/domain/model"
"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {

	id := "ce84c924-829e-4997-bd1f-094568e5b503"
	name := "Thaiane"
	email:="thaiane.tas@gmail.com"

	user, err := model.NewUser(id, name,email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}
