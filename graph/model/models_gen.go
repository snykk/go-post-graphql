// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Post   *Post  `json:"post"`
	Author *User  `json:"author"`
}

type Mutation struct {
}

type NewCommentInput struct {
	Text     string `json:"text"`
	PostID   string `json:"postID"`
	AuthorID string `json:"authorID"`
}

type NewPostInput struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"authorID"`
}

type NewUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Post struct {
	ID       string     `json:"id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	Author   *User      `json:"author"`
	Comments []*Comment `json:"comments,omitempty"`
}

type Query struct {
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Posts []*Post `json:"posts,omitempty"`
}
