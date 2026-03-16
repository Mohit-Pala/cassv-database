package models

type TableMetadata struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	TableCols []TableCols `json:"table_cols"`
	LastID    int         `json:"last_id"`
	RowCount  int         `json:"row_count"`
}
