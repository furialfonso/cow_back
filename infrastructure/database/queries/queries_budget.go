package queries

var (
	GetAll = `
			select id,
						code,
						debt,
						created_at
			from c_budget;
	`

	GetByCode = `
			select id,
						code,
						debt,
						created_at
			from c_budget
			where code = ?;
	`

	Create = `
		insert into c_budget(code)values(?)
	`

	Delete = `
		delete from c_budget
		where code = ?;
	`

	Update = `
		update c_budget 
		set debt = ?
		where code = ?
	`
)
