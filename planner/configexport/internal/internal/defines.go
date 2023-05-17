package internal

type XXConfig struct {
	Id    int
	Count string
	Award map[int]string
	Info  *struct {
		Id   int
		Name string
		Info *struct {
			Lv  int
			Exp *struct {
				Mux   int
				Count int
			}
		}
	}
	Other *struct {
		Id   int
		Name string
	}
}
