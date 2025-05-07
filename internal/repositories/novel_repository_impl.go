package repositories

type NovelRepositoryImpl struct {
	str string
}

func (nri *NovelRepositoryImpl) Save() {
}

func NewNovelRepository() NovelRepository {
	return &NovelRepositoryImpl{str: "Hello world\n"}
}
