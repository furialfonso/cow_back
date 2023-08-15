package templates

var (
	GetGroups = `
			select id,
						code,
						debt,
						created_at
			from c_group;
	`

	GetGroupByCode = `
			select id,
						code,
						debt,
						created_at
			from c_group
			where code = ?;
	`

	CreateGroup = `
		insert into c_group(code)values(?)
	`

	DeleteGroup = `
		delete from c_group
		where code = ?;
	`

	UpdateGroup = `
		update c_group 
		set debt = ?
		where code = ?
	`
)
