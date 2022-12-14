package post

import (
	"errors"

	"github.com/gabrielgaspar447/go-blog-api/constants"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gabrielgaspar447/go-blog-api/repositories"
)

func createPostService(input *models.Post) error {
	return repositories.PostCreate(input)
}

func listPostsService(posts *[]models.Post, includeUser bool) error {
	err := repositories.PostList(posts, includeUser)
	if err != nil {
		return err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.ID = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return nil
}

func getPostByIdService(post *models.Post, id uint, includeUser bool) error {
	err := repositories.PostGetById(post, id, includeUser)
	if err != nil {
		return err
	}

	if post.ID == 0 {
		return errors.New(constants.PostNotFound)
	}

	if includeUser {
		post.User.Password = ""
		post.User.ID = 0
		post.User.CreatedAt = nil
		post.User.UpdatedAt = nil
	}

	return nil
}

func searchPostsService(posts *[]models.Post, query string, includeUser bool) error {
	err := repositories.PostSearch(posts, query, includeUser)
	if err != nil {
		return err
	}

	if includeUser {
		for i := range *posts {
			(*posts)[i].User.Password = ""
			(*posts)[i].User.ID = 0
			(*posts)[i].User.CreatedAt = nil
			(*posts)[i].User.UpdatedAt = nil
		}
	}

	return nil
}

func updatePostService(input *models.Post, userId uint) error {
	post := &models.Post{}
	err := repositories.PostGetById(post, input.ID, false)
	if err != nil {
		return err
	}

	if post.ID == 0 {
		return errors.New(constants.PostNotFound)
	} else if post.UserID != userId {
		return errors.New(constants.PostNotOwned)
	}

	return repositories.PostUpdate(input, input.ID)
}

func deletePostService(id uint, userId uint) error {
	post := &models.Post{}
	err := repositories.PostGetById(post, id, false)
	if err != nil {
		return err
	}

	if post.ID == 0 {
		return errors.New(constants.PostNotFound)
	} else if post.UserID != userId {
		return errors.New(constants.PostNotOwned)
	}

	return repositories.PostDelete(id)
}
