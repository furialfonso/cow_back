package team

var (
	GetTeamByGroup = `
		select b.user_id
		from cow_local_db.c_group a,
					cow_local_db.c_team b
		where b.group_id = a.id
		and a.code = ?
	`

	GetTeamsByUser = `
		select b.id,
		       b.code,
           b.debt,
					 b.created_at 
		from cow_local_db.c_team a,
     		 cow_local_db.c_group b
		where a.user_id = ?
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
