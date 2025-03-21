package queries

var (
	GetTeamByBudget = `
		select b.user_id
		from c_budget a,
				 c_team b
		where b.budget_id = a.id
		and a.code = ?
	`

	GetTeamsByUser = `
		select b.id,
		       b.code,
           b.debt,
					 b.created_at 
		from c_team a,
     		 c_budget b
		where a.user_id = ?
	`

	GetUserByID = `
		select if (count(*) = 0,0,1) exist from c_team where user_id = ?
	`

	ComposeTeam = `
		INSERT INTO c_team(budget_id,
											 user_id)VALUES(?,?)
	`

	DecomposeTeam = `
		DELETE FROM c_team where budget_id = ? and user_id = ?
	`
)
