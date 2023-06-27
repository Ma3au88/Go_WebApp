/* Пакет model должен определить структуру, а указатель на экземпляр этой структуры должен быть передан
всем функциям и методам ui. В нашем приложении должен быть только один такой экземпляр — для дополнительной
уверенности вы можете реализовать это программно с помощью синглтона, но я не думаю, что это так уж необходимо.
*/

package model

type db interface {
	SelectPeople() ([]*Person, error)
}

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) People() ([]*Person, error) {
	return m.SelectPeople()
}

type Person struct {
	Id          int64
	First, Last string
}
