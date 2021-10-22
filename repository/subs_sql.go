package repository

const selectSubsite = `
	select
		s.*,
		is_me(:r.current_user_id, s.subsite_id) as is_me,
		is_subscribed(:r.current_user_id, s.subsite_id) as is_subscribed,
		get_subsite_subscribers(s.subsite_id, 12) as subscribers,
		get_subsite_subscriptions(s.subsite_id, 6) as subscriptions,
		get_subsite_rules(s.subsite_id, 3) as rules
	from subsite as s
	where s.status = 1 and s.is_user = false and s.slug = :r.slug 
`

const existSubsite = `
	select exists(select * from subsite where subsite_id = ?)
`

const selectSubs = `
	select
		s.*,
		is_me(:r.current_user_id, s.subsite_id) as is_me,
		is_subscribed(:r.current_user_id, s.subsite_id) as is_subscribed
	from subsite as s
	where s.status = 1 and s.is_user = false
`

const selectSubsRecommendations = `
	select
		s.*,
		is_me(:r.current_user_id, s.subsite_id) as is_me,
		is_subscribed(:r.current_user_id, s.subsite_id) as is_subscribed
	from subsite as s
	left join recommendations as r on r.subsite_id = s.subsite_id
	where s.status = 1 and s.is_user = false and r.subsite_id is not null
	order by r.sort_order asc
`

const Subscribe = `
	insert into subsite_subscription (` + "`" + `from` + "`" + `, ` + "`" + `to` + "`" + `, created)
	values (:r.from, :r.to, now())
`

const Unsubscribe = `
	delete from subsite_subscription
	where ` + "`" + `from` + "`" + ` = :r.from and ` + "`" + `to` + "`" + ` = :r.to
`
