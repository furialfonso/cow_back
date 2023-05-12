package templates

var (
	GetGroups = `
			select id,
						code,
						debt,
						created_at
			from cow_db.group;
	`

	GetGroupByCode = `
			select id,
						code,
						debt,
						created_at
			from cow_db.group
			where code = ?;
	`

	CreateGroup = `
		insert into cow_db.group(code)values(?)
	`

	UpdateGroup = `
		update cow_db.group 
		set debt = ?
		where code = ?
	`
)
