package mysql

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
	"gorm.io/gorm"
)

type likeStore struct {
	db *gorm.DB
}

func (ls *likeStore) Create(ctx context.Context, like model.Like, opt model.CreateOptions) error {
	return ls.db.Create(&like).Error
}

func (ls *likeStore) Update(ctx context.Context, like model.Like, opt model.UpdateOptions) error {
	return ls.db.Save(like).Error
}

func (ls *likeStore) GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Like, error) {
	ret := &model.Like{}
	if err := ls.db.Where("userid = ?", id).First(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func (ls *likeStore) DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error {
	return ls.db.Where("userid = ?", id).Delete(&model.Like{}).Error
}

func (ls *likeStore) List(ctx context.Context, title string, opt model.ListOptions) (*model.LikeList, error) {
	ret := &model.LikeList{}
	d := ls.db.Where("title like ?", title).
		Offset(int(opt.Offset)).
		Limit(int(opt.Limit)).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)
	return ret, d.Error
}

func newLikes(ds *datastore) *likeStore {
	return &likeStore{ds.db}
}
