package store

import (
	"path/filepath"

	"github.com/youkre/firefly/config"
	"github.com/youkre/firefly/meta"
)

type Store struct {
	Meta    *Dir
	Content *Dir
}

func NewStore(dataPath *config.DataPath, project config.ProjectConfig) *Store {
	return &Store{
		Meta: NewDir(filepath.Join(
			dataPath.Meta,
			project.DataDir,
		)),
		Content: NewDir(filepath.Join(
			dataPath.Content,
			project.DataDir,
		)),
	}
}

func (s *Store) LoadIndex() (meta.IndexPage, error) {
	var idxPage meta.IndexPage

	err := s.Meta.ReadJSON("index.json", &idxPage)
	if err != nil {
		return meta.IndexPage{}, err
	}

	return idxPage, nil
}

func (s *Store) SaveIndex(idxPage meta.IndexPage) error {
	_, err := s.Meta.WriteJSON("index.json", idxPage)
	if err != nil {
		return err
	}

	return nil
}
