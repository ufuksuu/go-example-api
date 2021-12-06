package server

func Init() {
	r := NewRouter()
	err := r.Run()
	if err != nil {
		panic("Failed to start routes!")
	}
}
