package team

var (
	GetUsersByGroup = `
		select c.nick_name 
		from cow_local_db.c_group a,
					cow_local_db.c_team b,
					cow_local_db.c_user c
		where b.user_id = c.id  
		and b.group_id = a.id
		and a.code = ?
	`

	GetUserByID = `
		select if (count(*) = 0,0,1) exist from c_team where user_id = ?
	`

	ComposeTeam = `
		INSERT INTO c_team(group_id,
											 user_id)VALUES(?,?)
	`

	DecomposeTeam = `
		DELETE FROM c_team where group_id = ? and user_id = ?
	`
)
