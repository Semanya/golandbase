package config

func StructToMap(c Config) map[string]interface{} {
	out := make(map[string]interface{})
	out["Name"] = c.Name
	out["Host"] = c.Host
	out["Port"] = c.Port
	out["Environment"] = c.Environment
	out["Volumes"] = c.Volumes
	return out
}
