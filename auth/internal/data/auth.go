package data

import (
	"context"

	"github.com/Skijetler/GoBillingService/auth/internal/biz"
	"github.com/Skijetler/GoBillingService/pkg/ent"
	"github.com/Skijetler/GoBillingService/pkg/ent/user"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			r.log.WithContext(ctx).Errorf("Rolling back transaction error: %v", rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	var userModel *ent.User
	var userMetadata *ent.UserMetadata

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		var err error

		if userModel, err = tx.User.
			Create().
			SetUsername(u.UserName).
			SetEmail(u.Email).
			SetPassword(u.Password).
			Save(ctx); err != nil {
			return err
		}

		if userMetadata, err = tx.UserMetadata.
			Create().
			SetFirstName(u.FirstName).
			SetLastName(u.LastName).
			SetGender(u.Gender).
			SetUser(userModel).
			Save(ctx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &biz.User{
		ID:        userModel.ID,
		UserName:  userModel.Username,
		FirstName: userMetadata.FirstName,
		LastName:  userMetadata.LastName,
		Gender:    userMetadata.Gender,
		Email:     userModel.Email,
		Password:  userModel.Password,
		Disabled:  userModel.Disabled,
	}, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	u, err := r.data.db.User.Query().Where(user.Username(username)).Only(ctx)
	if err != nil {
		return nil, err
	}

	m, err := u.QueryMetadata().Only(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:        u.ID,
		UserName:  u.Username,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Gender:    m.Gender,
		Email:     u.Email,
		Password:  u.Password,
		Disabled:  u.Disabled,
	}, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*biz.User, error) {
	u, err := r.data.db.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	m, err := u.QueryMetadata().Only(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:        u.ID,
		UserName:  u.Username,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Gender:    m.Gender,
		Email:     u.Email,
		Password:  u.Password,
		Disabled:  u.Disabled,
	}, nil
}
