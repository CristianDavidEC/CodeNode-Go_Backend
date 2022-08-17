package repository

type Repository interface {
	GetAllProgramsDb() ([]byte, error)
	GetProgramDb(id string) ([]byte, error)
	SaveProgramDb(newProgram []byte) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func GetAllProgramsDb() ([]byte, error) {
	return implementation.GetAllProgramsDb()
}

func GetProgramDb(id string) ([]byte, error) {
	return implementation.GetProgramDb(id)
}

func SaveProgramDb(newProgram []byte) error {
	return implementation.SaveProgramDb(newProgram)
}
