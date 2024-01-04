package template

var (
	GetAll = `
			select id,
						code,
						debt,
						created_at
			from c_group;
	`

	GetByCode = `
			select id,
						code,
						debt,
						created_at
			from c_group
			where code = ?;
	`

	Create = `
		insert into c_group(code)values(?)
	`

	Delete = `
		delete from c_group
		where code = ?;
	`

	Update = `
		update c_group 
		set debt = ?
		where code = ?
	`
)
