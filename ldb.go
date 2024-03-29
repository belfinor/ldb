package ldb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.011
// @date    2019-11-11

var store Storage = nil
var proxyConfig *Config

func Init(cfg *Config) {
	if store == nil {
		New(cfg, true)
	}
}

func TestInit() {
	Open("test=true")
}

func Close() {
	if store != nil {
		store.Close()
	}
	store = nil
}

func InitProxy(cfg *Config) {
	proxyConfig = cfg
}

func Use(base string) {

	cfg := &Config{
		Path:        proxyConfig.Path + "/" + base,
		FileSize:    proxyConfig.FileSize,
		Compression: proxyConfig.Compression,
		ReadOnly:    proxyConfig.ReadOnly,
	}

	Close()
	Init(cfg)
}

func Set(key []byte, value []byte) {
	store.Set(key, value)
}

func Get(key []byte) []byte {
	return store.Get(key)
}

func Has(key []byte) bool {
	return store.Has(key)
}

func Del(key []byte) {
	store.Del(key)
}

func Total(prefix []byte) int64 {
	return store.Total(prefix)
}

func List(prefix []byte, limit int, offset int, RemovePrefix bool) [][]byte {

	return store.List(prefix, limit, offset, RemovePrefix)
}

func ForEach(prefix []byte, RemovePrefix bool, fn FOR_EACH_FUNC) {

	store.ForEach(prefix, RemovePrefix, fn)
}

func ForEachKey(prefix []byte, limit int, offset int, RemovePrefix bool, fn FOR_EACH_KEY_FUNC) {

	store.ForEachKey(prefix, limit, offset, RemovePrefix, fn)
}
