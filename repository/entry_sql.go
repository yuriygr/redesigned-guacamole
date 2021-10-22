package repository

const selectTimelineEntries = `
	select
		e.*,
		ec.*,
		get_entry_subsite(e.subsite_id, :r.current_user_id) as subsite,
		get_entry_author(e.user_id, :r.current_user_id) as author,
		get_entry_counters(e.entry_id) as counters,
		get_entry_vote(e.entry_id, :r.current_user_id) as vote
	from entry as e
	left join entry_content as ec on ec.entry_id = e.entry_id
	left join subsite as es on es.subsite_id = e.subsite_id
	left join subsite as eu on eu.subsite_id = e.user_id
	inner join subsite_subscription as ss on ss.to = e.subsite_id
	where e.is_draft = 0 and ss.from = :r.current_user_id or e.subsite_id = :r.current_user_id
	group by e.entry_id, ec.content_id
`

const selectSubsiteEntries = `
	select
		e.*,
		ec.*,
		get_entry_subsite(e.subsite_id, :r.current_user_id) as subsite,
		get_entry_author(e.user_id, :r.current_user_id) as author,
		get_entry_counters(e.entry_id) as counters,
		get_entry_vote(e.entry_id, :r.current_user_id) as vote
	from entry as e
	left join entry_content as ec on ec.entry_id = e.entry_id
	left join subsite as es on es.subsite_id = e.subsite_id
	left join subsite as eu on eu.subsite_id = e.user_id
	where e.is_draft = 0 and (eu.subsite_id = :r.subsite_id or es.subsite_id = :r.subsite_id)
`

const selectEntry = `
	select
		e.*,
		ec.*,
		get_entry_subsite(e.subsite_id, :r.current_user_id) as subsite,
		get_entry_author(e.user_id, :r.current_user_id) as author,
		get_entry_counters(e.entry_id) as counters,
		get_entry_vote(e.entry_id, :r.current_user_id) as vote
	from entry as e
	left join entry_content as ec on ec.entry_id = e.entry_id
	left join subsite as es on es.subsite_id = e.subsite_id
	left join subsite as eu on eu.subsite_id = e.user_id
	where e.is_draft = 0 and e.entry_id = :r.entry_id
`
