package btree

import (
	"go-core-lessons/lesson-14/pkg/crawler"
	"sync"
)

// Tree - Двоичное дерево поиска
type Tree struct {
	mux  *sync.Mutex
	root *Element
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	Value       crawler.Document
}

// New - конструктор.
func New() *Tree {
	db := Tree{
		mux: new(sync.Mutex),
	}
	return &db
}

// Clear очищает хранилище
func (t *Tree) Clear() {
	t.root = nil
}

// StoreDocs добавляет новые документы.
func (t *Tree) StoreDocs(docs []crawler.Document) error {
	for _, doc := range docs {
		t.Insert(doc)
	}
	return nil
}

// Docs возвращает документы по их номерам.
func (t *Tree) Docs(ids []int) []crawler.Document {
	var result []crawler.Document
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, id := range ids {
		s := t.Search(id)
		result = append(result, s)
	}
	return result
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(doc crawler.Document) {
	e := &Element{Value: doc}
	if t.root == nil {
		t.root = e
		return
	}
	insert(t.root, e)
}

// inset рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.Value.ID < node.Value.ID {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Value.ID >= node.Value.ID {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search - поиск значения в дереве, выдаёт документ если найдено, иначе nil
func (t *Tree) Search(x int) crawler.Document {
	return search(t.root, x)
}

func search(el *Element, x int) crawler.Document {
	if el == nil {
		return crawler.Document{}
	}
	if el.Value.ID == x {
		return el.Value
	}
	if el.Value.ID < x {
		return search(el.right, x)
	}
	return search(el.left, x)
}
