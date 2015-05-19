package utils

func GetFilePath(base_path string, length int, extension string) string {
	file_name := base_path + "/" + RandomString(length) + "." + extension
	return file_name
}
