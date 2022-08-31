package migrator

type (
	MigrateFn interface {
		// Init 마이그레이션 관련 테이블 및 데이터를 초기화 한다.
		Init() error
	}
)
