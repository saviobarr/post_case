package service

import (
	"github.com/saviobarr/post_case/domain"
	"github.com/saviobarr/post_case/utils"
)

//GetPost gets a post by its id
func GetPost(id int64) (*domain.Post, *utils.ApplicationError) {
	return nil, nil
}

//CreatePost creates a new post
func CreatePost(domain.Post) *utils.ApplicationError {
	return nil
}

//GetComments gets comments from a post by its id postId
func GetComments(postID int64) (*[]domain.Comment, *utils.ApplicationError) {
	return nil, nil

}

//CreateComment creates a new comment
func CreateComment(domain.Comment) *utils.ApplicationError {
	return nil
}
