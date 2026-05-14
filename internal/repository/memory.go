// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package repository

import (
	"errors"
	"slices"
	"sync"

	"trackposter/internal/domain"
	"trackposter/internal/uid"
)

const (
	// trackNotFound returned when ID was not found in queue.
	trackNotFound int = -1
)

type MemoryQueue struct {
	mu    sync.RWMutex
	queue []*domain.TrackRecord
}

var _ domain.Repository = (*MemoryQueue)(nil)

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		mu:    sync.RWMutex{},
		queue: make([]*domain.TrackRecord, 0),
	}
}

// RemoveTrack looks up for track index and removes it from queue.
func (q *MemoryQueue) RemoveTrack(trackID string) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	idx := q.trackIndex(trackID)
	if idx == trackNotFound {
		return domain.ErrTrackNotFound
	}

	q.queue = slices.Delete(q.queue, idx, idx+1)

	return nil
}

// TrackByID finds and returns track from queue . If track not found - returns nil.
func (q *MemoryQueue) TrackByID(trackID string) *domain.Track {
	q.mu.RLock()
	defer q.mu.RUnlock()

	track := q.trackByID(trackID)

	return track
}

// Add new track to queue.
// Returns track id and error(or nil).
func (q *MemoryQueue) AddTrack(track *domain.Track) (string, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// generating id for track
	id, err := uid.New()
	if err != nil {
		return "", errors.Join(domain.ErrUUID, err)
	}

	q.queue = append(q.queue, &domain.TrackRecord{
		ID:    id,
		Track: track,
	})

	return id, nil
}

// Queue returns current tracks queue.
func (q *MemoryQueue) Queue() []*domain.Track {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return q.tracksFromQueue()
}

// trackIndex looks for track index in queue. Can return -1 if nothing was found.
func (q *MemoryQueue) trackIndex(trackID string) int {
	return slices.IndexFunc(q.queue, func(t *domain.TrackRecord) bool {
		return t.ID == trackID
	})
}

func (q *MemoryQueue) trackByID(trackID string) *domain.Track {
	idx := q.trackIndex(trackID)
	if idx == trackNotFound {
		return nil
	}

	return q.queue[idx].Track
}

func (q *MemoryQueue) tracksFromQueue() []*domain.Track {
	tracks := make([]*domain.Track, 0, len(q.queue))
	for _, record := range q.queue {
		tracks = append(tracks, record.Track)
	}

	return tracks
}
