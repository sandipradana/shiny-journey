package repository

import (
	"context"
	"shiny-journey/ent"
	"shiny-journey/ent/admin"
)

type AdminRepository interface {
	GetAdminByID(ctx context.Context, id int) (*ent.Admin, error)
	GetAdminByEmail(ctx context.Context, email string) (*ent.Admin, error)
	CreateAdmin(ctx context.Context, admin *ent.Admin) (*ent.Admin, error)
}

type adminRepository struct {
	client *ent.Client
}

func NewAdminRepository(client *ent.Client) AdminRepository {
	return &adminRepository{
		client: client,
	}
}

func (r *adminRepository) GetAdminByID(ctx context.Context, id int) (*ent.Admin, error) {
	admin, err := r.client.Admin.
		Query().
		Where(admin.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) GetAdminByEmail(ctx context.Context, email string) (*ent.Admin, error) {
	admin, err := r.client.Admin.
		Query().
		Where(admin.EmailEQ(email)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *adminRepository) CreateAdmin(ctx context.Context, admin *ent.Admin) (*ent.Admin, error) {
	createdAdmin, err := r.client.Admin.
		Create().
		SetName(admin.Name).
		SetEmail(admin.Email).
		SetPassword(admin.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return createdAdmin, nil
}
