package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"port-radar/internal/models"
	"strconv"
	"time"
)

// InitDB 初始化数据库
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS app_marks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		port INTEGER NOT NULL,
		protocol TEXT NOT NULL DEFAULT 'tcp',
		app_name TEXT NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(port, protocol)
	);
	CREATE INDEX IF NOT EXISTS idx_port ON app_marks(port);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetAppMark 获取端口的标记
func GetAppMark(db *sql.DB, port int, protocol string) (*models.AppMark, error) {
	mark := &models.AppMark{}
	err := db.QueryRow(
		"SELECT id, port, protocol, app_name, description, created_at, updated_at FROM app_marks WHERE port = ? AND protocol = ?",
		port, protocol,
	).Scan(&mark.ID, &mark.Port, &mark.Protocol, &mark.AppName, &mark.Description, &mark.CreatedAt, &mark.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return mark, nil
}

// GetAllAppMarks 获取所有标记
func GetAllAppMarks(db *sql.DB) (map[string]*models.AppMark, error) {
	rows, err := db.Query("SELECT id, port, protocol, app_name, description, created_at, updated_at FROM app_marks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	marks := make(map[string]*models.AppMark)
	for rows.Next() {
		mark := &models.AppMark{}
		err := rows.Scan(&mark.ID, &mark.Port, &mark.Protocol, &mark.AppName, &mark.Description, &mark.CreatedAt, &mark.UpdatedAt)
		if err != nil {
			return nil, err
		}
		marks[mark.Protocol+":"+strconv.Itoa(mark.Port)] = mark
	}
	return marks, nil
}

// SaveAppMark 保存或更新标记
func SaveAppMark(db *sql.DB, port int, protocol, appName, description string) error {
	now := time.Now()
	_, err := db.Exec(`
		INSERT INTO app_marks (port, protocol, app_name, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(port, protocol) DO UPDATE SET
			app_name = excluded.app_name,
			description = excluded.description,
			updated_at = excluded.updated_at
	`, port, protocol, appName, description, now, now)
	return err
}

// DeleteAppMark 删除标记
func DeleteAppMark(db *sql.DB, port int, protocol string) error {
	_, err := db.Exec("DELETE FROM app_marks WHERE port = ? AND protocol = ?", port, protocol)
	return err
}
