package converter

import "time"

type LogEntry struct {
	Level  string    `json:"level"`
	Msg    string    `json:"msg"`
	Source string    `json:"source"`
	Stats  *Stats    `json:"stats"`
	Time   time.Time `json:"time"`
}

type Stats struct {
	Bytes               int64   `json:"bytes"`
	Checks              int     `json:"checks"`
	DeletedDirs         int     `json:"deletedDirs"`
	Deletes             int     `json:"deletes"`
	ElapsedTime         float64 `json:"elapsedTime"`
	Errors              int     `json:"errors"`
	ETA                 int     `json:"eta"`
	FatalError          bool    `json:"fatalError"`
	Renames             int     `json:"renames"`
	RetryError          bool    `json:"retryError"`
	ServerSideCopies    int     `json:"serverSideCopies"`
	ServerSideCopyBytes int64   `json:"serverSideCopyBytes"`
	ServerSideMoveBytes int64   `json:"serverSideMoveBytes"`
	ServerSideMoves     int     `json:"serverSideMoves"`
	Speed               float64 `json:"speed"`
	TotalBytes          int64   `json:"totalBytes"`
	TotalChecks         int     `json:"totalChecks"`
	TotalTransfers      int     `json:"totalTransfers"`
	TransferTime        float64 `json:"transferTime"`
	Transfers           int     `json:"transfers"`
}
