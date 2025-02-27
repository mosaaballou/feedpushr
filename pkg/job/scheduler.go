package job

import (
	"github.com/ncarlier/feedpushr/pkg/config"
	"github.com/ncarlier/feedpushr/pkg/store"
)

// Scheduler is a job scheduler
type Scheduler struct {
	cleanCacheJob *CleanCacheJob
}

// StartNewScheduler create and start new job scheduler
func StartNewScheduler(_db store.DB, conf config.Config) *Scheduler {
	return &Scheduler{
		cleanCacheJob: NewCleanCacheJob(_db, conf.CacheRetention),
	}
}

// Shutdown job scheduler
func (s *Scheduler) Shutdown() {
	s.cleanCacheJob.Stop()
}
