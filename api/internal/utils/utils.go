package utils

func IsProductionEnv(env string) bool {
	return env == "production"
}
