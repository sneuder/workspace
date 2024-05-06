package inputs

import (
	"workspace/models"
)

var collectionInputs = models.GenerateEmptyCollectionInputs()

func SetInputs() models.CollectionInputs {
	setWorkspaceInputs()
	return collectionInputs
}
