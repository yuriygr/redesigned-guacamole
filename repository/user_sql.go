package repository

const selectUser = `
	select
		s.*,
		get_subsite_subscribers(s.subsite_id, 12) as subscribers,
		get_subsite_subscriptions(s.subsite_id, 6) as subscriptions
	from subsite as s
	where s.status = 1 and s.is_user = true and s.subsite_id = :r.subsite_id
`
