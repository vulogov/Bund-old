package vm

func (ns *NS) InLCache(name string) bool {
	if _, ok := ns.LCache.Load(name); ok {
		return true
	}
	return false
}

func (ns *NS) AddToLCache(name string, e *Elem) {
	if !ns.InLCache(name) {
		c := make([]Elem, 0)
		c = append(c, *e)
		ns.LCache.Store(name, c)
	} else {
		if _c, ok := ns.LCache.Load(name); ok {
			c := _c.([]Elem)
			c = append(c, *e)
			ns.LCache.Store(name, c)
		}
	}
}

func (ns *NS) LCacheGet(name string) []Elem {
	if c, ok := ns.LCache.Load(name); ok {
		for i := 0; i < len(c.([]Elem)); i++ {
			ns.VM.Debug("f(%v %v:%v)", name, i, c.([]Elem)[i])
		}
		return c.([]Elem)
	}
	return nil
}

func (ns *NS) LCacheDel(name string) {
	ns.VM.Debug("Removing %v from LCache", name)
	ns.LCache.Delete(name)
}
