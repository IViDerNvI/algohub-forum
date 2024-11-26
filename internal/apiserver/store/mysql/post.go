package mysql

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
	"gorm.io/gorm"
)

type postStore struct {
	db *gorm.DB
}

func (ps *postStore) Create(ctx context.Context, post model.Post, opt model.CreateOptions) error {
	return ps.db.Create(&post).Error
}

func (ps *postStore) Update(ctx context.Context, post model.Post, opt model.UpdateOptions) error {
	return ps.db.Save(post).Error
}

func (ps *postStore) GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Post, error) {
	ret := &model.Post{}
	if err := ps.db.Where("id = ?", id).First(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (ps *postStore) DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error {
	return ps.db.Where("id = ?", id).Delete(&model.Post{}).Error
}

func (ps *postStore) List(ctx context.Context, title string, opt model.ListOptions) (*model.PostList, error) {
	ret := &model.PostList{}
	d := ps.db.Where("title like ?", title).
		Offset(int(*opt.Offset)).
		Limit(int(*opt.Limit)).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)
	return ret, d.Error
}

func newPosts(ds *datastore) *postStore {
	return &postStore{ds.db}
}
