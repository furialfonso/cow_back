package team

var (
	GetUsersByGroup = `
		select b.user_id
		from cow_local_db.c_group a,
					cow_local_db.c_team b
		where b.group_id = a.id
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
