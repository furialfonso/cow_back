package templates

var (
	GetGroups = `
			select id,
						code,
						debt,
						created_at
			from cow_local_db.group;
	`

	GetGroupByCode = `
			select id,
						code,
						debt,
						created_at
			from cow_local_db.group
			where code = ?;
	`

	CreateGroup = `
		insert into cow_local_db.group(code)values(?)
	`

	UpdateGroup = `
		update cow_local_db.group 
		set debt = ?
		where code = ?
	`
)
