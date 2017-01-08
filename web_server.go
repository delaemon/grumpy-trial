package main
import (
	πg "grumpy"
	π_grumpyΓlibΓselect "grumpy/lib/select"
	π_grumpyΓlibΓsocket "grumpy/lib/socket"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ß := πg.InternStr("")
	ßClient := πg.InternStr("Client")
	ßIOError := πg.InternStr("IOError")
	ßSOL_SOCKET := πg.InternStr("SOL_SOCKET")
	ßSO_REUSEADDR := πg.InternStr("SO_REUSEADDR")
	ßServer := πg.InternStr("Server")
	ß__init__ := πg.InternStr("__init__")
	ß__main__ := πg.InternStr("__main__")
	ß__metaclass__ := πg.InternStr("__metaclass__")
	ß__module__ := πg.InternStr("__module__")
	ß__name__ := πg.InternStr("__name__")
	ßaccept := πg.InternStr("accept")
	ßbind := πg.InternStr("bind")
	ßbuf := πg.InternStr("buf")
	ßclose := πg.InternStr("close")
	ßcon := πg.InternStr("con")
	ßevloop := πg.InternStr("evloop")
	ßfileno := πg.InternStr("fileno")
	ßkeys := πg.InternStr("keys")
	ßlisten := πg.InternStr("listen")
	ßobject := πg.InternStr("object")
	ßon_acceptable := πg.InternStr("on_acceptable")
	ßon_readable := πg.InternStr("on_readable")
	ßon_writable := πg.InternStr("on_writable")
	ßpop := πg.InternStr("pop")
	ßread_waits := πg.InternStr("read_waits")
	ßrecv := πg.InternStr("recv")
	ßselect := πg.InternStr("select")
	ßsend := πg.InternStr("send")
	ßserver := πg.InternStr("server")
	ßsetblocking := πg.InternStr("setblocking")
	ßsetsockopt := πg.InternStr("setsockopt")
	ßsocket := πg.InternStr("socket")
	ßstart := πg.InternStr("start")
	ßwait_read := πg.InternStr("wait_read")
	ßwait_write := πg.InternStr("wait_write")
	ßwrite_waits := πg.InternStr("write_waits")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 *πg.Dict
	_ = πTemp003
	var πTemp004 []πg.FunctionArg
	_ = πTemp004
	var πTemp005 *πg.Object
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 *πg.Object
	_ = πTemp007
	var πTemp008 *πg.Object
	_ = πTemp008
	var πTemp009 *πg.Object
	_ = πTemp009
	var πTemp010 bool
	_ = πTemp010
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 1: import socket
		πF.SetLineno(1)
		if πTemp002, πE = πg.ImportModule(πF, "socket", []*πg.Code{π_grumpyΓlibΓsocket.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsocket.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 2: import select
		πF.SetLineno(2)
		if πTemp002, πE = πg.ImportModule(πF, "select", []*πg.Code{π_grumpyΓlibΓselect.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßselect.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 4: read_waits = {}
		πF.SetLineno(4)
		πTemp003 = πg.NewDict()
		πTemp001 = πTemp003.ToObject()
		if πE = πF.Globals().SetItem(πF, ßread_waits.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 5: write_waits = {}
		πF.SetLineno(5)
		πTemp003 = πg.NewDict()
		πTemp001 = πTemp003.ToObject()
		if πE = πF.Globals().SetItem(πF, ßwrite_waits.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 7: def wait_read(con, callback):
		πF.SetLineno(7)
		πTemp004 = make([]πg.FunctionArg, 2)
		πTemp004[0] = πg.FunctionArg{Name: "con", Def: nil}
		πTemp004[1] = πg.FunctionArg{Name: "callback", Def: nil}
		πTemp001 = πg.NewFunction(πg.NewCode("wait_read", "web_server.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µcon *πg.Object = πArgs[0]; _ = µcon
			var µcallback *πg.Object = πArgs[1]; _ = µcallback
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 8: read_waits[con.fileno()] = callback
				πF.SetLineno(8)
				if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
					continue
				}
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcallback); πE != nil {
					continue
				}
				if πTemp002, πE = πg.ResolveGlobal(πF, ßread_waits); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µcon, "con"); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, µcon, ßfileno, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp003 = πTemp005
				if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßwait_read.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 10: def wait_write(con, callback):
		πF.SetLineno(10)
		πTemp004 = make([]πg.FunctionArg, 2)
		πTemp004[0] = πg.FunctionArg{Name: "con", Def: nil}
		πTemp004[1] = πg.FunctionArg{Name: "callback", Def: nil}
		πTemp005 = πg.NewFunction(πg.NewCode("wait_write", "web_server.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µcon *πg.Object = πArgs[0]; _ = µcon
			var µcallback *πg.Object = πArgs[1]; _ = µcallback
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 11: write_waits[con.fileno()] = callback
				πF.SetLineno(11)
				if πE = πg.CheckLocal(πF, µcallback, "callback"); πE != nil {
					continue
				}
				if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcallback); πE != nil {
					continue
				}
				if πTemp002, πE = πg.ResolveGlobal(πF, ßwrite_waits); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µcon, "con"); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, µcon, ßfileno, nil); πE != nil {
					continue
				}
				if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp003 = πTemp005
				if πE = πg.SetItem(πF, πTemp002, πTemp003, πTemp001); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßwait_write.ToObject(), πTemp005); πE != nil {
			continue
		}
		// line 13: def evloop():
		πF.SetLineno(13)
		πTemp004 = make([]πg.FunctionArg, 0)
		πTemp006 = πg.NewFunction(πg.NewCode("evloop", "web_server.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µrs *πg.Object = πg.UnboundLocal; _ = µrs
			var µws *πg.Object = πg.UnboundLocal; _ = µws
			var µxs *πg.Object = πg.UnboundLocal; _ = µxs
			var µrfd *πg.Object = πg.UnboundLocal; _ = µrfd
			var µwfd *πg.Object = πg.UnboundLocal; _ = µwfd
			var πTemp001 bool
			_ = πTemp001
			var πTemp002 []*πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 []*πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 14: while 1:
				πF.SetLineno(14)
			Label1:
				if πTemp001, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
					continue
				}
				if !πTemp001 {
					goto Label2
				}
				// line 15: rs, ws, xs, = select.select(read_waits.keys(), write_waits.keys(), [])
				πF.SetLineno(15)
				πTemp002 = πF.MakeArgs(3)
				if πTemp003, πE = πg.ResolveGlobal(πF, ßread_waits); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp002[0] = πTemp003
				if πTemp003, πE = πg.ResolveGlobal(πF, ßwrite_waits); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßkeys, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, nil, nil); πE != nil {
					continue
				}
				πTemp002[1] = πTemp003
				πTemp005 = make([]*πg.Object, 0)
				πTemp003 = πg.NewList(πTemp005...).ToObject()
				πTemp002[2] = πTemp003
				if πTemp003, πE = πg.ResolveGlobal(πF, ßselect); πE != nil {
					continue
				}
				if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßselect, nil); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp004.Call(πF, πTemp002, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp002)
				if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp004}, πg.TieTarget{Target: &πTemp006}, πg.TieTarget{Target: &πTemp007}}}, πTemp003); πE != nil {
					continue
				}
				µrs = πTemp004
				µws = πTemp006
				µxs = πTemp007
				// line 16: for rfd in rs:
				πF.SetLineno(16)
				if πE = πg.CheckLocal(πF, µrs, "rs"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.Iter(πF, µrs); πE != nil {
					continue
				}
			Label3:
				if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
					isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
					if exc != nil {
						πE = exc
						continue
					}
					if !isStop {
						continue
					}
					πE = nil
					πF.RestoreExc(nil, nil)
					goto Label4
				}
				µrfd = πTemp004
				// line 17: read_waits.pop(rfd)()
				πF.SetLineno(17)
				πTemp002 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µrfd, "rfd"); πE != nil {
					continue
				}
				πTemp002[0] = µrfd
				if πTemp006, πE = πg.ResolveGlobal(πF, ßread_waits); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßpop, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp002)
				if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
					continue
				}
				goto Label3
				goto Label4
			Label4:
				// line 18: for wfd in ws:
				πF.SetLineno(18)
				if πE = πg.CheckLocal(πF, µws, "ws"); πE != nil {
					continue
				}
				if πTemp003, πE = πg.Iter(πF, µws); πE != nil {
					continue
				}
			Label5:
				if πTemp004, πE = πg.Next(πF, πTemp003); πE != nil {
					isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
					if exc != nil {
						πE = exc
						continue
					}
					if !isStop {
						continue
					}
					πE = nil
					πF.RestoreExc(nil, nil)
					goto Label6
				}
				µwfd = πTemp004
				// line 19: write_waits.pop(wfd)()
				πF.SetLineno(19)
				πTemp002 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µwfd, "wfd"); πE != nil {
					continue
				}
				πTemp002[0] = µwfd
				if πTemp006, πE = πg.ResolveGlobal(πF, ßwrite_waits); πE != nil {
					continue
				}
				if πTemp007, πE = πg.GetAttr(πF, πTemp006, ßpop, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp007.Call(πF, πTemp002, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp002)
				if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
					continue
				}
				goto Label5
				goto Label6
			Label6:
				goto Label1
				goto Label2
			Label2:
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßevloop.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 21: class Server(object):
		πF.SetLineno(21)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
			continue
		}
		πTemp002[0] = πTemp009
		πTemp003 = πg.NewDict()
		if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
			continue
		}
		_, πE = πg.NewCode("Server", "web_server.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp003
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 []πg.FunctionArg
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 22: def __init__(self, con):
				πF.SetLineno(22)
				πTemp002 = make([]πg.FunctionArg, 2)
				πTemp002[0] = πg.FunctionArg{Name: "self", Def: nil}
				πTemp002[1] = πg.FunctionArg{Name: "con", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "web_server.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µcon *πg.Object = πArgs[1]; _ = µcon
					var πTemp001 *πg.Object
					_ = πTemp001
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 23: self.con = con
						πF.SetLineno(23)
						if πE = πg.CheckLocal(πF, µcon, "con"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcon); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßcon, πTemp001); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 25: def start(self):
				πF.SetLineno(25)
				πTemp002 = make([]πg.FunctionArg, 1)
				πTemp002[0] = πg.FunctionArg{Name: "self", Def: nil}
				πTemp003 = πg.NewFunction(πg.NewCode("start", "web_server.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 26: wait_read(self.con, self.on_acceptable)
						πF.SetLineno(26)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßon_acceptable, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßwait_read); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßstart.ToObject(), πTemp003); πE != nil {
					continue
				}
				// line 28: def on_acceptable(self):
				πF.SetLineno(28)
				πTemp002 = make([]πg.FunctionArg, 1)
				πTemp002[0] = πg.FunctionArg{Name: "self", Def: nil}
				πTemp004 = πg.NewFunction(πg.NewCode("on_acceptable", "web_server.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µcon *πg.Object = πg.UnboundLocal; _ = µcon
					var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
					var πTemp001 bool
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 *πg.Object
					_ = πTemp004
					var πTemp005 []*πg.Object
					_ = πTemp005
					var πTemp006 *πg.BaseException
					_ = πTemp006
					var πTemp007 *πg.Traceback
					_ = πTemp007
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						case 1: goto Label1
						default: panic("unexpected function state")
						}
						// line 29: try:
						πF.SetLineno(29)
						πF.PushCheckpoint(1)
						// line 30: while 1:
						πF.SetLineno(30)
					Label3:
						if πTemp001, πE = πg.IsTrue(πF, πg.NewInt(1).ToObject()); πE != nil {
							continue
						}
						if !πTemp001 {
							goto Label4
						}
						// line 31: con, _ = self.con.accept()
						πF.SetLineno(31)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßaccept, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Children: []πg.TieTarget{πg.TieTarget{Target: &πTemp003}, πg.TieTarget{Target: &πTemp004}}}, πTemp002); πE != nil {
							continue
						}
						µcon = πTemp003
						µ_ = πTemp004
						// line 32: con.setblocking(0)
						πF.SetLineno(32)
						πTemp005 = πF.MakeArgs(1)
						πTemp005[0] = πg.NewInt(0).ToObject()
						if πE = πg.CheckLocal(πF, µcon, "con"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µcon, ßsetblocking, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						// line 33: Client(con)
						πF.SetLineno(33)
						πTemp005 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µcon, "con"); πE != nil {
							continue
						}
						πTemp005[0] = µcon
						if πTemp002, πE = πg.ResolveGlobal(πF, ßClient); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						goto Label3
						goto Label4
					Label4:
						πF.PopCheckpoint()
						goto Label2
					Label1:
						πTemp006, πTemp007 = πF.ExcInfo()
						if πTemp002, πE = πg.ResolveGlobal(πF, ßIOError); πE != nil {
							continue
						}
						if πTemp001, πE = πg.IsInstance(πF, πTemp006.ToObject(), πTemp002); πE != nil {
							continue
						}
						if πTemp001 {
							goto Label5
						}
						πE = πF.Raise(πTemp006.ToObject(), nil, πTemp007.ToObject())
						continue
						// line 34: except IOError:
						πF.SetLineno(34)
					Label5:
						// line 35: wait_read(self.con, self.on_acceptable)
						πF.SetLineno(35)
						πTemp005 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						πTemp005[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßon_acceptable, nil); πE != nil {
							continue
						}
						πTemp005[1] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßwait_read); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp005)
						πE = nil
						πF.RestoreExc(nil, nil)
						goto Label2
					Label2:
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßon_acceptable.ToObject(), πTemp004); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			return nil, πE
		}
		if πTemp008, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			return nil, πE
		}
		if πTemp008 == nil {
			πTemp008 = πg.TypeType.ToObject()
		}
		if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Server").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßServer.ToObject(), πTemp009); πE != nil {
			continue
		}
		// line 37: class Client(object):
		πF.SetLineno(37)
		πTemp002 = make([]*πg.Object, 1)
		if πTemp009, πE = πg.ResolveGlobal(πF, ßobject); πE != nil {
			continue
		}
		πTemp002[0] = πTemp009
		πTemp003 = πg.NewDict()
		if πTemp007, πE = πF.Globals().GetItem(πF, ß__name__.ToObject()); πE != nil {
			continue
		}
		if πE = πTemp003.SetItem(πF, ß__module__.ToObject(), πTemp007); πE != nil {
			continue
		}
		_, πE = πg.NewCode("Client", "web_server.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
			πClass := πTemp003
			_ = πClass
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 []πg.FunctionArg
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 38: def __init__(self, con):
				πF.SetLineno(38)
				πTemp002 = make([]πg.FunctionArg, 2)
				πTemp002[0] = πg.FunctionArg{Name: "self", Def: nil}
				πTemp002[1] = πg.FunctionArg{Name: "con", Def: nil}
				πTemp001 = πg.NewFunction(πg.NewCode("__init__", "web_server.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µcon *πg.Object = πArgs[1]; _ = µcon
					var πTemp001 *πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 39: self.con = con
						πF.SetLineno(39)
						if πE = πg.CheckLocal(πF, µcon, "con"); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp001}, µcon); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßcon, πTemp001); πE != nil {
							continue
						}
						// line 40: self.on_readable()
						πF.SetLineno(40)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp001, πE = πg.GetAttr(πF, µself, ßon_readable, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ß__init__.ToObject(), πTemp001); πE != nil {
					continue
				}
				// line 42: def on_readable(self):
				πF.SetLineno(42)
				πTemp002 = make([]πg.FunctionArg, 1)
				πTemp002[0] = πg.FunctionArg{Name: "self", Def: nil}
				πTemp003 = πg.NewFunction(πg.NewCode("on_readable", "web_server.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µdata *πg.Object = πg.UnboundLocal; _ = µdata
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 bool
					_ = πTemp004
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 43: data = self.con.recv(32 * 1024)
						πF.SetLineno(43)
						πTemp001 = πF.MakeArgs(1)
						if πTemp002, πE = πg.Mul(πF, πg.NewInt(32).ToObject(), πg.NewInt(1024).ToObject()); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßrecv, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µdata = πTemp002
						if πE = πg.CheckLocal(πF, µdata, "data"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.IsTrue(πF, µdata); πE != nil {
							continue
						}
						πTemp002 = πg.GetBool(!πTemp004).ToObject()
						if πTemp004, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							return nil, πE
						}
						if πTemp004 {
							goto Label1
						}
						goto Label2
						// line 44: if not data:
						πF.SetLineno(44)
					Label1:
						// line 45: wait_read(self.con, self.on_readable)
						πF.SetLineno(45)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßon_readable, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßwait_read); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						// line 46: return
						πF.SetLineno(46)
						return nil, nil
						goto Label2
					Label2:
						// line 47: self.buf = b"""HTTP/1.1 200 OK \r
						πF.SetLineno(47)
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πg.NewStr("HTTP/1.1 200 OK \r\nContent-Type: text/plain\r\nContent-Length: 6\r\n\r\nhello\n").ToObject()); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßbuf, πTemp002); πE != nil {
							continue
						}
						// line 53: self.on_writable()
						πF.SetLineno(53)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßon_writable, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, nil, nil); πE != nil {
							continue
						}
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßon_readable.ToObject(), πTemp003); πE != nil {
					continue
				}
				// line 55: def on_writable(self):
				πF.SetLineno(55)
				πTemp002 = make([]πg.FunctionArg, 1)
				πTemp002[0] = πg.FunctionArg{Name: "self", Def: nil}
				πTemp004 = πg.NewFunction(πg.NewCode("on_writable", "web_server.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
					var µself *πg.Object = πArgs[0]; _ = µself
					var µwrote *πg.Object = πg.UnboundLocal; _ = µwrote
					var πTemp001 []*πg.Object
					_ = πTemp001
					var πTemp002 *πg.Object
					_ = πTemp002
					var πTemp003 *πg.Object
					_ = πTemp003
					var πTemp004 *πg.Object
					_ = πTemp004
					var πTemp005 bool
					_ = πTemp005
					var πE *πg.BaseException; _ = πE
					for ; πF.State() >= 0; πF.PopCheckpoint() {
						switch πF.State() {
						case 0:
						default: panic("unexpected function state")
						}
						// line 56: wrote = self.con.send(self.buf)
						πF.SetLineno(56)
						πTemp001 = πF.MakeArgs(1)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßsend, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						µwrote = πTemp002
						// line 57: self.buf = self.buf[wrote:]
						πF.SetLineno(57)
						if πE = πg.CheckLocal(πF, µwrote, "wrote"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.SliceType.Call(πF, πg.Args{µwrote, πg.None, πg.None}, nil); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp004, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetItem(πF, πTemp004, πTemp002); πE != nil {
							continue
						}
						if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp002}, πTemp003); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πE = πg.SetAttr(πF, µself, ßbuf, πTemp002); πE != nil {
							continue
						}
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßbuf, nil); πE != nil {
							continue
						}
						if πTemp005, πE = πg.IsTrue(πF, πTemp002); πE != nil {
							return nil, πE
						}
						if πTemp005 {
							goto Label1
						}
						goto Label2
						// line 58: if self.buf:
						πF.SetLineno(58)
					Label1:
						// line 59: wait_write(self.con, self.on_writable)
						πF.SetLineno(59)
						πTemp001 = πF.MakeArgs(2)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						πTemp001[0] = πTemp002
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßon_writable, nil); πE != nil {
							continue
						}
						πTemp001[1] = πTemp002
						if πTemp002, πE = πg.ResolveGlobal(πF, ßwait_write); πE != nil {
							continue
						}
						if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
							continue
						}
						πF.FreeArgs(πTemp001)
						goto Label3
					Label2:
						// line 61: self.con.close()
						πF.SetLineno(61)
						if πE = πg.CheckLocal(πF, µself, "self"); πE != nil {
							continue
						}
						if πTemp002, πE = πg.GetAttr(πF, µself, ßcon, nil); πE != nil {
							continue
						}
						if πTemp003, πE = πg.GetAttr(πF, πTemp002, ßclose, nil); πE != nil {
							continue
						}
						if πTemp002, πE = πTemp003.Call(πF, nil, nil); πE != nil {
							continue
						}
						goto Label3
					Label3:
						return nil, nil
					}
					return nil, πE
				}), πF.Globals()).ToObject()
				if πE = πClass.SetItem(πF, ßon_writable.ToObject(), πTemp004); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}).Eval(πF, πF.Globals(), nil, nil)
		if πE != nil {
			return nil, πE
		}
		if πTemp008, πE = πTemp003.GetItem(πF, ß__metaclass__.ToObject()); πE != nil {
			return nil, πE
		}
		if πTemp008 == nil {
			πTemp008 = πg.TypeType.ToObject()
		}
		if πTemp009, πE = πTemp008.Call(πF, []*πg.Object{πg.NewStr("Client").ToObject(), πg.NewTuple(πTemp002...).ToObject(), πTemp003.ToObject()}, nil); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßClient.ToObject(), πTemp009); πE != nil {
			continue
		}
		// line 63: def server():
		πF.SetLineno(63)
		πTemp004 = make([]πg.FunctionArg, 0)
		πTemp007 = πg.NewFunction(πg.NewCode("server", "web_server.py", πTemp004, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µsock *πg.Object = πg.UnboundLocal; _ = µsock
			var µserver *πg.Object = πg.UnboundLocal; _ = µserver
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 64: sock = socket.socket()
				πF.SetLineno(64)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßsocket, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, nil, nil); πE != nil {
					continue
				}
				µsock = πTemp001
				// line 65: sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
				πF.SetLineno(65)
				πTemp003 = πF.MakeArgs(3)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSOL_SOCKET, nil); πE != nil {
					continue
				}
				πTemp003[0] = πTemp002
				if πTemp001, πE = πg.ResolveGlobal(πF, ßsocket); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßSO_REUSEADDR, nil); πE != nil {
					continue
				}
				πTemp003[1] = πTemp002
				πTemp003[2] = πg.NewInt(1).ToObject()
				if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, µsock, ßsetsockopt, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				// line 66: sock.setblocking(0)
				πF.SetLineno(66)
				πTemp003 = πF.MakeArgs(1)
				πTemp003[0] = πg.NewInt(0).ToObject()
				if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, µsock, ßsetblocking, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				// line 67: sock.bind(('', 8000))
				πF.SetLineno(67)
				πTemp003 = πF.MakeArgs(1)
				πTemp004 = make([]*πg.Object, 2)
				πTemp004[0] = ß.ToObject()
				πTemp004[1] = πg.NewInt(8000).ToObject()
				πTemp001 = πg.NewTuple(πTemp004...).ToObject()
				πTemp003[0] = πTemp001
				if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, µsock, ßbind, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				// line 68: sock.listen(128)
				πF.SetLineno(68)
				πTemp003 = πF.MakeArgs(1)
				πTemp003[0] = πg.NewInt(128).ToObject()
				if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, µsock, ßlisten, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				// line 69: server = Server(sock)
				πF.SetLineno(69)
				πTemp003 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µsock, "sock"); πE != nil {
					continue
				}
				πTemp003[0] = µsock
				if πTemp001, πE = πg.ResolveGlobal(πF, ßServer); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µserver = πTemp002
				// line 70: server.start()
				πF.SetLineno(70)
				if πE = πg.CheckLocal(πF, µserver, "server"); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, µserver, ßstart, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
					continue
				}
				// line 71: evloop()
				πF.SetLineno(71)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßevloop); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, nil, nil); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßserver.ToObject(), πTemp007); πE != nil {
			continue
		}
		if πTemp009, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp008, πE = πg.Eq(πF, πTemp009, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp010, πE = πg.IsTrue(πF, πTemp008); πE != nil {
			return nil, πE
		}
		if πTemp010 {
			goto Label1
		}
		goto Label2
		// line 73: if __name__ == '__main__':
		πF.SetLineno(73)
	Label1:
		// line 74: server()
		πF.SetLineno(74)
		if πTemp008, πE = πg.ResolveGlobal(πF, ßserver); πE != nil {
			continue
		}
		if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
			continue
		}
		goto Label2
	Label2:
		return nil, nil
	}
	return nil, πE
}
var Code *πg.Code
func main() {
	Code = πg.NewCode("<module>", "web_server.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
