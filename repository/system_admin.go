package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
)

var SystemAdminR SystemAdminRepository

func RegisterSystemAdmin(p SystemAdminRepository) error {
	if SystemAdminR != nil {
		return errors.New("repository: RegisterSystemAdmin called twice")
	}
	SystemAdminR = p
	return nil
}

func CreateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) (int, error) {
	return SystemAdminR.CreateSystemAdmin(ctx, SystemAdmin)
}

func GetSystemAdmin(ctx context.Context, id int) (entity.SystemAdmin, error) {
	return SystemAdminR.GetSystemAdmin(ctx, id)
}

func GetSystemAdmins(ctx context.Context) ([]entity.SystemAdmin, error) {
	return SystemAdminR.GetSystemAdmins(ctx)
}

func UpdateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) error {
	return SystemAdminR.UpdateSystemAdmin(ctx, SystemAdmin)
}

func DeleteSystemAdmin(ctx context.Context, id int) error {
	return SystemAdminR.DeleteSystemAdmin(ctx, id)
}
