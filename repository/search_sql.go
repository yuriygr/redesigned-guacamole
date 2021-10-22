package repository

const selectFastSearch = `
	select s.name, s.avatar from subsite as s
	where s.status = 1
`
