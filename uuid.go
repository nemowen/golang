package main

func UUID() (string, error) {
	file, err := os.Open("/dev/urandom")

	if err != nil {
		return "", err
	}

	defer file.Close()

	b := make([]byte, 16)
	file.Read(b)

	uuid := fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}
