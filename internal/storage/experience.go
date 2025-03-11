package storage

func GetExperices(storage *Storage, filter interface{}) []Experience {
	var experiences []*Experience

	cursor, err := storage.collection.Find(ctx, filter)
}
