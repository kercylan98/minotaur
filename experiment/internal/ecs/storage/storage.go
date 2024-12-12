package storage

type Storage[PK, Col comparable] interface {
	Migrate(target Storage[PK, Col], primaryKes ...PK)

	SetColumn(col Col, defaultGetter func() any)

	Get(primaryKey PK, col Col) any

	GetRow(primaryKey PK) map[Col]any

	AddRow(primaryKey PK)

	AddRows(primaryKeys []PK)

	AddRowsWithValues(primaryKeys []PK, values map[Col][]any)

	DelRow(primaryKey PK)

	DelRows(primaryKeys []PK)
}
