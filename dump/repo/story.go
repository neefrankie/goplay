package repo

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"example.com/dump/pkg"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sem        = semaphore.NewWeighted(int64(maxWorkers))
)

type StoryRepo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) StoryRepo {
	return StoryRepo{
		db: db,
	}
}

const stmtStory = `
SELECT story.id AS id,
    story.fileupdatetime AS created_at,
    story.last_publish_time AS updated_at,
    story.accessright AS access_right,
    story.cheadline AS title_cn,
    story.eheadline AS title_en,
    story.cbyline_description AS byline_desc_cn,
    story.ebyline_description AS byline_desc_en,
    story.cauthor AS byline_author_cn,
    story.eauthor AS byline_author_en,
    story.cbyline_status AS byline_status_cn,
    story.ebyline_status AS byline_status_en,
    story.clongleadbody AS long_lead_cn,
    picture.piclink  AS cover_url,
    story.tag AS tag,
    story.area AS area,
    story.genre AS genre,
    story.industry AS industry,
    story.topic AS topic,
    story.cbody AS body_cn,
    story.ebody AS body_en
FROM cmstmp01.story AS story
    LEFT JOIN cmstmp01.story_pic AS storyToPic
        ON story.id = storyToPic.storyid 
    LEFT JOIN cmstmp01.picture AS picture
        ON picture.id = storyToPic.picture_id
WHERE story.id = ?
LIMIT 1`

func (repo StoryRepo) retrieveStory(id string) (pkg.Story, error) {

	var story pkg.RawStory

	if err := repo.db.Get(&story, stmtStory, id); err != nil {
		log.Println(err)
		return pkg.Story{}, err
	}

	story.Sanitize()
	story.Normalize()

	return pkg.NewStory(story), nil
}

func (repo StoryRepo) LoadAndSave(id string, dataDir string) error {

	log.Printf("Start retrieving story %s\n", id)

	story, err := repo.retrieveStory(id)
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(story, "", "\t")
	if err != nil {
		return err
	}

	return saveStory(buildFileName(id, dataDir), b)
}

func ReadIDs(fileName string) ([]string, error) {

	content, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	ids := strings.Split(string(content), "\n")

	return ids, nil
}

func (repo StoryRepo) StartDump(ids []string, dataDir string) {
	ctx := context.Background()

	for _, id := range ids {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(id string) {
			defer sem.Release(1)

			err := repo.LoadAndSave(id, dataDir)
			if err != nil {
				log.Printf("Error when load and save story %s\n", err)
				return
			}
		}(id)
	}

	err := sem.Acquire(ctx, int64(maxWorkers))
	if err != nil {
		log.Printf("Error waiting finish: %s", err)
		return
	}

	log.Println("Done")
}

func buildFileName(id string, dir string) string {
	return filepath.Join(dir, id+".json")
}

func saveStory(fileName string, data []byte) error {
	log.Printf("Saving file %s\n", fileName)
	return os.WriteFile(fileName, data, 0666)
}
