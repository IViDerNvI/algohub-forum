package mysql

import (
	"context"

	"github.com/ividernvi/algohub-forum/internal/model"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func (us *userStore) Create(ctx context.Context, user model.User, opt model.CreateOptions) error {
	return us.db.Create(&user).Error
}

func (us *userStore) Update(ctx context.Context, user model.User, opt model.UpdateOptions) error {
	return us.db.Save(user).Error
}

func (us *userStore) GetById(ctx context.Context, id int, opt model.GetOptions) (*model.User, error) {
	user := &model.User{}
	if err := us.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userStore) GetByAccount(ctx context.Context, account string, opt model.GetOptions) (*model.User, error) {
	user := &model.User{}
	if err := us.db.Where("username = ?", account).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userStore) DeleteById(ctx context.Context, id int, opt model.DeleteOptions) error {
	return us.db.Where("id = ?", id).Delete(&model.User{}).Error
}

func (us *userStore) DeleteByUsername(ctx context.Context, username string, opt model.DeleteOptions) error {
	return us.db.Where("username = ?", username).Delete(&model.User{}).Error
}

func (us *userStore) List(ctx context.Context, username string, opt model.ListOptions) (*model.UserList, error) {
	ret := &model.UserList{}

	// 先计算总数
	var totalCount int64
	d := us.db.Model(&model.User{}).Where("username like ?", "%"+opt.Pattern+"%").Count(&totalCount)
	if d.Error != nil {
		return nil, d.Error
	}
	ret.TotalCount = totalCount

	// 应用分页和排序
	d = us.db.Where("username like ?", "%"+opt.Pattern+"%").
		Offset(int(opt.Offset)).
		Limit(int(opt.Limit)).
		Order("id desc").
		Find(&ret.Users)

	if d.Error != nil {
		return nil, d.Error
	}

	return ret, nil
}

func newUsers(ds *datastore) *userStore {
	return &userStore{ds.db}
}

func DatabaseName() string {
	return "users"
}
