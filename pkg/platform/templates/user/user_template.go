package user

var (
	GetAll = `
		select id,
					 name,
					 second_name,
					 last_name,
					 second_last_name,
					 email,
					 nick_name,
					 created_at
		from c_user
	`

	GetByCode = `
		select id,
					 name,
					 second_name,
					 last_name,
					 second_last_name,
					 email,
					 nick_name,
					 created_at
		from c_user
		where nick_name = ?;
	`

	Create = `
		insert into c_user(name,
											 second_name,
											 last_name,
											 second_last_name,
											 email,
											 nick_name)values(?,?,?,?,?,?)
	`

	Delete = `
		delete from c_user
		where nick_name = ?;
	`
)
