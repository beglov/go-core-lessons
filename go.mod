module go-core-lessons

go 1.15

replace (
	go-core-lessons/lesson-7/pkg/crawler v0.0.0 => ./lesson-7/pkg/crawler
	go-core-lessons/lesson-7/pkg/crawler/spider v0.0.0 => ./lesson-7/pkg/crawler/spider
	go-core-lessons/lesson-7/pkg/engine v0.0.0 => ./lesson-7/pkg/engine
	go-core-lessons/lesson-7/pkg/index v0.0.0 => ./lesson-7/pkg/index
	go-core-lessons/lesson-7/pkg/index/hash v0.0.0 => ./lesson-7/pkg/index/hash
	go-core-lessons/lesson-7/pkg/storage v0.0.0 => ./lesson-7/pkg/storage
	go-core-lessons/lesson-7/pkg/storage/btree v0.0.0 => ./lesson-7/pkg/storage/btree
)

require golang.org/x/net v0.0.0-20201224014010-6772e930b67b
