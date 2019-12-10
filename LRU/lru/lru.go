package lru

import (
	"errors"
	"container/list"
)

const LEN  = 10


type  LruCache struct{
	size int
	dataList *list.List        //存的是value
	items  map[interface{}]*list.Element
}

type entry struct{
	key interface{}
	value interface{}
}


func NewLruCache(size int)(*LruCache, error){
	if(size<0){
		return nil, errors.New("Must provide a positive size")
	}

	c:= &LruCache{
		size:     size,
		dataList: list.New(),
		items:    make(map[interface{}]*list.Element),
	}
	return c,nil
}

func (l *LruCache) Contain(key interface{}) (bool){
	_,ok := l.items[key]
	return ok
}

func (l *LruCache)  Add(key ,value interface{}){
	if(l.Contain(key)){
		v := l.items[key]
		l.dataList.Remove(v)
		l.dataList.PushFront(&entry{key,value})
		l.items[key] = l.dataList.Front()
	}else{
		if(len(l.items) < l.size ) {
			l.dataList.PushFront(&entry{key,value})
			l.items[key] = l.dataList.Front()
		}else{
			entryValue := l.dataList.Back().Value.(*entry).key
			l.dataList.Remove(l.dataList.Back())
			delete(l.items,entryValue)
		}
	}
}

func (l *LruCache) Get(key interface{})(value interface{}){
	if(l.Contain(key)){
		temp := l.items[key]
		l.dataList.MoveToFront(temp)
		value := temp.Value.(*entry).value
		return value
	}
	return nil
}

func (l *LruCache) Remove(key interface{}){
	if(l.Contain(key)){
		temp := l.items[key]
		l.dataList.Remove(temp)
		delete(l.items,key)
	}
}

func (l *LruCache) String() (string){
	var res string
	for e := l.dataList.Front(); e != nil; e = e.Next() {
		key := (e.Value.(*entry).key).(string)
		value := e.Value.(*entry).value.(string)
		res += ( key + "对应的值为:"  + value + "\n")
	}
	return res
}


