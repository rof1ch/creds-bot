package session

import "time"

type Session struct {
	UserID     int64
	LastUsed   time.Time
	DecryptKey string
}

type ListSession map[int64]*Session

func NewList() *ListSession {
	return &ListSession{}
}

func (ls ListSession) NewSession(userID int64, key string) {
	ls[userID] = &Session{
		UserID:     int64(userID),
		DecryptKey: key,
		LastUsed:   time.Now(),
	}

}

func (ls ListSession) GetSession(userID int64) (*Session, bool) {
	session, ok := ls[userID]
	if !ok {
		return nil, false
	}
	return session, true

}
func (ls ListSession) UpdateLastUsed(userID int64) {
	ls[userID].LastUsed = time.Now()
}

func (ls ListSession) NeedsReauth(userID int64) bool {
	session, _ := ls.GetSession(userID)
	return session.LastUsed.Before(time.Now().Add(time.Minute * -5))
}
