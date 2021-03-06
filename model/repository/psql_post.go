package repository

import (
	"database/sql"
	"errors"

	"NoticeBoard/entity"
)

type PostRepositoryImpl struct {
	conn *sql.DB
}

func NewPostRepositoryImpl(Conn *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{conn: Conn}
}

func (pri *PostRepositoryImpl) Posts() ([]entity.Post, error) {
	
	rows, err := pri.conn.Query("SELECT * FROM posts;")
	if err != nil {
		return nil, errors.New("could not query the database")
	}
	defer rows.Close()

	posts := []entity.Post{}

	for rows.Next() {
		post := entity.Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Description, &post.Image, &post.Type)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	
	return posts, nil

}

func (pri *PostRepositoryImpl) Post(id int) (entity.Post, error) {
	
	row := pri.conn.QueryRow("SELECT * FROM posts WHERE id = $1", id)

	post := entity.Post{}

	err := row.Scan(&post.Id, &post.Title, &post.Description, &post.Image, &post.Type)
	if err != nil {
		return post, err
	}

	return post, nil

}

func (pri *PostRepositoryImpl) UpdatePost(post entity.Post) error {
	
	_, err := pri.conn.Exec("UPDATE posts SET title=$1, description=$2, image=$3, type=$4 WHERE id=$5", post.Title, post.Description, post.Image, post.Type, post.Id)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

func (pri *PostRepositoryImpl) DeletePost(id int) error {
	
	_, err := pri.conn.Exec("DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (pri *PostRepositoryImpl) StorePost(post entity.Post) error {
	
	_, err := pri.conn.Exec("INSERT INTO posts (title,description,image,type) values($1, $2, $3, $4)", post.Title, post.Description, post.Image, post.Type)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}