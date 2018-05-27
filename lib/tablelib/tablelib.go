package tablelib

import rt "github.com/arnodel/golua/runtime"

func Load(r *rt.Runtime) {
	pkg := rt.NewTable()
	rt.SetEnv(r.GlobalEnv(), "table", pkg)
	rt.SetEnvGoFunc(pkg, "concat", concat, 4, false)
	rt.SetEnvGoFunc(pkg, "insert", insert, 3, false)
	rt.SetEnvGoFunc(pkg, "move", move, 5, false)
	rt.SetEnvGoFunc(pkg, "pack", pack, 0, true)
	rt.SetEnvGoFunc(pkg, "remove", remove, 2, false)
}

func concat(t *rt.Thread, c *rt.GoCont) (rt.Cont, *rt.Error) {
	if err := c.Check1Arg(); err != nil {
		return nil, err.AddContext(c)
	}
	tbl, err := c.TableArg(0)
	if err != nil {
		return nil, err.AddContext(c)
	}
	var sep rt.String
	i := rt.Int(1)
	j, err := rt.Len(t, tbl)
	if err != nil {
		return nil, err.AddContext(c)
	}
Switch:
	switch nargs := c.NArgs(); {
	case nargs >= 4:
		j, err = c.IntArg(3)
		if err != nil {
			break
		}
		fallthrough
	case nargs >= 3:
		i, err = c.IntArg(2)
		if err != nil {
			break
		}
		fallthrough
	case nargs >= 2:
		sep, err = c.StringArg(1)
		if err != nil {
			break
		}
		fallthrough
	default:
		var res rt.Value
		if i > j {
			return c.Next(), nil
		}
		res, err = rt.Index(t, tbl, i)
		if err != nil {
			break
		}
		for i++; i <= j; i++ {
			res, err = rt.Concat(t, res, sep)
			if err != nil {
				break Switch
			}
			v, err := rt.Index(t, tbl, i)
			if err != nil {
				break Switch
			}
			res, err = rt.Concat(t, res, v)
			if err != nil {
				break Switch
			}
		}
		return c.PushingNext(res), nil
	}
	return nil, err.AddContext(c)
}

func insert(t *rt.Thread, c *rt.GoCont) (rt.Cont, *rt.Error) {
	if err := c.CheckNArgs(2); err != nil {
		return nil, err.AddContext(c)
	}
	tbl, err := c.TableArg(0)
	if err != nil {
		return nil, err.AddContext(c)
	}
	var val rt.Value
	var pos rt.Int
	tblLen, err := rt.Len(t, tbl)
	if err != nil {
		return nil, err.AddContext(c)
	}
	if c.NArgs() >= 3 {
		pos, err = c.IntArg(1)
		if err != nil {
			return nil, err.AddContext(c)
		}
		if pos <= 0 {
			return nil, rt.NewErrorS("#2 out of range").AddContext(c)
		}
		val = c.Arg(2)
	} else {
		pos = tblLen + 1
		val = c.Arg(1)
	}
	var oldVal rt.Value
	for pos <= tblLen {
		oldVal, err = rt.Index(t, tbl, pos)
		if err != nil {
			return nil, err.AddContext(c)
		}
		err = rt.SetIndex(t, tbl, pos, val)
		if err != nil {
			return nil, err.AddContext(c)
		}
		val = oldVal
		pos++
	}
	err = rt.SetIndex(t, tbl, pos, val)
	if err != nil {
		return nil, err.AddContext(c)
	}
	return c.Next(), nil
}

func move(t *rt.Thread, c *rt.GoCont) (rt.Cont, *rt.Error) {
	if err := c.CheckNArgs(4); err != nil {
		return nil, err.AddContext(c)
	}
	src, err := c.TableArg(0)
	if err != nil {
		return nil, err.AddContext(c)
	}
	srcStart, err := c.IntArg(1)
	if err != nil {
		return nil, err.AddContext(c)
	}
	srcEnd, err := c.IntArg(2)
	if err != nil {
		return nil, err.AddContext(c)
	}
	dstStart, err := c.IntArg(3)
	if err != nil {
		return nil, err.AddContext(c)
	}
	dst := src
	if c.NArgs() >= 5 {
		dst, err = c.TableArg(4)
		if err != nil {
			return nil, err.AddContext(c)
		}
	}
	if srcStart > srcEnd {
		// Nothing to do apparently!
	} else if dstStart >= srcStart {
		// Move in descending order to avoid writing at a position
		// before moving it
		dstStart += srcEnd - srcStart
		for srcEnd >= srcStart {
			v, err := rt.Index(t, src, srcEnd)
			if err == nil {
				err = rt.SetIndex(t, dst, dstStart, v)
			}
			if err != nil {
				return nil, err.AddContext(c)
			}
			srcEnd--
			dstStart--
		}
	} else {
		// Move in ascending order to avoid writing at a position
		// before moving it
		for srcStart <= srcEnd {
			v, err := rt.Index(t, src, srcStart)
			if err == nil {
				err = rt.SetIndex(t, dst, dstStart, v)
			}
			if err != nil {
				return nil, err.AddContext(c)
			}
			srcStart++
			dstStart++
		}
	}
	return c.PushingNext(dst), nil
}

func pack(t *rt.Thread, c *rt.GoCont) (rt.Cont, *rt.Error) {
	tbl := rt.NewTable()
	// We can use tbl.Set() because tbl has no metatable
	for i, v := range c.Etc() {
		tbl.Set(rt.Int(i+1), v)
	}
	tbl.Set(rt.String("n"), rt.Int(len(c.Etc())))
	return c.PushingNext(tbl), nil
}

func remove(t *rt.Thread, c *rt.GoCont) (rt.Cont, *rt.Error) {
	if err := c.Check1Arg(); err != nil {
		return nil, err.AddContext(c)
	}
	tbl, err := c.TableArg(0)
	if err != nil {
		return nil, err.AddContext(c)
	}
	tblLen, err := rt.Len(t, tbl)
	if err != nil {
		return nil, err.AddContext(c)
	}
	pos := tblLen
	if c.NArgs() >= 2 {
		pos, err = c.IntArg(1)
		if err != nil {
			return nil, err.AddContext(c)
		}
	}
	var val rt.Value
	switch {
	case pos == tblLen || pos == tblLen+1:
		val, err = rt.Index(t, tbl, pos)
		if err == nil {
			err = rt.SetIndex(t, tbl, pos, nil)
		}
		if err != nil {
			return nil, err.AddContext(c)
		}
	case pos <= 0 || pos > tblLen:
		return nil, rt.NewErrorS("#2 out of range").AddContext(c)
	default:
		var newVal rt.Value = nil
		for pos <= tblLen {
			val, err = rt.Index(t, tbl, tblLen)
			rt.SetIndex(t, tbl, tblLen, newVal)
			tblLen--
			newVal = val
		}
	}
	return c.PushingNext(val), nil
}