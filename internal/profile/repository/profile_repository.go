package repository

import (
	"context"
	"database/sql"

	entityAuth "server/internal/auth/entity"
	dtoProfile "server/internal/profile/dto"
	"server/pkg/database/transactor"
)

type ProfileRepository interface {
	PutMyProfile(ctx context.Context, reqBody *dtoProfile.RequestPutMyProfile, userId int64, roleId int64, imgUrl *string) (*entityAuth.UserDetail, error)
}

type profileRepositoryImpl struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *profileRepositoryImpl {
	return &profileRepositoryImpl{
		db: db,
	}
}

func (ar *profileRepositoryImpl) PutMyProfile(ctx context.Context, reqBody *dtoProfile.RequestPutMyProfile, userId int64, roleId int64, imgUrl *string) (*entityAuth.UserDetail, error) {
	query := `
		UPDATE user_details
		SET full_name = COALESCE(NULLIF($1, ''), full_name),
			whatsapp_number = COALESCE(NULLIF($2, ''), whatsapp_number),
			image_url = COALESCE(NULLIF($3, ''), image_url),
			updated_at = NOW()
		WHERE user_id = $4
		RETURNING id, user_id, full_name, sipa_number, whatsapp_number, years_of_experience, image_url, created_at, updated_at, deleted_at;
    `
	tx := transactor.ExtractTx(ctx)
	var userDetail entityAuth.UserDetail
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, reqBody.Fullname, reqBody.WhatsappNumber, *imgUrl, userId).Scan(
			&userDetail.ID,
			&userDetail.UserId,
			&userDetail.Fullname,
			&userDetail.WhatsappNumber,
			&userDetail.ImageUrl,
			&userDetail.CreatedAt,
			&userDetail.UpdatedAt,
			&userDetail.DeletedAt,
		)
	} else {
		err = ar.db.QueryRowContext(ctx, query, reqBody.Fullname, reqBody.WhatsappNumber, *imgUrl, userId).Scan(
			&userDetail.ID,
			&userDetail.UserId,
			&userDetail.Fullname,
			&userDetail.WhatsappNumber,
			&userDetail.ImageUrl,
			&userDetail.CreatedAt,
			&userDetail.UpdatedAt,
			&userDetail.DeletedAt,
		)
	}
	if err != nil {
		return nil, err
	}

	return &userDetail, nil
}
