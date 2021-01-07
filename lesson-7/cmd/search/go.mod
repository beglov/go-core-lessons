module go-core-lessons/lesson-7/cmd/search

go 1.15

replace (
	gosearch/pkg/crawler => ../../pkg/crawler
	gosearch/pkg/crawler/spider => ../../pkg/crawler/spider
	gosearch/pkg/engine => ../../pkg/engine
	gosearch/pkg/index => ../../pkg/index
	gosearch/pkg/index/hash => ../../pkg/index/hash
	gosearch/pkg/storage => ../../pkg/storage
	gosearch/pkg/storage/btree => ../../pkg/storage/btree
)

require (
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000
	gosearch/pkg/crawler/spider v0.0.0-00010101000000-000000000000
	gosearch/pkg/engine v0.0.0-00010101000000-000000000000
	gosearch/pkg/index v0.0.0-00010101000000-000000000000
	gosearch/pkg/index/hash v0.0.0-00010101000000-000000000000
	gosearch/pkg/storage v0.0.0-00010101000000-000000000000
	gosearch/pkg/storage/btree v0.0.0-00010101000000-000000000000
)
