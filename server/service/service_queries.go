package service

import (
	"context"

	"github.com/kolide/fleet/server/contexts/viewer"
	"github.com/kolide/fleet/server/kolide"
	"github.com/pkg/errors"
)

func (svc service) ApplyQueryYaml(ctx context.Context, yml string) error {
	vc, ok := viewer.FromContext(ctx)
	if !ok {
		return errors.New("user must be authenticated to apply queries")
	}

	queries, err := kolide.LoadQueriesFromYaml(yml)
	if err != nil {
		return errors.Wrap(err, "loading query yaml")
	}

	err = svc.ds.ApplyQueries(vc.UserID(), queries)
	return errors.Wrap(err, "applying queries")
}

func (svc service) GetQueryYaml(ctx context.Context) (string, error) {
	queries, err := svc.ds.ListQueries(kolide.ListOptions{})
	if err != nil {
		return "", errors.Wrap(err, "getting queries")
	}

	yml, err := kolide.WriteQueriesToYaml(queries)
	if err != nil {
		return "", errors.Wrap(err, "writing query yaml")
	}

	return yml, nil
}

func (svc service) ListQueries(ctx context.Context, opt kolide.ListOptions) ([]*kolide.Query, error) {
	return svc.ds.ListQueries(opt)
}

func (svc service) GetQuery(ctx context.Context, id uint) (*kolide.Query, error) {
	return svc.ds.Query(id)
}

func (svc service) NewQuery(ctx context.Context, p kolide.QueryPayload) (*kolide.Query, error) {
	query := &kolide.Query{Saved: true}

	if p.Name != nil {
		query.Name = *p.Name
	}

	if p.Description != nil {
		query.Description = *p.Description
	}

	if p.Query != nil {
		query.Query = *p.Query
	}

	vc, ok := viewer.FromContext(ctx)
	if ok {
		query.AuthorID = vc.UserID()
		query.AuthorName = vc.FullName()
	}

	query, err := svc.ds.NewQuery(query)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (svc service) ModifyQuery(ctx context.Context, id uint, p kolide.QueryPayload) (*kolide.Query, error) {
	query, err := svc.ds.Query(id)
	if err != nil {
		return nil, err
	}

	if p.Name != nil {
		query.Name = *p.Name
	}

	if p.Description != nil {
		query.Description = *p.Description
	}

	if p.Query != nil {
		query.Query = *p.Query
	}

	err = svc.ds.SaveQuery(query)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (svc service) DeleteQuery(ctx context.Context, id uint) error {
	return svc.ds.DeleteQuery(id)
}

func (svc service) DeleteQueries(ctx context.Context, ids []uint) (uint, error) {
	return svc.ds.DeleteQueries(ids)
}
