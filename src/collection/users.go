package collection

type User struct {
	ID uint64 `db:"id,omitempty" json:"-"`
  Name string `db:"name" json:"name"`
}

