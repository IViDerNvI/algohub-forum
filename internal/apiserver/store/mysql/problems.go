package mysql

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
	"gorm.io/gorm"
)

type probStore struct {
	db *gorm.DB
}

func (ps *probStore) Create(ctx context.Context, prob *model.Problem, opt model.CreateOptions) error {
	return ps.db.Create(prob).Error
}

func (ps *probStore) Update(ctx context.Context, prob *model.Problem, opt model.UpdateOptions) error {
	return ps.db.Save(prob).Error
}

func (ps *probStore) GetById(ctx context.Context, id string, opt model.GetOptions) (*model.Problem, error) {
	prob := &model.Problem{}
	if err := ps.db.Where("id = ?", id).First(prob).Error; err != nil {
		return nil, err
	}
	return prob, nil
}

func (ps *probStore) DeleteById(ctx context.Context, id string, opt model.DeleteOptions) error {
	return ps.db.Where("id = ?", id).Delete(&model.Problem{}).Error
}

func (ps *probStore) List(ctx context.Context, opt model.ListOptions) (*model.ProblemList, error) {
	ret := &model.ProblemList{}

	// 先计算总数
	var totalCount int64
	d := ps.db.Model(&model.Problem{}).Where("title like ?", "%"+opt.Pattern+"%").Count(&totalCount)
	if d.Error != nil {
		return nil, d.Error
	}
	ret.TotalCount = totalCount

	// 应用分页和排序
	d = ps.db.Where("title like ?", "%"+opt.Pattern+"%").
		Offset(int(opt.Offset)).
		Limit(int(opt.Limit)).
		Order("id desc").
		Find(&ret.Items)

	if d.Error != nil {
		return nil, d.Error
	}

	return ret, nil
}

func newProblems(ds *datastore) *probStore {
	return &probStore{ds.db}
}
