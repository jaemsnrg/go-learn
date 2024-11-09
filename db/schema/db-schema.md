
table accounts as A {
  id bigserial [pk]
  owner varchar [not null]
  balance bigint [not null]
  currency varchar [not null]
  created_at timestamptz [default: `now()`, not null]

  Indexes {
    owner
  }
}

table entries {
  id bigserial [pk]
  account_id bigint [ref: > A.id, not null]
  amount bigint [not null, note: 'can be negative or positive']
  created_at timestamptz [default: `now()`, not null]

  Indexes {
    account_id
  }
}

table transfers {
  id bigserial [pk]
  from_account_id bigint [ref: > A.id, not null]
  to_account_id bigint [ref: > A.id, not null]
  amount bigint [not null, note: 'must be positive']
  created_at timestamptz [default: `now()`, not null]

  Indexes {
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
}
