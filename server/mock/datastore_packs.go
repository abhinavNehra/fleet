// Automatically generated by mockimpl. DO NOT EDIT!

package mock

import "github.com/kolide/fleet/server/kolide"

var _ kolide.PackStore = (*PackStore)(nil)

type ApplyPackSpecFunc func(spec *kolide.PackSpec) error

type NewPackFunc func(pack *kolide.Pack, opts ...kolide.OptionalArg) (*kolide.Pack, error)

type SavePackFunc func(pack *kolide.Pack) error

type DeletePackFunc func(pid uint) error

type PackFunc func(pid uint) (*kolide.Pack, error)

type ListPacksFunc func(opt kolide.ListOptions) ([]*kolide.Pack, error)

type PackByNameFunc func(name string, opts ...kolide.OptionalArg) (*kolide.Pack, bool, error)

type AddLabelToPackFunc func(lid uint, pid uint, opts ...kolide.OptionalArg) error

type RemoveLabelFromPackFunc func(lid uint, pid uint) error

type ListLabelsForPackFunc func(pid uint) ([]*kolide.Label, error)

type AddHostToPackFunc func(hid uint, pid uint) error

type RemoveHostFromPackFunc func(hid uint, pid uint) error

type ListHostsInPackFunc func(pid uint, opt kolide.ListOptions) ([]uint, error)

type ListExplicitHostsInPackFunc func(pid uint, opt kolide.ListOptions) ([]uint, error)

type PackStore struct {
	ApplyPackSpecFunc        ApplyPackSpecFunc
	ApplyPackSpecFuncInvoked bool

	NewPackFunc        NewPackFunc
	NewPackFuncInvoked bool

	SavePackFunc        SavePackFunc
	SavePackFuncInvoked bool

	DeletePackFunc        DeletePackFunc
	DeletePackFuncInvoked bool

	PackFunc        PackFunc
	PackFuncInvoked bool

	ListPacksFunc        ListPacksFunc
	ListPacksFuncInvoked bool

	PackByNameFunc        PackByNameFunc
	PackByNameFuncInvoked bool

	AddLabelToPackFunc        AddLabelToPackFunc
	AddLabelToPackFuncInvoked bool

	RemoveLabelFromPackFunc        RemoveLabelFromPackFunc
	RemoveLabelFromPackFuncInvoked bool

	ListLabelsForPackFunc        ListLabelsForPackFunc
	ListLabelsForPackFuncInvoked bool

	AddHostToPackFunc        AddHostToPackFunc
	AddHostToPackFuncInvoked bool

	RemoveHostFromPackFunc        RemoveHostFromPackFunc
	RemoveHostFromPackFuncInvoked bool

	ListHostsInPackFunc        ListHostsInPackFunc
	ListHostsInPackFuncInvoked bool

	ListExplicitHostsInPackFunc        ListExplicitHostsInPackFunc
	ListExplicitHostsInPackFuncInvoked bool
}

func (s *PackStore) ApplyPackSpec(spec *kolide.PackSpec) error {
	s.ApplyPackSpecFuncInvoked = true
	return s.ApplyPackSpecFunc(spec)
}

func (s *PackStore) NewPack(pack *kolide.Pack, opts ...kolide.OptionalArg) (*kolide.Pack, error) {
	s.NewPackFuncInvoked = true
	return s.NewPackFunc(pack, opts...)
}

func (s *PackStore) SavePack(pack *kolide.Pack) error {
	s.SavePackFuncInvoked = true
	return s.SavePackFunc(pack)
}

func (s *PackStore) DeletePack(pid uint) error {
	s.DeletePackFuncInvoked = true
	return s.DeletePackFunc(pid)
}

func (s *PackStore) Pack(pid uint) (*kolide.Pack, error) {
	s.PackFuncInvoked = true
	return s.PackFunc(pid)
}

func (s *PackStore) ListPacks(opt kolide.ListOptions) ([]*kolide.Pack, error) {
	s.ListPacksFuncInvoked = true
	return s.ListPacksFunc(opt)
}

func (s *PackStore) PackByName(name string, opts ...kolide.OptionalArg) (*kolide.Pack, bool, error) {
	s.PackByNameFuncInvoked = true
	return s.PackByNameFunc(name, opts...)
}

func (s *PackStore) AddLabelToPack(lid uint, pid uint, opts ...kolide.OptionalArg) error {
	s.AddLabelToPackFuncInvoked = true
	return s.AddLabelToPackFunc(lid, pid, opts...)
}

func (s *PackStore) RemoveLabelFromPack(lid uint, pid uint) error {
	s.RemoveLabelFromPackFuncInvoked = true
	return s.RemoveLabelFromPackFunc(lid, pid)
}

func (s *PackStore) ListLabelsForPack(pid uint) ([]*kolide.Label, error) {
	s.ListLabelsForPackFuncInvoked = true
	return s.ListLabelsForPackFunc(pid)
}

func (s *PackStore) AddHostToPack(hid uint, pid uint) error {
	s.AddHostToPackFuncInvoked = true
	return s.AddHostToPackFunc(hid, pid)
}

func (s *PackStore) RemoveHostFromPack(hid uint, pid uint) error {
	s.RemoveHostFromPackFuncInvoked = true
	return s.RemoveHostFromPackFunc(hid, pid)
}

func (s *PackStore) ListHostsInPack(pid uint, opt kolide.ListOptions) ([]uint, error) {
	s.ListHostsInPackFuncInvoked = true
	return s.ListHostsInPackFunc(pid, opt)
}

func (s *PackStore) ListExplicitHostsInPack(pid uint, opt kolide.ListOptions) ([]uint, error) {
	s.ListExplicitHostsInPackFuncInvoked = true
	return s.ListExplicitHostsInPackFunc(pid, opt)
}
