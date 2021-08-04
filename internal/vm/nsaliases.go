package vm

func (ns *NS) SetAlias(key string, val interface{}) {
	ns.VM.Debug("ALIAS(%s) OPTION %v=%v ", ns.Name, key, val)
	ns.Aliases.Store(key, val)
}

func (ns *NS) GetAlias(key string) interface{} {
	if res, ok := ns.Aliases.Load(key); ok {
		ns.VM.Debug("ALIAS(%s)GET %v ", ns.Name, key)
		return res
	}
	return nil
}
