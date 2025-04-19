package repository

type MasterRepository interface {
	ReplicaRepository
	Create()
	Update()
	Delete()
}

type ReplicaRepository interface {
	Read()
	GetAll()
	Filter()
}


type repository struct {
	dbUrl string
}

func NewMasterRepository(dbUrl string) MasterRepository {
	return &repository{dbUrl: dbUrl}
}

func NewReplicaRepository(dbUrl string) ReplicaRepository {
	return &repository{dbUrl: dbUrl}
}

