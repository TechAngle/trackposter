// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package repository

import (
	"fmt"
	"slices"
	"sync"
	"trackposter/internal/server/models"
	"trackposter/internal/utils"
)

type TrackRecord struct {
	Track *models.Track
	ID    string
}

func (r *TrackRecord) ToTrack() *models.Track {
	return r.Track
}

type MemoryQueue struct {
	mutex sync.RWMutex
	queue []*TrackRecord
}

// looks for track index in queue. Can return -1 if nothing found
func (q *MemoryQueue) trackIndex(trackId string) int {
	return slices.IndexFunc(q.queue, func(t *TrackRecord) bool {
		return t.ID == trackId
	})
}

// finds track by its id
func (q *MemoryQueue) trackById(trackId string) *models.Track {
	idx := q.trackIndex(trackId)
	if idx == -1 {
		return nil
	}

	return q.queue[idx].Track
}

// gets all tracks from queue and returns its model
func (q *MemoryQueue) tracksFromQueue() []*models.Track {
	tracks := make([]*models.Track, 0, len(q.queue))
	for _, record := range q.queue {
		tracks = append(tracks, record.ToTrack())
	}

	return tracks
}

// Remove track using its ID
func (q *MemoryQueue) RemoveTrack(trackId string) error {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	idx := q.trackIndex(trackId)
	if idx == -1 {
		return fmt.Errorf("track with ID %s do not exist", trackId)
	}

	q.queue = slices.Delete(q.queue, idx, idx+1)

	return nil
}

// Finds and returns track by its ID. If track not found - returns nil.
func (q *MemoryQueue) TrackByID(trackId string) *models.Track {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	track := q.trackById(trackId)
	return track
}

// Add new track to queue.
// Returns track id and error(or nil).
func (q *MemoryQueue) AddTrack(track *models.Track) (string, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// generating id for track
	id, err := utils.GenerateUUID()
	if err != nil {
		return "", fmt.Errorf("failed to generate track id: %v", err)
	}

	q.queue = append(q.queue, &TrackRecord{
		ID:    id,
		Track: track,
	})

	return id, nil
}

// Gets tracks queue
func (q *MemoryQueue) Queue() []*models.Track {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	return q.tracksFromQueue()
}

func NewMemoryQueue() *MemoryQueue {
	return &MemoryQueue{
		queue: make([]*TrackRecord, 0),
	}
}
