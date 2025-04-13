package topic

type TopicId string

type Topic interface {
	Insert([]byte) (int, error)
	Extract(int) ([]byte, error)
	ExtractLatest() ([]byte, error)
}
